locals {
  literal_string = "plain"
  literal_number = 7
  literal_bool   = true
  literal_null   = null
  literal_list   = [1, 2, 3]
  literal_map    = { a = 1, b = "two" }

  from_var       = var.string_default
  interpolated   = "${var.string_default}-${local.literal_number}"
  conditional    = var.bool_true ? "yes" : "no"
  for_list       = [for s in var.list_of_strings : upper(s)]
  for_map        = { for k, v in var.map_of_numbers : k => v * 2 if v > 1 }
  splat          = random_pet.names[*].id
  function_chain = join(",", sort(distinct(concat(var.list_of_strings, ["z"]))))
  many_refs      = merge(var.map_of_objects, { extra = { instance_type = local.literal_string } }, module.child.passthrough)
  pipe_value     = "left | right"
  heredoc_value  = <<-EOT
    line one
    line two
  EOT
}

# Second locals block in the same file.
locals {
  regions_list = tolist(var.regions)
  primary      = local.regions_list[0]
}
