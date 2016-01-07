package ast

// Table represents a table in the SQL query.
type Table struct {
	Name  string
	Alias string
}

func (t *Table) BuildQuery() string {
	query := t.Name
	if t.Alias != "" {
		query += " " + t.Alias
	}
	return query
}
