package ast

// InnerJoin represents an inner join table relation in the SQL query.
type InnerJoin struct {
	Table *Table
	Left  *Field
	Right *Field
}

func (j *InnerJoin) SetTable(table *Table) {
	j.Table = table
}

func (j *InnerJoin) GetTable() *Table {
	return j.Table
}

func (j *InnerJoin) BuildQuery() string {
	return "INNER JOIN " + j.Table.BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
