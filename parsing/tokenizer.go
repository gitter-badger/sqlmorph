package parsing

import (
	"io"

	"github.com/s2gatev/sqlmorph/lexing"
)

// Tokenizer reads SQL tokens from a query.
type Tokenizer struct {
	lexer  *lexing.Lexer
	buffer struct {
		token     lexing.Token
		value     string
		available bool
	}
}

// NewTokenizer creates Tokenizer instance that reads SQL tokens from the provided query.
func NewTokenizer(queryReader io.Reader) *Tokenizer {
	return &Tokenizer{lexer: lexing.NewLexer(queryReader)}
}

// ReadToken returns the next token from the SQL query ignoring whitespace.
func (t *Tokenizer) ReadToken() (lexing.Token, string) {
	token, value := t.read()
	if token == lexing.WHITESPACE {
		token, value = t.read()
	}
	return token, value
}

// UnreadToken brings back the previously read token into the SQL query.
func (t *Tokenizer) UnreadToken() {
	t.buffer.available = true
}

// read returns the next token in the SQL query.
func (t *Tokenizer) read() (lexing.Token, string) {
	if t.buffer.available {
		t.buffer.available = false
	} else {
		t.buffer.token, t.buffer.value = t.lexer.NextToken()
	}

	return t.buffer.token, t.buffer.value
}
