package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const UpdateWithoutTargetError = "UPDATE statement must be followed by a target class."

// UpdateState parses UPDATE SQL clauses along with the target table.
// UPDATE User u ...
type UpdateState struct {
	BaseState
}

func (s *UpdateState) Name() string {
	return "UPDATE"
}

func (s *UpdateState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	target := ast.NewUpdate()

	if token, _ := tokenizer.ReadToken(); token != UPDATE {
		tokenizer.UnreadToken()
		return result, false
	}

	table := &ast.Table{}

	if token, tableName := tokenizer.ReadToken(); token == LITERAL {
		table.Name = tableName
	} else {
		wrongTokenPanic(UpdateWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == LITERAL {
		table.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	target.SetTable(table)

	return target, true
}
