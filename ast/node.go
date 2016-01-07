package ast

// Node is a generic representation of an element in the SQL AST.
type Node interface {
	// BuildQuery creates a valid SQL query from the AST node.
	BuildQuery() string
}
