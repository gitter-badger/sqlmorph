package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

// State is a state in the parser state machine.
type State interface {

	// Name returns the name of this state.
	Name() string

	// Parse parses a node from the tokens available in the tokenizer.
	Parse(ast.Node, *Tokenizer) (ast.Node, bool)

	// Next returns the next reachable states in the parser.
	Next() []State

	// IsEndState returns if the state is a possible final state for the  parser.
	IsEndState() bool
}
