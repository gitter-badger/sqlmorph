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

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isSymbol(ch rune) bool {
	return (ch == '_' || ch == '.')
}
