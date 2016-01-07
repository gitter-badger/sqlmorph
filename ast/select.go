package ast

// Select represents a SQL SELECT statement.
type Select struct {
	Fields     []*Field
	Table      *Table
	JoinTables []Join
	Conditions []*EqualsCondition
	Limit      string
	Offset     string
}

func NewSelect() *Select {
	return &Select{}
}

func (ss *Select) AddField(field *Field) {
	ss.Fields = append(ss.Fields, field)
}

func (ss *Select) AddCondition(condition *EqualsCondition) {
	ss.Conditions = append(ss.Conditions, condition)
}

func (ss *Select) SetTable(table *Table) {
	ss.Table = table
}

func (ss *Select) GetTable() *Table {
	return ss.Table
}

func (ss *Select) SetLimit(limit string) {
	ss.Limit = limit
}

func (ss *Select) SetOffset(offset string) {
	ss.Offset = offset
}

func (ss *Select) AddJoin(join Join) {
	ss.JoinTables = append(ss.JoinTables, join)
}

func (ss *Select) BuildQuery() string {
	query := ""

	// Build SELECT part.
	fieldsPart := ""
	for index, field := range ss.Fields {
		if index != 0 {
			fieldsPart += ", "
		}
		fieldsPart += field.BuildQuery()
	}
	query += "SELECT " + fieldsPart

	// Build FROM part.
	query += " FROM " + ss.Table.BuildQuery()

	for _, join := range ss.JoinTables {
		query += " " + join.BuildQuery()
	}

	// Build WHERE part.
	if len(ss.Conditions) > 0 {
		conditionsPart := ""
		for index, condition := range ss.Conditions {
			if index != 0 {
				conditionsPart += " AND "
			}
			conditionsPart += condition.BuildQuery()
		}
		query += " WHERE " + conditionsPart
	}

	// Build LIMIT part.
	if ss.Limit != "" {
		query += " LIMIT " + ss.Limit
	}

	// Build OFFSET part.
	if ss.Offset != "" {
		query += " OFFSET " + ss.Offset
	}
	return query
}
