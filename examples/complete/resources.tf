# --- Managed resources ------------------------------------------------------

# Minimal resource: no attributes, no meta-arguments.
resource "null_resource" "empty" {}

resource "random_pet" "names" {
  count = 3

  length    = 2
  separator = "-"
  prefix    = var.string_default
}

resource "random_string" "per_key" {
  for_each = var.map_of_numbers

  length  = each.value + 8
  special = false
}

resource "random_password" "db" {
  length           = 32
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}

/*
 * Resource with every meta-argument and a full lifecycle block.
 */
resource "aws_s3_bucket" "kitchen_sink" {
  for_each = { for name, cfg in var.map_of_objects : name => cfg if cfg.max_size > 1 }
  provider = aws.by_region["us-east-1"]

  bucket        = "complete-${each.key}-${random_pet.names[0].id}"
  force_destroy = false

  tags = {
    Name  = each.key
    Owner = "platform | infra"
  }

  depends_on = [
    random_pet.names,
    null_resource.empty,
    module.child,
  ]

  lifecycle {
    create_before_destroy = true
    prevent_destroy       = true
    ignore_changes        = [tags, force_destroy]
    replace_triggered_by  = [random_password.db.id, null_resource.empty]

    precondition {
      condition     = length(each.key) < 20
      error_message = "Bucket key too long."
    }

    precondition {
      condition     = each.value.min_size <= each.value.max_size
      error_message = "min_size must not exceed max_size."
    }

    postcondition {
      condition     = self.bucket != ""
      error_message = "Bucket name must not be empty."
    }
  }
}

resource "aws_s3_bucket" "replica" {
  provider = aws.replica
  bucket   = "complete-replica"
}

resource "aws_iam_role" "ignore_all" {
  provider = aws.audit
  name     = "complete-audit"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect    = "Allow"
      Action    = "sts:AssumeRole"
      Principal = { Service = "ec2.amazonaws.com" }
    }]
  })

  lifecycle {
    ignore_changes = all
  }
}

# Nested blocks, dynamic blocks, and provisioners.
resource "aws_security_group" "nested" {
  name_prefix = "complete-"
  description = "Nested and dynamic block capture."

  dynamic "ingress" {
    for_each = var.simple_object.enabled ? [var.simple_object.port, 443] : []
    iterator = port
    content {
      from_port   = port.value
      to_port     = port.value
      protocol    = "tcp"
      cidr_blocks = ["0.0.0.0/0"]
    }
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  timeouts {
    create = "10m"
    delete = "15m"
  }

  provisioner "local-exec" {
    command = "echo created ${self.id}"
  }

  provisioner "local-exec" {
    when       = destroy
    command    = "echo destroyed"
    on_failure = continue
  }

  connection {
    type = "ssh"
    host = "127.0.0.1"
  }
}

resource "tls_private_key" "ca" {
  algorithm   = "ECDSA"
  ecdsa_curve = "P384"
}

resource "local_file" "rendered" {
  filename        = "${path.module}/out/rendered.txt"
  content         = templatefile("${path.module}/templates/example.tpl", { names = random_pet.names[*].id })
  file_permission = "0640"
}

resource "time_sleep" "wait" {
  create_duration = "30s"
  depends_on      = [aws_s3_bucket.kitchen_sink]
}

resource "terraform_data" "replacement" {
  input            = var.string_default
  triggers_replace = [var.number_default, var.bool_true]
}

# --- Data sources -----------------------------------------------------------

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "per_region" {
  for_each = var.regions
  provider = aws.by_region[each.key]

  statement {
    actions   = ["s3:GetObject"]
    resources = ["arn:aws:s3:::complete-${each.key}/*"]
  }
}

data "http" "health" {
  count = var.bool_true ? 1 : 0

  url = "https://example.com/health"

  request_headers = {
    Accept = "application/json"
  }

  lifecycle {
    postcondition {
      condition     = self.status_code == 200
      error_message = "Health endpoint returned ${self.status_code}."
    }
  }
}

data "local_file" "readme" {
  filename   = "${path.module}/README.md"
  depends_on = [local_file.rendered]
}

# --- Ephemeral resources ----------------------------------------------------

ephemeral "random_password" "session" {
  length  = 24
  special = true
}

ephemeral "aws_secretsmanager_secret_version" "db" {
  count    = var.bool_true ? 1 : 0
  provider = aws.audit

  secret_id = "complete/db"

  depends_on = [aws_iam_role.ignore_all]
}

ephemeral "tls_private_key" "per_region" {
  for_each = var.regions

  algorithm = "ED25519"
}
