package ast

// HasTable is an AST node with table.
type HasTable interface {

	// SetTable sets the table of the node.
	SetTable(*Table)

	GetTable() *Table
}
