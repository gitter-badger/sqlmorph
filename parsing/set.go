package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
	"github.com/s2gatev/sqlmorph/lexing"
)

const SetWithoutFieldsError = "SET statement must be followed by field assignment list."

// SetState parses SET SQL clauses along with their fields.
// ... SET u.Name=?, u.Age=? ...
type SetState struct {
	BaseState
}

func (s *SetState) Name() string {
	return "SET"
}

func (s *SetState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	target := result.(ast.HasFields)

	if token, _ := tokenizer.ReadToken(); token != lexing.SET {
		tokenizer.UnreadToken()
		return result, false
	}

	// Parse WHERE conditions.
	for {
		field := &ast.Field{}

		if token, fieldName := tokenizer.ReadToken(); token == lexing.LITERAL {
			field = parseField(fieldName)
		} else {
			wrongTokenPanic(SetWithoutFieldsError, fieldName)
		}

		if token, operator := tokenizer.ReadToken(); token != lexing.EQUALS {
			wrongTokenPanic(SetWithoutFieldsError, operator)
		}

		if token, value := tokenizer.ReadToken(); token == lexing.LITERAL || token == lexing.PLACEHOLDER {
			field.Value = value
		} else {
			wrongTokenPanic(SetWithoutFieldsError, value)
		}

		target.AddField(field)

		if token, _ := tokenizer.ReadToken(); token != lexing.COMMA {
			tokenizer.UnreadToken()
			break
		}
	}

	return result, true
}
