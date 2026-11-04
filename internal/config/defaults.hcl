# Built-in defaults for open-docs. A user's .open-docs.hcl is decoded on top
# of these values, so any attribute set there overrides the one here.
main {
  docsfile = "README.md"
  mode     = "stdout"
  # mode = "replace"
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
  alias = true
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

outputs {
  render = true
  name = true
  value = true
  references = true
  description = true
  comment = true
  sensitive = true
  ephemeral = true
  depends_on = true
  position = true
}

locals {
  render = true
  name = true
  value = true
  references = true
  position = true
}

resources {
  render    = true
	name      = true
	mode      = false
	type      = true
	provider  = false
	count     = true
	for_each   = true
	depends_on = true
	nested_body= true
	lifecycle = true
	comment   = true
	position  = true
}
