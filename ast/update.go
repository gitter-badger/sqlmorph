package ast

// Update represents an Update SQL query AST node.
type Update struct {
	Fields     []*Field
	Table      *Table
	Conditions []*EqualsCondition
}

func NewUpdate() *Update {
	return &Update{}
}

func (u *Update) AddField(field *Field) {
	u.Fields = append(u.Fields, field)
}

func (u *Update) AddCondition(condition *EqualsCondition) {
	u.Conditions = append(u.Conditions, condition)
}

func (u *Update) SetTable(table *Table) {
	u.Table = table
}

func (u *Update) GetTable() *Table {
	return u.Table
}

func (u *Update) BuildQuery() string {
	query := "UPDATE " + u.Table.BuildQuery()

	fieldsPart := ""
	for index, field := range u.Fields {
		if index != 0 {
			fieldsPart += ", "
		}
		fieldsPart += field.BuildQuery()
	}
	query += " SET " + fieldsPart

	// Build WHERE part.
	if len(u.Conditions) > 0 {
		conditionsPart := ""
		for index, condition := range u.Conditions {
			if index != 0 {
				conditionsPart += " AND "
			}
			conditionsPart += condition.BuildQuery()
		}
		query += " WHERE " + conditionsPart
	}

	return query
}
