package parsing

import (
	"fmt"
	"strings"

	"github.com/s2gatev/sqlmorph/ast"
)

// wrongTokenPanic causes a panic because of an unexpected token.
func wrongTokenPanic(message string, value string) {
	if value != "" {
		message += fmt.Sprintf(" Found %s.", value)
	}
	panic(message)
}

// parseField parses field extracting its name and target.
func parseField(literal string) *ast.Field {
	field := &ast.Field{}
	literalParts := strings.Split(literal, ".")
	if len(literalParts) > 1 {
		field.Target = literalParts[0]
		field.Name = literalParts[1]
	} else {
		field.Name = literalParts[0]
	}
	return field
}
