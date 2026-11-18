# Built-in defaults for open-docs. A user's .open-docs.hcl is decoded on top
# of these values, so any attribute set there overrides the one here.
main {
  docsfile = "README.md" # the markdown file to update
  mode     = "stdout"    # must be `stdout`, `replace`, or `inject`
  header   = true        # `This documentation was generated with [open-docs v1](https://github.com/RemoteRabbit/open-docs)`
  verbose = false        # include byte and column locations in position renders
  version = true         # include the open-docs version header in markdown render
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

data {
  render    = true
	name      = true
	mode      = true
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

ephemeral {
  render    = true
	name      = true
	mode      = true
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

modules {
  render = true
  name = true
  source = true
  version = true
  count = true
  for_each = true
  depends_on = true
  providers = true
  inputs = true
  comment = true
  position = true
}

diagnostics {
  render = true
  summary = true
  detail = true
  subject = true
  context = true
}

moved {
  render = true
  from = true
  to = true
  position = true
}

imports {
  render = true
  to = true
  id = true
  provider = true
  position = true
}

removed {
  render = true
  from = true
  destroy_on_drop = true
  position = true
}
