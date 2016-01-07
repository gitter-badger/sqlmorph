package parsing_test

import (
	"testing"

	. "github.com/s2gatev/sqlmorph/ast"
)

func TestSelectParsing(t *testing.T) {
	runSuccessTests(t, []successTest{
		{
			Query: `SELECT Name FROM User`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Name: "Name"},
				},
				Table: &Table{Name: "User"},
			},
		},
		{
			Query: `SELECT Name, Location, Age FROM User`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Name: "Name"}, &Field{Name: "Location"}, &Field{Name: "Age"},
				},
				Table: &Table{Name: "User"},
			},
		},
		{
			Query: `SELECT * FROM User`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Name: "*"},
				},
				Table: &Table{Name: "User"},
			},
		},
		{
			Query: `SELECT u.Name, u.Location, u.Age FROM User u`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
					&Field{Target: "u", Name: "Location"},
					&Field{Target: "u", Name: "Age"},
				},
				Table: &Table{Name: "User", Alias: "u"},
			},
		},
		{
			Query: `SELECT u.Name FROM User u WHERE u.Age=21`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
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
			Query: `SELECT u.Name, u.Location, u.Age FROM User u LIMIT 10`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
					&Field{Target: "u", Name: "Location"},
					&Field{Target: "u", Name: "Age"},
				},
				Limit: "10",
				Table: &Table{Name: "User", Alias: "u"},
			},
		},
		{
			Query: `SELECT u.Name, u.Location, u.Age FROM User u LIMIT 10 OFFSET 20`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
					&Field{Target: "u", Name: "Location"},
					&Field{Target: "u", Name: "Age"},
				},
				Limit:  "10",
				Offset: "20",
				Table:  &Table{Name: "User", Alias: "u"},
			},
		},
		{
			Query: `SELECT Name, Location FROM User INNER JOIN Address ON ID=UserID`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Name: "Name"},
					&Field{Name: "Location"},
				},
				Table: &Table{Name: "User"},
				JoinTables: []Join{
					&InnerJoin{
						Table: &Table{Name: "Address"},
						Left:  &Field{Name: "ID"},
						Right: &Field{Name: "UserID"},
					},
				},
			},
		},
		{
			Query: `SELECT u.Name, a.Location FROM User u INNER JOIN Address a ON u.ID=a.UserID`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
					&Field{Target: "a", Name: "Location"},
				},
				Table: &Table{Name: "User", Alias: "u"},
				JoinTables: []Join{
					&InnerJoin{
						Table: &Table{Name: "Address", Alias: "a"},
						Left:  &Field{Target: "u", Name: "ID"},
						Right: &Field{Target: "a", Name: "UserID"},
					},
				},
			},
		},
		{
			Query: `SELECT u.Name FROM User u CROSS JOIN Client`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
				},
				Table: &Table{Name: "User", Alias: "u"},
				JoinTables: []Join{
					&CrossJoin{
						Table: &Table{Name: "Client"},
					},
				},
			},
		},
		{
			Query: `SELECT u.Name, c.Name FROM User u CROSS JOIN Client c`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
					&Field{Target: "c", Name: "Name"},
				},
				Table: &Table{Name: "User", Alias: "u"},
				JoinTables: []Join{
					&CrossJoin{
						Table: &Table{Name: "Client", Alias: "c"},
					},
				},
			},
		},
		{
			Query: `SELECT Name, Location FROM User LEFT JOIN Address ON ID=UserID`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Name: "Name"},
					&Field{Name: "Location"},
				},
				Table: &Table{Name: "User"},
				JoinTables: []Join{
					&LeftJoin{
						Table: &Table{Name: "Address"},
						Left:  &Field{Name: "ID"},
						Right: &Field{Name: "UserID"},
					},
				},
			},
		},
		{
			Query: `SELECT u.Name, a.Location FROM User u LEFT JOIN Address a ON u.ID=a.UserID`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
					&Field{Target: "a", Name: "Location"},
				},
				Table: &Table{Name: "User", Alias: "u"},
				JoinTables: []Join{
					&LeftJoin{
						Table: &Table{Name: "Address", Alias: "a"},
						Left:  &Field{Target: "u", Name: "ID"},
						Right: &Field{Target: "a", Name: "UserID"},
					},
				},
			},
		},
		{
			Query: `SELECT Name, Location FROM User RIGHT JOIN Address ON ID=UserID`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Name: "Name"},
					&Field{Name: "Location"},
				},
				Table: &Table{Name: "User"},
				JoinTables: []Join{
					&RightJoin{
						Table: &Table{Name: "Address"},
						Left:  &Field{Name: "ID"},
						Right: &Field{Name: "UserID"},
					},
				},
			},
		},
		{
			Query: `SELECT u.Name, a.Location FROM User u RIGHT JOIN Address a ON u.ID=a.UserID`,
			Expected: &Select{
				Fields: []*Field{
					&Field{Target: "u", Name: "Name"},
					&Field{Target: "a", Name: "Location"},
				},
				Table: &Table{Name: "User", Alias: "u"},
				JoinTables: []Join{
					&RightJoin{
						Table: &Table{Name: "Address", Alias: "a"},
						Left:  &Field{Target: "u", Name: "ID"},
						Right: &Field{Target: "a", Name: "UserID"},
					},
				},
			},
		},
	})

	runErrorTests(t, []errorTest{
		{
			Query:        `SELECT u.Name, u.Location, u.Age FROM LIMIT 10`,
			ErrorMessage: "FROM statement must be followed by a target class. Found LIMIT.",
		},
		{
			Query:        `SELECT u.Name, u.Location, u.Age FROM User u LIMIT OFFSET`,
			ErrorMessage: "LIMIT statement must be followed by a number. Found OFFSET.",
		},
		{
			Query:        `SELECT u.Name, u.Location, u.Age FROM User u LIMIT 10 OFFSET`,
			ErrorMessage: "OFFSET statement must be followed by a number.",
		},
		{
			Query:        `SELECT FROM User`,
			ErrorMessage: "SELECT statement must be followed by field list. Found FROM.",
		},
		{
			Query:        `SELECT u.Name FROM User u WHERE LIMIT`,
			ErrorMessage: "WHERE statement must be followed by condition list. Found LIMIT.",
		},
		{
			Query:        `SELECT u.Name FROM User u WHERE a!=2 LIMIT`,
			ErrorMessage: "WHERE statement must be followed by condition list. Found !.",
		},
		{
			Query:        `SELECT u.Name FROM User u WHERE a=! LIMIT`,
			ErrorMessage: "WHERE statement must be followed by condition list. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u INNER Address a ON u.ID=a.UserID`,
			ErrorMessage: "Expected JOIN following INNER. Found Address.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u INNER JOIN ON u.ID=a.UserID`,
			ErrorMessage: "INNER JOIN statement must be followed by a target class. Found ON.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u INNER JOIN Address a`,
			ErrorMessage: "INNER JOIN statement must have an ON clause.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u INNER JOIN Address a ON !u=a.UserID`,
			ErrorMessage: "Wrong join fields in INNER JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u INNER JOIN Address a ON u.ID!=a.UserID`,
			ErrorMessage: "Wrong join fields in INNER JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u INNER JOIN Address a ON u.ID=!a.UserID`,
			ErrorMessage: "Wrong join fields in INNER JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, c.Name FROM User u CROSS Client c`,
			ErrorMessage: "Expected JOIN following CROSS. Found Client.",
		},
		{
			Query:        `SELECT u.Name, c.Name FROM User u CROSS JOIN`,
			ErrorMessage: "CROSS JOIN statement must be followed by a target class.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u LEFT Address a ON u.ID=a.UserID`,
			ErrorMessage: "Expected JOIN following LEFT. Found Address.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u LEFT JOIN ON u.ID=a.UserID`,
			ErrorMessage: "LEFT JOIN statement must be followed by a target class. Found ON.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u LEFT JOIN Address a`,
			ErrorMessage: "LEFT JOIN statement must have an ON clause.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u LEFT JOIN Address a ON !u=a.UserID`,
			ErrorMessage: "Wrong join fields in LEFT JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u LEFT JOIN Address a ON u.ID!=a.UserID`,
			ErrorMessage: "Wrong join fields in LEFT JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u LEFT JOIN Address a ON u.ID=!a.UserID`,
			ErrorMessage: "Wrong join fields in LEFT JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u RIGHT Address a ON u.ID=a.UserID`,
			ErrorMessage: "Expected JOIN following RIGHT. Found Address.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u RIGHT JOIN ON u.ID=a.UserID`,
			ErrorMessage: "RIGHT JOIN statement must be followed by a target class. Found ON.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u RIGHT JOIN Address a`,
			ErrorMessage: "RIGHT JOIN statement must have an ON clause.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u RIGHT JOIN Address a ON !u=a.UserID`,
			ErrorMessage: "Wrong join fields in RIGHT JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u RIGHT JOIN Address a ON u.ID!=a.UserID`,
			ErrorMessage: "Wrong join fields in RIGHT JOIN statement. Found !.",
		},
		{
			Query:        `SELECT u.Name, a.Location FROM User u RIGHT JOIN Address a ON u.ID=!a.UserID`,
			ErrorMessage: "Wrong join fields in RIGHT JOIN statement. Found !.",
		},
	})
}
