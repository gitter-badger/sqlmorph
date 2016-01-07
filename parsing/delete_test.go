package parsing_test

import (
	"testing"

	. "github.com/s2gatev/sqlmorph/ast"
)

func TestDeleteParsing(t *testing.T) {
	runSuccessTests(t, []successTest{
		{
			Query: `DELETE FROM User u WHERE u.Name=?`,
			Expected: &Delete{
				Conditions: []*EqualsCondition{
					&EqualsCondition{
						Field: &Field{Target: "u", Name: "Name"},
						Value: "?",
					},
				},
				Table: &Table{Name: "User", Alias: "u"},
			},
		},
	})

	runErrorTests(t, []errorTest{
		{
			Query:        `DELETE WHERE u.Name=?`,
			ErrorMessage: "Query is not complete. Expected FROM statement.",
		},
	})
}
