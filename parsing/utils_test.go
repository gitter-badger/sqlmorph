package parsing_test

import (
	"reflect"
	"strings"
	"testing"

	. "github.com/s2gatev/sqlmorph/ast"
	. "github.com/s2gatev/sqlmorph/parsing"
)

type successTest struct {
	Query    string
	Expected Node
}

type errorTest struct {
	Query        string
	ErrorMessage string
}

func testPanic(t *testing.T, expectedMessage string) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic.")
	} else {
		actualMessage := r.(string)
		if expectedMessage != actualMessage {
			t.Errorf("The panic message was not correct.\n"+
				"\tExpected: %v\n"+
				"\tActual: %v\n", expectedMessage, actualMessage)
		}
	}
}

func parseQuery(query string) (Node, error) {
	return NewParser(strings.NewReader(query)).Parse()
}

func testParseCorrectQuery(t *testing.T,
	query string,
	expectedNode Node) {

	actualNode, _ := parseQuery(query)
	if !reflect.DeepEqual(expectedNode, actualNode) {
		t.Errorf("Parsed node is not correct.\n"+
			"Expected: %+v\n"+
			"Actual: %+v", expectedNode, actualNode)
	}
}

func runSuccessTests(t *testing.T, successTests []successTest) {
	execute := func(test successTest) {
		testParseCorrectQuery(t, test.Query, test.Expected)
	}

	for _, test := range successTests {
		execute(test)
	}
}

func runErrorTests(t *testing.T, errorTests []errorTest) {
	execute := func(test errorTest) {
		defer testPanic(t, test.ErrorMessage)

		parseQuery(test.Query)
	}

	for _, test := range errorTests {
		execute(test)
	}
}
