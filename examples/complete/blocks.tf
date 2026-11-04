# --- moved ------------------------------------------------------------------

moved {
  from = null_resource.old_name
  to   = null_resource.empty
}

moved {
  from = random_pet.single
  to   = random_pet.names[0]
}

moved {
  from = module.old_child
  to   = module.child
}

moved {
  from = module.child.aws_s3_bucket.legacy
  to   = module.child.aws_s3_bucket.this
}

# --- import -----------------------------------------------------------------

import {
  to = null_resource.empty
  id = "static-id-123"
}

import {
  to       = aws_iam_role.ignore_all
  id       = "complete-audit"
  provider = aws.audit
}

import {
  to = aws_s3_bucket.replica
  id = "${local.literal_string}-replica"
}

# OpenTofu / Terraform 1.7+ import with for_each.
import {
  for_each = var.map_of_objects
  to       = aws_s3_bucket.kitchen_sink[each.key]
  id       = "complete-${each.key}"
}

# --- removed ----------------------------------------------------------------

removed {
  from = null_resource.gone
}

removed {
  from = aws_s3_bucket.retired

  lifecycle {
    destroy = false
  }
}

removed {
  from = module.deprecated_child

  lifecycle {
    destroy = true
  }
}

# --- check ------------------------------------------------------------------

check "health_endpoint" {
  data "http" "probe" {
    url = "https://example.com/health"
  }

  assert {
    condition     = data.http.probe.status_code == 200
    error_message = "Health check returned ${data.http.probe.status_code}."
  }

  assert {
    condition     = can(jsondecode(data.http.probe.response_body).status)
    error_message = "Health body is not JSON with a status field."
  }
}

# Check without a scoped data source.
check "inputs_sane" {
  assert {
    condition     = var.number_default > 0
    error_message = "number_default must be positive."
  }
}

check "three_assertions" {
  assert {
    condition     = length(var.list_of_strings) > 0
    error_message = "list_of_strings must not be empty."
  }

  assert {
    condition     = var.map_of_numbers["small"] < var.map_of_numbers["large"]
    error_message = "small must be less than large."
  }

  assert {
    condition     = local.literal_bool
    error_message = "Always true."
  }
}
