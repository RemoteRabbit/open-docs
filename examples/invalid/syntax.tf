# Missing closing brace: a hard HCL syntax error with a source range.
resource "null_resource" "broken" {
  triggers = {
    name = "missing closing brace"
}
