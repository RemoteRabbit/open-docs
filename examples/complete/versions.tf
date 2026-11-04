# OpenTofu-flavored module that exercises every block type open-inspector
# understands, plus formatting edge cases for the Markdown renderer. It is
# not meant to be applied.

terraform {
  required_version = ">= 1.10.0"

  required_providers {
    aws = {
      source                = "hashicorp/aws"
      version               = "~> 5.70"
      configuration_aliases = [aws.replica, aws.shared_services]
    }
    random = {
      source  = "hashicorp/random"
      version = ">= 3.6.0, < 4.0.0"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "~> 4.0"
    }
    http = {
      source = "hashicorp/http"
    }
    local = {
      source  = "registry.opentofu.org/hashicorp/local"
      version = "2.5.2"
    }
    # No source: implied hashicorp/<name>.
    time = {}
  }

  encryption {
    key_provider "pbkdf2" "primary" {
      passphrase = var.state_passphrase
    }

    method "aes_gcm" "primary" {
      keys = key_provider.pbkdf2.primary
    }

    method "unencrypted" "migrate" {}

    state {
      method = method.aes_gcm.primary

      fallback {
        method = method.unencrypted.migrate
      }
    }

    plan {
      method = method.aes_gcm.primary
    }
  }
}

# A second terraform block: both constraints should be collected.
terraform {
  required_version = "!= 1.10.1"
}
