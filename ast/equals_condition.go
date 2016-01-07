package ast

// EqualsCondition represents an equality field-value relation.
type EqualsCondition struct {
	Field *Field
	Value string
}

func (c *EqualsCondition) BuildQuery() string {
	return c.Field.BuildQuery() + "=" + c.Value
}
