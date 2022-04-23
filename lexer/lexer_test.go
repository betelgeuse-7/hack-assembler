package lexer

import (
	"assembler/token"
	"strings"
	"testing"
)

func TestLexerLex(t *testing.T) {
	l := New("@10D=M;JGE")
	want := `
			(AT, @)
		    (NUMBER, 10)
			(D, D)
			(EQ, =)
			(M, M)
			(SEMICOLON, ;)
			(JGE, JGE)
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
