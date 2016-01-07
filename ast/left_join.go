package ast

// LeftJoin represents an inner join table relation in the SQL query.
type LeftJoin struct {
	Table *Table
	Left  *Field
	Right *Field
}

func (j *LeftJoin) SetTable(table *Table) {
	j.Table = table
}

func (j *LeftJoin) GetTable() *Table {
	return j.Table
}

func (j *LeftJoin) BuildQuery() string {
	return "LEFT JOIN " + j.Table.BuildQuery() + " ON " +
		j.Left.BuildQuery() + "=" + j.Right.BuildQuery()
}
