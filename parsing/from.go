package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
)

const FromWithoutTargetError = "FROM statement must be followed by a target class."

// FromState parses FROM SQL clauses along with the table name and alias.
// ... FROM User u ...
type FromState struct {
	BaseState
}

func (s *FromState) Name() string {
	return "FROM"
}

func (s *FromState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	target := result.(ast.HasTable)

	if token, _ := tokenizer.ReadToken(); token != FROM {
		tokenizer.UnreadToken()
		return result, false
	}

	table := &ast.Table{}

	if token, tableName := tokenizer.ReadToken(); token == LITERAL {
		table.Name = tableName
	} else {
		wrongTokenPanic(FromWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == LITERAL {
		table.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	target.SetTable(table)

	return result, true
}
