package ast

// HasOffset is an AST node with offset value.
type HasOffset interface {
	// SetOffset sets the offset value of the node.
	SetOffset(string)
}
