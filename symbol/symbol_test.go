package symbol

import (
	"assembler/lexer"
	"assembler/token"
	"strings"
	"testing"
)

func TestResolve(t *testing.T) {
	l := lexer.New(`
		@10
			(OUTPUT_FIRST)
		@R11
		@LEX
		@ARG
		@SP
		@LCL
			(LEX)
		@THIS
		@THAT
		@SCREEN
		@KBD
		@OUTPUT_FIRST
		@R5
		D=M;JMP
	`)
	tt := []token.Token{}
	for {
		tok := l.Lex()
		tt = append(tt, tok)
		if tok.Tok == token.EOF {
			break
		}
	}
	want := `
		(A_INSTR, @10)
		(A_INSTR, @11)
		(A_INSTR, @7)
		(A_INSTR, @2)
		(A_INSTR, @0)
		(A_INSTR, @1)
		(A_INSTR, @3)
		(A_INSTR, @4)
		(A_INSTR, @16384)
		(A_INSTR, @24576)
		(A_INSTR, @1)
		(A_INSTR, @5)		
		(C_INSTR, D=M;JMP)
		`
	want = strings.Trim(want, " ")
	want = strings.ReplaceAll(want, "\t", "")
	want = strings.ReplaceAll(want, "\n", "")
	want = strings.ReplaceAll(want, " ", "")
	want = strings.ReplaceAll(want, ",", ", ")

	got := ""
	toks := Resolve(tt)

	for _, v := range toks {
		if v.Tok == token.EOF {
			break
		}
		got += v.String()
	}
	if got != want {
		t.Errorf("\nEXPECTED: '%s'\nGOT     : '%s'\n", want, got)
	}
}

func TestIsNotAllDigits(t *testing.T) {
	if isNotAllDigits("afaef1") == false {
		t.Errorf("1")
	}
	if isNotAllDigits("") == true {
		t.Errorf("2")
	}
	if isNotAllDigits("11111") == true {
		t.Errorf("3")
	}
}

func TestIsUpperCase(t *testing.T) {
	if isUpperCase("afagegaA") == true {
		t.Errorf("1")
	}
	if isUpperCase("AAAAAAAAAAAAA") == false {
		t.Errorf("2")
	}
	if isUpperCase("AZAASGAHHASWRH") == false {
		t.Errorf("3")
	}
	if isUpperCase("ASFAE1") == true {
		t.Errorf("4")
	}
	if isUpperCase("aAFGgA") == true {
		t.Errorf("5")
	}
	if isUpperCase("") == true {
		t.Errorf("6")
	}
}
