package parsing

import (
	"github.com/s2gatev/sqlmorph/ast"
	"github.com/s2gatev/sqlmorph/lexing"
)

const (
	LeftWithoutJoinError         = "Expected JOIN following LEFT."
	LeftJoinWithoutTargetError   = "LEFT JOIN statement must be followed by a target class."
	LeftJoinWithoutOnError       = "LEFT JOIN statement must have an ON clause."
	LeftJoinWrongJoinFieldsError = "Wrong join fields in LEFT JOIN statement."
)

type LeftJoinState struct {
	BaseState
}

func (s *LeftJoinState) Name() string {
	return "LEFT JOIN"
}

func (s *LeftJoinState) Parse(result ast.Node, tokenizer *Tokenizer) (ast.Node, bool) {
	target := result.(ast.HasJoin)

	if token, _ := tokenizer.ReadToken(); token != lexing.LEFT {
		tokenizer.UnreadToken()
		return result, false
	}

	if token, value := tokenizer.ReadToken(); token != lexing.JOIN {
		wrongTokenPanic(LeftWithoutJoinError, value)
	}

	join := &ast.LeftJoin{}
	table := &ast.Table{}

	if token, tableName := tokenizer.ReadToken(); token == lexing.LITERAL {
		table.Name = tableName
	} else {
		wrongTokenPanic(LeftJoinWithoutTargetError, tableName)
	}

	if token, tableAlias := tokenizer.ReadToken(); token == lexing.LITERAL {
		table.Alias = tableAlias
	} else {
		tokenizer.UnreadToken()
	}

	join.Table = table

	if token, value := tokenizer.ReadToken(); token != lexing.ON {
		wrongTokenPanic(LeftJoinWithoutOnError, value)
	}

	if token, leftField := tokenizer.ReadToken(); token == lexing.LITERAL {
		join.Left = parseField(leftField)
	} else {
		wrongTokenPanic(LeftJoinWrongJoinFieldsError, leftField)
	}

	if token, operator := tokenizer.ReadToken(); token != lexing.EQUALS {
		wrongTokenPanic(LeftJoinWrongJoinFieldsError, operator)
	}

	if token, rightField := tokenizer.ReadToken(); token == lexing.LITERAL {
		join.Right = parseField(rightField)
	} else {
		wrongTokenPanic(LeftJoinWrongJoinFieldsError, rightField)
	}

	target.AddJoin(join)

	return result, true
}
