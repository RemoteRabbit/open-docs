# Built-in defaults for open-docs. A user's .open-docs.hcl is decoded on top
# of these values, so any attribute set there overrides the one here.
main {
  docsfile = "README.md"
  mode     = "stdout"
  header   = true
}

required {
  render = true
  name = true
  source = true
  version = true
  alias = false
  position = true
}

providers {
  render = true
  name = true
  alias = false
  for_each = false
  position = true
}

variables {
  render = true
  name = true
  type = true
  default = true
  required = true
  description = true
  comment = true
  sensitive = true
  nullable = true
  ephemeral = true
  validation = true
  position = true
}
