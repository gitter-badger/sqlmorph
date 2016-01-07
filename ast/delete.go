package ast

// Delete represents an DELETE SQL query AST node.
type Delete struct {
	Table      *Table
	Conditions []*EqualsCondition
}

func NewDelete() *Delete {
	return &Delete{}
}

func (d *Delete) AddCondition(condition *EqualsCondition) {
	d.Conditions = append(d.Conditions, condition)
}

func (d *Delete) SetTable(table *Table) {
	d.Table = table
}

func (d *Delete) GetTable() *Table {
	return d.Table
}

func (d *Delete) BuildQuery() string {
	query := "DELETE FROM " + d.Table.BuildQuery()

	if len(d.Conditions) > 0 {
		conditionsPart := ""
		for index, condition := range d.Conditions {
			if index != 0 {
				conditionsPart += " AND "
			}
			conditionsPart += condition.BuildQuery()
		}
		query += " WHERE " + conditionsPart
	}

	return query
}
