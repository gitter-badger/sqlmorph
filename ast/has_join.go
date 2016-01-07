package ast

// HasJoin is an AST node with join table.
type HasJoin interface {
	AddJoin(Join)
}
