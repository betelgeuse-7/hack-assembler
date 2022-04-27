package lexer

import (
	"assembler/token"
	"strings"
	"testing"
)

func TestLexerLex(t *testing.T) {
	l := New(`@10
			D=M;JGE
		(OUTPUT_FIRST)
			@R11
			D=M
			@hello
			D;JMP 
			D|M 
			@ARG 
			e
			!A`)
	want := `
				(AT, @)
			    (NUMBER, 10)
				(D, D)
				(EQ, =)
				(M, M)
				(SEMICOLON, ;)
				(JGE, JGE)
				(LABEL, (OUTPUT_FIRST))
				(AT, @)
				(R11, R11)
				(D, D)
				(EQ, =)
				(M, M)
				(AT, @)
				(SYMBOL, hello)
				(D, D)
				(SEMICOLON, ;)
				(JMP, JMP)
				(D, D)
				(OR, |)
				(M, M)
				(AT, @)
				(ARG, ARG)
				(SYMBOL, e)
				(BANG, !)
				(A, A)
		`
	want = strings.Trim(want, " ")
	want = strings.ReplaceAll(want, "\n", "")
	want = strings.ReplaceAll(want, "\t", "")
	want = strings.ReplaceAll(want, " ", "")
	want = strings.ReplaceAll(want, ",", ", ")

	got := ""

	for {
		tok := l.Lex()
		if tok.Tok == token.EOF {
			break
		}
		got += tok.String()
	}
	if got != want {
		t.Errorf("\nEXPECTED: '%s'\nGOT     : '%s'\n", want, got)
	}
}
