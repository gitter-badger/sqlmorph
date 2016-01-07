package ast

// HasLimit is an AST node with limit value.
type HasLimit interface {
	// SetLimit sets the limit of the node.
	SetLimit(string)
}
