package codegen

import (
	"assembler/lexer"
	"assembler/parsing"
	"assembler/token"
	"testing"
)

func TestNewAInstruction(t *testing.T) {
	input := `
		@21
		@16
	`
	want := []string{
		"0000000000010101",
		"0000000000010000",
	}
	l := lexer.New(input)
	var tokens []token.Token
	for {
		tok := l.Lex()
		tokens = append(tokens, tok)
		if tok.Tok == token.EOF {
			break
		}
	}
	p := parsing.New(tokens)
	i := 0
	for {
		if p.CurTokIsEOF() {
			break
		}
		parsed := p.Next()
		bin := CompileAInstruction(parsed.(*parsing.AInstruction))
		if bin != want[i] {
			t.Errorf("at index %d, expected %s, but got %s\n", i, want[i], bin)
		}
		i++
	}
}
