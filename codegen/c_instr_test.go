package codegen

import (
	"assembler/lexer"
	"assembler/parsing"
	"assembler/token"
	"fmt"
	"testing"
)

func TestCompileCInstruction(t *testing.T) {
	input := `
		D=M
		D;JGE
		A=D;JMP
		`
	want := []string{
		"1111110000010000",
		"1110001100000011",
		"1110001100100111",
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
		fmt.Println(parsed.(*parsing.CInstruction).String())
		bin := CompileCInstruction(parsed.(*parsing.CInstruction))
		if bin != want[i] {
			t.Errorf("at index %d, expected %s, but got %s\n", i, want[i], bin)
		}
		i++
	}
}
