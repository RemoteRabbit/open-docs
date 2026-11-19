package render

import (
	"fmt"
	"net/url"
	"path/filepath"
	"sort"
	"strings"

	"github.com/remoterabbit/open-inspector/pkg/model"
)

type Position struct {
	Position model.Position
	LinkBase string
	Verbose  bool
}

// position formats a source Position as a document-relative source link followed
// by its line:column:byte range.
func position(opts Position) string {
	if opts.Position.Filename == "" {
		return ""
	}

	target, err := filepath.Rel(opts.LinkBase, filepath.FromSlash(opts.Position.Filename))
	if err != nil {
		target = opts.Position.Filename
	} else if !strings.HasPrefix(target, ".") {
		target = "./" + target
	}
	target = filepath.ToSlash(target)
	href := (&url.URL{Path: target}).EscapedPath()
	if opts.Position.Start.Line > 0 {
		href += fmt.Sprintf("#L%d", opts.Position.Start.Line)
		if opts.Position.End.Line > opts.Position.Start.Line {
			href += fmt.Sprintf("-L%d", opts.Position.End.Line)
		}
	}

	if opts.Verbose {
		return fmt.Sprintf("[`%s`](%s) [%d:%d:%d -> %d:%d:%d]",
			target, href,
			opts.Position.Start.Line, opts.Position.Start.Column, opts.Position.Start.Byte,
			opts.Position.End.Line, opts.Position.End.Column, opts.Position.End.Byte)
	} else {
		return fmt.Sprintf("[`%s`](%s)",
			target, href)
	}
}

func posPtr(p *model.Position, pos func(model.Position) string) string {
	if p == nil {
		return "-"
	}
	return pos(*p)
}

func exprCell(e *model.Expression) string {
	if e == nil {
		return "-"
	}
	return code(oneLine(e.Source))
}

func exprValueCell(e model.Expression) string {
	return code(oneLine(e.Source))
}

// typeCell renders a variable's type. It prefers the structured TypeSpec
// (open-inspector v0.7.0+), falling back to the raw typeexpr string.
func typeCell(raw string, spec *model.TypeSpec) string {
	if spec != nil {
		return code(typeSpecString(spec))
	}
	return code(raw)
}

// typeSpecString renders a structured TypeSpec tree as a single-line type
// expression, e.g. `object({name = string, tags = optional(map(string))})`.
func typeSpecString(t *model.TypeSpec) string {
	if t == nil {
		return ""
	}
	switch t.Kind {
	case model.TypeList, model.TypeSet, model.TypeMap:
		return string(t.Kind) + "(" + typeSpecString(t.Element) + ")"
	case model.TypeObject:
		names := make([]string, 0, len(t.Attributes))
		for name := range t.Attributes {
			names = append(names, name)
		}
		sort.Strings(names)
		parts := make([]string, 0, len(names))
		for _, name := range names {
			attr := t.Attributes[name]
			inner := typeSpecString(attr.Type)
			if attr.Optional {
				if attr.Default != nil {
					inner = "optional(" + inner + ", " + valueString(attr.Default) + ")"
				} else {
					inner = "optional(" + inner + ")"
				}
			}
			parts = append(parts, name+" = "+inner)
		}
		return "object({" + strings.Join(parts, ", ") + "})"
	case model.TypeTuple:
		parts := make([]string, 0, len(t.Elements))
		for _, el := range t.Elements {
			parts = append(parts, typeSpecString(el))
		}
		return "tuple([" + strings.Join(parts, ", ") + "])"
	case model.TypeDynamic:
		return "any"
	default:
		return string(t.Kind)
	}
}

// defaultCell renders a variable default. It prefers the decoded literal
// Value (open-inspector v0.7.0+), falling back to the raw expression source.
func defaultCell(expr *model.Expression, value *model.Value) string {
	if value != nil {
		return code(oneLine(valueString(value)))
	}
	return exprCell(expr)
}

// valueString renders a decoded constant Value as a single-line literal.
func valueString(v *model.Value) string {
	if v == nil {
		return ""
	}
	switch v.Kind {
	case model.ValueNull:
		return "null"
	case model.ValueString:
		return fmt.Sprintf("%q", v.String)
	case model.ValueNumber:
		return v.Number
	case model.ValueBool:
		return fmt.Sprintf("%t", v.Bool)
	case model.ValueList:
		parts := make([]string, 0, len(v.List))
		for _, el := range v.List {
			parts = append(parts, valueString(el))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case model.ValueTuple:
		parts := make([]string, 0, len(v.Tuple))
		for _, el := range v.Tuple {
			parts = append(parts, valueString(el))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case model.ValueMap:
		return objectString(v.Map)
	case model.ValueObject:
		return objectString(v.Object)
	default:
		return ""
	}
}

func objectString(entries map[string]*model.Value) string {
	names := make([]string, 0, len(entries))
	for name := range entries {
		names = append(names, name)
	}
	sort.Strings(names)
	parts := make([]string, 0, len(names))
	for _, name := range names {
		parts = append(parts, name+" = "+valueString(entries[name]))
	}
	return "{" + strings.Join(parts, ", ") + "}"
}

// refsCell renders the references an expression makes (open-inspector
// v0.6.0+) as a comma-separated list of canonical addresses.
func refsCell(refs []model.Reference) string {
	if len(refs) == 0 {
		return "-"
	}
	parts := make([]string, 0, len(refs))
	for _, r := range refs {
		addr := r.Address
		if r.Attribute != "" {
			addr += "." + r.Attribute
		}
		parts = append(parts, addr)
	}
	return code(strings.Join(parts, ", "))
}

// attrNamesCell lists the top-level attribute names captured in a resource's
// NestedBody (open-inspector v0.7.0+ replaces the old AttrNames slice).
func attrNamesCell(body *model.Body) string {
	if body == nil || len(body.Attributes) == 0 {
		return "-"
	}
	names := make([]string, 0, len(body.Attributes))
	for name := range body.Attributes {
		names = append(names, name)
	}
	sort.Strings(names)
	return code(strings.Join(names, ", "))
}

// inputsCell lists the input argument names passed to a module call.
func inputsCell(inputs map[string]model.Expression) string {
	if len(inputs) == 0 {
		return "-"
	}
	names := make([]string, 0, len(inputs))
	for name := range inputs {
		names = append(names, name)
	}
	sort.Strings(names)
	return code(strings.Join(names, ", "))
}

func boolCell(v bool) string {
	if v {
		return "✓"
	}
	return "-"
}

func boolPtrCell(v *bool) string {
	if v == nil {
		return "-"
	}
	if *v {
		return "✓"
	}
	return "✗"
}

func lifecycleCell(l *model.Lifecycle) string {
	if l == nil {
		return "-"
	}
	var parts []string
	if l.CreateBeforeDestroy != nil {
		parts = append(parts, fmt.Sprintf("create_before_destroy=%t", *l.CreateBeforeDestroy))
	}
	if l.PreventDestroy != nil {
		parts = append(parts, fmt.Sprintf("prevent_destroy=%t", *l.PreventDestroy))
	}
	if len(l.IgnoreChanges) > 0 {
		parts = append(parts, "ignore_changes=["+strings.Join(l.IgnoreChanges, ",")+"]")
	}
	if len(l.ReplaceTriggeredBy) > 0 {
		parts = append(parts, "replace_triggered_by=["+strings.Join(l.ReplaceTriggeredBy, ",")+"]")
	}
	if len(l.Preconditions) > 0 {
		parts = append(parts, fmt.Sprintf("preconditions=%d", len(l.Preconditions)))
	}
	if len(l.Postconditions) > 0 {
		parts = append(parts, fmt.Sprintf("postconditions=%d", len(l.Postconditions)))
	}
	if len(parts) == 0 {
		return "-"
	}
	return code(strings.Join(parts, "; "))
}

func mapCell(m map[string]string) string {
	if len(m) == 0 {
		return "-"
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var parts []string
	for _, k := range keys {
		parts = append(parts, k+"="+m[k])
	}
	return code(strings.Join(parts, ", "))
}

func code(s string) string {
	if s == "" {
		return "-"
	}
	return "`" + s + "`"
}

// oneLine collapses whitespace and escapes pipe characters so a value is safe
// to place inside a Markdown table cell.
func oneLine(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.Join(strings.Fields(s), " ")
	if s == "" {
		return "-"
	}
	return strings.ReplaceAll(s, "|", "\\|")
}

func sortedKeys(m map[string]model.ProviderRequirement) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// col builds a left-aligned column; chain .center() for other alignment.
func col[T any](header string, enabled bool, cell func(T) string) column[T] {
	return column[T]{header: header, enabled: enabled, cell: cell}
}

//nolint:unused // Kept for column alignment; no table centers a column yet.
func (c column[T]) center() column[T] { c.align = "center"; return c }

// sepFor returns a delimiter cell as wide as the rendered header cell
// (" Name " = len+2), with optional alignment colons.
func sepFor(header, align string) string {
	n := len(header) + 2
	switch align {
	case "center":
		return ":" + strings.Repeat("-", n-2) + ":"
	case "right":
		return strings.Repeat("-", n-1) + ":"
	default:
		return strings.Repeat("-", n)
	}
}

func writeTable[T any](b *strings.Builder, cols []column[T], rows []T) {
	var active []column[T]
	for _, c := range cols {
		if c.enabled {
			active = append(active, c)
		}
	}
	if len(active) == 0 {
		return
	}
	headers := make([]string, len(active))
	seps := make([]string, len(active))
	for i, c := range active {
		headers[i], seps[i] = c.header, sepFor(c.header, c.align)
	}
	fmt.Fprintf(b, "| %s |\n|%s|\n", strings.Join(headers, " | "), strings.Join(seps, "|"))
	for _, r := range rows {
		cells := make([]string, len(active))
		for i, c := range active {
			cells[i] = c.cell(r)
		}
		fmt.Fprintf(b, "| %s |\n", strings.Join(cells, " | "))
	}
	b.WriteString("\n")
}
