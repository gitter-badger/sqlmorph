package parsing_test

import (
	"testing"

	. "github.com/s2gatev/sqlmorph/ast"
)

func TestUpdateParsing(t *testing.T) {
	runSuccessTests(t, []successTest{
		{
			Query: `UPDATE User u SET u.Name=? WHERE u.Age=21`,
			Expected: &Update{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name", Value: "?"},
				},
				Conditions: []*EqualsCondition{
					&EqualsCondition{
						Field: &Field{Target: "u", Name: "Age"},
						Value: "21",
					},
				},
				Table: &Table{Name: "User", Alias: "u"},
			},
		},
		{
			Query: `UPDATE User SET Name=? WHERE Age=21`,
			Expected: &Update{
				Fields: []*Field{
					&Field{Name: "Name", Value: "?"},
				},
				Conditions: []*EqualsCondition{
					&EqualsCondition{
						Field: &Field{Name: "Age"},
						Value: "21",
					},
				},
				Table: &Table{Name: "User"},
			},
		},
	})

	runErrorTests(t, []errorTest{
		{
			Query:        `UPDATE User u WHERE u.Age=21`,
			ErrorMessage: "Query is not complete. Expected SET statement.",
		},
		{
			Query:        `UPDATE User u SET`,
			ErrorMessage: "SET statement must be followed by field assignment list.",
		},
		{
			Query:        `UPDATE User u SET ! WHERE u.Age=21`,
			ErrorMessage: "SET statement must be followed by field assignment list. Found !.",
		},
		{
			Query:        `UPDATE User u SET u.Name!=? WHERE u.Age=21`,
			ErrorMessage: "SET statement must be followed by field assignment list. Found !.",
		},
		{
			Query:        `UPDATE User u SET u.Name=! WHERE u.Age=21`,
			ErrorMessage: "SET statement must be followed by field assignment list. Found !.",
		},
		{
			Query:        `UPDATE SET u.Name=? WHERE u.Age=21`,
			ErrorMessage: "UPDATE statement must be followed by a target class. Found SET.",
		},
	})
}
