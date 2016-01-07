package ast

// Field represents a struct field in the SQL query.
type Field struct {
	Target string
	Name   string
	Value  string
}

func (f *Field) BuildQuery() string {
	field := ""
	if f.Target != "" {
		field += f.Target + "."
	}
	field += f.Name
	if f.Value != "" {
		field += "=" + f.Value
	}
	return field
}
