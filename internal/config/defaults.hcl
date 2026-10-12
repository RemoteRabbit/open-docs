# Built-in defaults for open-docs. A user's .open-docs.hcl is decoded on top
# of these values, so any attribute set there overrides the one here.
main {
  docsfile = "README.md"
  mode     = "stdout"
  header   = true
}
