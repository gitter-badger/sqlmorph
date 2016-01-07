package ast

// RightJoin represents an inner join table relation in the SQL query.
type RightJoin struct {
	Table *Table
	Left  *Field
	Right *Field
}

func (j *RightJoin) SetTable(table *Table) {
	j.Table = table
}

func (j *RightJoin) GetTable() *Table {
	return j.Table
}

func (j *RightJoin) BuildQuery() string {
	return "RIGHT JOIN " + j.Table.BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
