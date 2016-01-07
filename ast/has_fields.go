package ast

// HasFields is an AST nodes with Field children.
type HasFields interface {
	//  AddField adds a Field to the node.
	AddField(*Field)
}
