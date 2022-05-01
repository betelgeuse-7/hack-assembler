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
			!A
			D
			M=M+1;JNE`)
	want := "" +
		"(AT, @)" +
		"(NUMBER, 10)" +
		"(NEWLINE, \n)" +
		"(D, D)" +
		"(EQ, =)" +
		"(M, M)" +
		"(SEMICOLON, ;)" +
		"(JGE, JGE)" +
		"(NEWLINE, \n)" +
		"(LABEL, (OUTPUT_FIRST))" +
		"(NEWLINE, \n)" +
		"(AT, @)" +
		"(R11, R11)" +
		"(NEWLINE, \n)" +
		"(D, D)" +
		"(EQ, =)" +
		"(M, M)" +
		"(NEWLINE, \n)" +
		"(AT, @)" +
		"(SYMBOL, hello)" +
		"(NEWLINE, \n)" +
		"(D, D)" +
		"(SEMICOLON, ;)" +
		"(JMP, JMP)" +
		"(NEWLINE, \n)" +
		"(D, D)" +
		"(OR, |)" +
		"(M, M)" +
		"(NEWLINE, \n)" +
		"(AT, @)" +
		"(ARG, ARG)" +
		"(NEWLINE, \n)" +
		"(SYMBOL, e)" +
		"(NEWLINE, \n)" +
		"(BANG, !)" +
		"(A, A)" +
		"(NEWLINE, \n)" +
		"(D, D)" +
		"(NEWLINE, \n)" +
		"(M, M)" +
		"(EQ, =)" +
		"(M, M)" +
		"(PLUS, +)" +
		"(NUMBER, 1)" +
		"(SEMICOLON, ;)" +
		"(JNE, JNE)"

	want = strings.Trim(want, " ")
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
