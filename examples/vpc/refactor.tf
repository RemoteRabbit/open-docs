# Subnets used to be indexed by count; they are now keyed by AZ.
moved {
  from = aws_subnet.public[0]
  to   = aws_subnet.public["us-east-1a"]
}

moved {
  from = aws_subnet.public[1]
  to   = aws_subnet.public["us-east-1b"]
}

moved {
  from = aws_nat_gateway.main
  to   = aws_nat_gateway.this[0]
}

moved {
  from = module.endpoints_legacy
  to   = module.endpoints
}

import {
  to = aws_vpc.this
  id = var.import_existing_vpc_id
}

import {
  to = aws_internet_gateway.this[0]
  id = "igw-0123456789abcdef0"
}

removed {
  from = aws_default_security_group.this

  lifecycle {
    destroy = false
  }
}

removed {
  from = aws_vpn_connection.legacy
}
