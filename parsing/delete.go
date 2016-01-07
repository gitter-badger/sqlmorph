package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

// DeleteState parses SELECT SQL clauses along with the desired fields.
// SELECT u.Name, u.Age ...
type DeleteState struct {
	BaseState
}

func (s *DeleteState) Name() string {
	return "DELETE"
}

func (s *DeleteState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	if token, _ := tokenizer.ReadToken(); token != DELETE {
		tokenizer.UnreadToken()
		return result, false
	} else {
		target := ast.NewDelete()

		return target, true
	}
}
