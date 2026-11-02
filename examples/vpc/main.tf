# The VPC itself. Protected from accidental destruction in every
# environment because subnets, ENIs, and peering all hang off it.
resource "aws_vpc" "this" {
  cidr_block           = var.cidr_block
  enable_dns_hostnames = var.enable_dns_hostnames
  enable_dns_support   = var.enable_dns_support

  tags = merge(local.tags, { Name = local.name_prefix })

  lifecycle {
    prevent_destroy = true
    ignore_changes  = [tags["CreatedAt"]]
  }
}

resource "aws_vpc_ipv4_cidr_block_association" "secondary" {
  for_each = toset(var.secondary_cidr_blocks)

  vpc_id     = aws_vpc.this.id
  cidr_block = each.value
}

resource "aws_internet_gateway" "this" {
  count = length(var.public_subnets) > 0 ? 1 : 0

  vpc_id = aws_vpc.this.id
  tags   = { Name = "${local.name_prefix}-igw" }
}

# --- Subnets ---------------------------------------------------------------

resource "aws_subnet" "public" {
  for_each = local.public_subnets

  vpc_id                  = aws_vpc.this.id
  cidr_block              = each.value
  availability_zone       = each.key
  map_public_ip_on_launch = true

  tags = merge(var.subnet_tags.public, {
    Name = "${local.name_prefix}-public-${each.key}"
    Tier = "public"
  })
}

resource "aws_subnet" "private" {
  for_each = local.private_subnets

  vpc_id            = aws_vpc.this.id
  cidr_block        = each.value
  availability_zone = each.key

  tags = merge(var.subnet_tags.private, {
    Name = "${local.name_prefix}-private-${each.key}"
    Tier = "private"
  })
}

resource "aws_subnet" "database" {
  count = length(var.database_subnets)

  vpc_id            = aws_vpc.this.id
  cidr_block        = var.database_subnets[count.index]
  availability_zone = local.azs[count.index]

  tags = merge(var.subnet_tags.database, {
    Name = "${local.name_prefix}-db-${local.azs[count.index]}"
    Tier = "database"
  })
}

resource "aws_db_subnet_group" "database" {
  count = local.create_database_subnets ? 1 : 0

  name       = "${local.name_prefix}-db"
  subnet_ids = aws_subnet.database[*].id
}

# --- NAT -------------------------------------------------------------------

resource "aws_eip" "nat" {
  count = var.nat_gateway_config.reuse_eips ? 0 : local.nat_gateway_count

  domain = "vpc"
  tags   = { Name = "${local.name_prefix}-nat-${count.index}" }

  depends_on = [aws_internet_gateway.this]
}

resource "aws_nat_gateway" "this" {
  count = local.nat_gateway_count

  allocation_id     = var.nat_gateway_config.reuse_eips ? var.nat_gateway_config.eip_allocation_ids[count.index] : aws_eip.nat[count.index].id
  subnet_id         = values(aws_subnet.public)[count.index].id
  connectivity_type = var.nat_gateway_config.connectivity_type

  tags = merge(var.nat_gateway_config.tags, { Name = "${local.name_prefix}-nat-${count.index}" })

  depends_on = [aws_internet_gateway.this, aws_eip.nat]

  lifecycle {
    create_before_destroy = true

    precondition {
      condition     = !var.nat_gateway_config.reuse_eips || length(var.nat_gateway_config.eip_allocation_ids) >= local.nat_gateway_count
      error_message = "reuse_eips requires one allocation ID per NAT gateway."
    }
  }
}

# --- Routing ---------------------------------------------------------------

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.this.id
  tags   = { Name = "${local.name_prefix}-public" }
}

resource "aws_route" "public_internet" {
  route_table_id         = aws_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.this[0].id

  timeouts {
    create = "5m"
  }
}

resource "aws_route_table" "private" {
  for_each = local.private_subnets

  vpc_id = aws_vpc.this.id
  tags   = { Name = "${local.name_prefix}-private-${each.key}" }

  dynamic "route" {
    for_each = local.nat_gateway_count > 0 ? [1] : []
    content {
      cidr_block     = "0.0.0.0/0"
      nat_gateway_id = aws_nat_gateway.this[var.single_nat_gateway ? 0 : index(local.azs, each.key)].id
    }
  }

  lifecycle {
    ignore_changes = [propagating_vgws]
  }
}

resource "aws_route_table_association" "public" {
  for_each = aws_subnet.public

  subnet_id      = each.value.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "private" {
  for_each = aws_subnet.private

  subnet_id      = each.value.id
  route_table_id = aws_route_table.private[each.key].id
}

# --- Flow logs -------------------------------------------------------------

resource "aws_cloudwatch_log_group" "flow_logs" {
  count = var.flow_logs.enabled && var.flow_logs.destination == "cloud-watch-logs" ? 1 : 0

  name              = "/aws/vpc/${local.name_prefix}/flow-logs"
  retention_in_days = var.flow_logs.retention_in_days

  lifecycle {
    prevent_destroy = false
  }
}

resource "aws_iam_role" "flow_logs" {
  count = var.flow_logs.enabled ? 1 : 0

  name_prefix        = "${local.name_prefix}-flow-logs-"
  assume_role_policy = data.aws_iam_policy_document.flow_logs_assume.json
}

resource "aws_iam_role_policy" "flow_logs" {
  count = var.flow_logs.enabled ? 1 : 0

  role   = aws_iam_role.flow_logs[0].id
  policy = data.aws_iam_policy_document.flow_logs_publish.json
}

resource "aws_flow_log" "this" {
  count = var.flow_logs.enabled ? 1 : 0

  vpc_id               = aws_vpc.this.id
  traffic_type         = var.flow_logs.traffic_type
  log_destination_type = var.flow_logs.destination
  log_destination      = local.flow_log_destination_arn
  iam_role_arn         = aws_iam_role.flow_logs[0].arn

  depends_on = [aws_iam_role_policy.flow_logs]

  lifecycle {
    replace_triggered_by = [aws_iam_role.flow_logs[0].id]

    postcondition {
      condition     = self.traffic_type == var.flow_logs.traffic_type
      error_message = "Flow log traffic type drifted from configuration."
    }
  }
}

# --- Optional VPN ----------------------------------------------------------

resource "aws_vpn_gateway" "this" {
  count = local.create_vpn ? 1 : 0

  vpc_id = aws_vpc.this.id
}

resource "random_id" "vpn_suffix" {
  count = local.create_vpn ? 1 : 0

  byte_length = 4
  keepers = {
    vpc_id = aws_vpc.this.id
  }
}
