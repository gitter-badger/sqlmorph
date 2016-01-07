package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const WhereWithoutConditionsError = "WHERE statement must be followed by condition list."

// WhereState parses WHERE SQL clauses along with their conditions.
// ... WHERE u.Age=? ...
type WhereState struct {
	BaseState
}

func (s *WhereState) Name() string {
	return "WHERE"
}

func (s *WhereState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	target := result.(ast.HasConditions)

	if token, _ := tokenizer.ReadToken(); token != WHERE {
		tokenizer.UnreadToken()
		return result, false
	}

	// Parse WHERE conditions.
	for {
		condition := &ast.EqualsCondition{}

		if token, field := tokenizer.ReadToken(); token == LITERAL {
			condition.Field = parseField(field)
		} else {
			wrongTokenPanic(WhereWithoutConditionsError, field)
		}

		if token, operator := tokenizer.ReadToken(); token != EQUALS {
			wrongTokenPanic(WhereWithoutConditionsError, operator)
		}

		if token, value := tokenizer.ReadToken(); token == LITERAL || token == PLACEHOLDER {
			condition.Value = value
		} else {
			wrongTokenPanic(WhereWithoutConditionsError, value)
		}

		target.AddCondition(condition)

		if token, _ := tokenizer.ReadToken(); token != AND {
			tokenizer.UnreadToken()
			break
		}
	}

	return result, true
}
