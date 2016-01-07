package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
	"github.com/s2gatev/sqlmorph/lexing"
)

const SelectWithoutFieldsError = "SELECT statement must be followed by field list."

// SelectState parses SELECT SQL clauses along with the desired fields.
// SELECT u.Name, u.Age ...
type SelectState struct {
	BaseState
}

func (s *SelectState) Name() string {
	return "SELECT"
}

func (s *SelectState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	if token, _ := tokenizer.ReadToken(); token != lexing.SELECT {
		tokenizer.UnreadToken()
		return result, false
	}

	target := ast.NewSelect()

	// Parse fields.
	for {
		token, value := tokenizer.ReadToken()
		if token == lexing.LITERAL || token == lexing.ASTERISK {
			target.AddField(parseField(value))
		} else {
			wrongTokenPanic(SelectWithoutFieldsError, value)
		}

		if token, _ := tokenizer.ReadToken(); token != lexing.COMMA {
			tokenizer.UnreadToken()
			break
		}
	}

	return target, true
}