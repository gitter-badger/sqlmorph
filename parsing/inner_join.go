package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
	"github.com/s2gatev/sqlmorph/lexing"
)

const (
	InnerWithoutJoinError         = "Expected JOIN following INNER."
	InnerJoinWithoutTargetError   = "INNER JOIN statement must be followed by a target class."
	InnerJoinWithoutOnError       = "INNER JOIN statement must have an ON clause."
	InnerJoinWrongJoinFieldsError = "Wrong join fields in INNER JOIN statement."
)

type InnerJoinState struct {
	BaseState
}

func (s *InnerJoinState) Name() string {
	return "INNER JOIN"
}

func (s *InnerJoinState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	target := result.(ast.HasJoin)

	if token, _ := tokenizer.ReadToken(); token != lexing.INNER {
		tokenizer.UnreadToken()
		return result, false
	}

	if token, value := tokenizer.ReadToken(); token != lexing.JOIN {
		wrongTokenPanic(InnerWithoutJoinError, value)
	}

	join := &ast.InnerJoin{}
	table := &ast.Table{}

	if token, tableName := tokenizer.ReadToken(); token == lexing.LITERAL {
		table.Name = tableName
	} else {
		wrongTokenPanic(InnerJoinWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == lexing.LITERAL {
		table.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	join.Table = table

	if token, value := tokenizer.ReadToken(); token != lexing.ON {
		wrongTokenPanic(InnerJoinWithoutOnError, value)
	}

	if token, leftField := tokenizer.ReadToken(); token == lexing.LITERAL {
		join.Left = parseField(leftField)
	} else {
		wrongTokenPanic(InnerJoinWrongJoinFieldsError, leftField)
	}

	if token, operator := tokenizer.ReadToken(); token != lexing.EQUALS {
		wrongTokenPanic(InnerJoinWrongJoinFieldsError, operator)
	}

	if token, rightField := tokenizer.ReadToken(); token == lexing.LITERAL {
		join.Right = parseField(rightField)
	} else {
		wrongTokenPanic(InnerJoinWrongJoinFieldsError, rightField)
	}

	target.AddJoin(join)

	return result, true
}
