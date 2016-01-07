package parsing_test

import (
	"testing"
)

func TestMiscParsing(t *testing.T) {
	runErrorTests(t, []errorTest{
		{
			Query:        `INVALID FROM User u WHERE u.Name=?`,
			ErrorMessage: "Query is not complete. Expected SELECT, UPDATE or DELETE statement.",
		},
	})
}
