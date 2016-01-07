package parsing

import (
	"fmt"
	"io"
	"strings"

	"github.com/s2gatev/sqlmorph/ast"
)

const QueryNotCompleteError = "Query is not complete. Expected %s statement."

// Parser parses SQL query into AST.
type Parser struct {
	tokenizer *Tokenizer
}

// NewParser creates a Parser instance for the provided query.
func NewParser(queryReader io.Reader) *Parser {
	return &Parser{tokenizer: NewTokenizer(queryReader)}
}

// Parse parses the query into a Node.
func (p *Parser) Parse() (ast.Node, error) {
	result := p.parseStates(
		nil,
		false,
		selectState(),
		updateState(),
		deleteState())

	return result, nil
}

func (p *Parser) parseStates(result ast.Node, isEndState bool, availableStates ...State) ast.Node {
	var stateNames []string
	for _, next := range availableStates {
		stateNames = append(stateNames, next.Name())
		if result, ok := next.Parse(result, p.tokenizer); ok {
			return p.parseStates(result, next.IsEndState(), next.Next()...)
		}
	}
	if isEndState {
		return result
	} else {
		statesList := stateNames[0]
		if len(stateNames) > 2 {
			statesList = strings.Join(stateNames[:len(stateNames)-1], ", ") + " or " + stateNames[len(stateNames)-1]
		}
		panic(fmt.Sprintf(QueryNotCompleteError, statesList))
	}
}

func selectState() *SelectState {
	offsetState := &OffsetState{}
	offsetState.MakeEndState()

	limitState := &LimitState{}
	limitState.MakeEndState()

	whereState := &WhereState{}
	whereState.MakeEndState()

	innerJoinState := &InnerJoinState{}
	innerJoinState.MakeEndState()

	leftJoinState := &LeftJoinState{}
	leftJoinState.MakeEndState()

	rightJoinState := &RightJoinState{}
	rightJoinState.MakeEndState()

	crossJoinState := &CrossJoinState{}
	crossJoinState.MakeEndState()

	fromState := &FromState{}
	fromState.MakeEndState()

	selectState := &SelectState{}

	selectState.SetNext(fromState)
	fromState.SetNext(leftJoinState, rightJoinState, innerJoinState, crossJoinState, whereState, limitState)
	innerJoinState.SetNext(whereState, limitState, leftJoinState, rightJoinState, innerJoinState, crossJoinState)
	leftJoinState.SetNext(whereState, limitState, leftJoinState, rightJoinState, innerJoinState, crossJoinState)
	rightJoinState.SetNext(whereState, limitState, leftJoinState, rightJoinState, innerJoinState, crossJoinState)
	crossJoinState.SetNext(whereState, limitState, leftJoinState, rightJoinState, innerJoinState, crossJoinState)
	whereState.SetNext(limitState)
	limitState.SetNext(offsetState)

	return selectState
}

func updateState() *UpdateState {
	whereState := &WhereState{}
	whereState.MakeEndState()

	setState := &SetState{}
	setState.SetNext(whereState)

	updateState := &UpdateState{}
	updateState.SetNext(setState)

	return updateState
}

func deleteState() *DeleteState {
	whereState := &WhereState{}
	whereState.MakeEndState()

	fromState := &FromState{}
	fromState.MakeEndState()
	fromState.SetNext(whereState)

	deleteState := &DeleteState{}
	deleteState.SetNext(fromState)

	return deleteState
}
