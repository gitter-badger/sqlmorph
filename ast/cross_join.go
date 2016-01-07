package ast

// CrossJoin represents an inner join table relation in the SQL query.
type CrossJoin struct {
	Table *Table
}

func (j *CrossJoin) SetTable(table *Table) {
	j.Table = table
}

func (j *CrossJoin) GetTable() *Table {
	return j.Table
}

func (j *CrossJoin) BuildQuery() string {
	return "CROSS JOIN " + j.Table.BuildQuery()
}
