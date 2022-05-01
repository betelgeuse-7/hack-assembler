package lexer

import (
	"assembler/token"
	"strings"
	"testing"
)

func TestLexerLex(t *testing.T) {
	l := New(`
			@10
			D=M;JGE
			(OUTPUT_FIRST)
			@R11
			D=M
			D
			@ARG
			!A
			M=M+1;JNE
			@hello
			1
			-1
		`)
	want := `
		(A_INSTR, @10)
		(C_INSTR, D=M;JGE)
		(LABEL, (OUTPUT_FIRST))
		(A_INSTR, @R11)
		(C_INSTR, D=M)
		(C_INSTR, D)
		(A_INSTR, @ARG)
		(C_INSTR, !A)
		(C_INSTR, M=M+1;JNE)
		(A_INSTR, @hello)
		(C_INSTR, 1)
		(C_INSTR, -1)
	`

	want = strings.Trim(want, " ")
	want = strings.ReplaceAll(want, "\t", "")
	want = strings.ReplaceAll(want, "\n", "")
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
