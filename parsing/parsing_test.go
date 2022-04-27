package parsing

import (
	"assembler/token"
	"fmt"
	"testing"
)

func TestParserNext(t *testing.T) {
	input := tokens{
		{Tok: token.AT, Lit: "@"},
		{Tok: token.NUMBER, Lit: "16"},
		{Tok: token.D, Lit: "D"},
		{Tok: token.EQ, Lit: "="},
		{Tok: token.M, Lit: "M"},
		{Tok: token.SEMICOLON, Lit: ";"},
		{Tok: token.JMP, Lit: "JMP"},
	}
	p := New(input)
	for {
		if p.curTok.Tok == token.EOF {
			break
		}
		instr := p.Next()
		fmt.Printf("TYPE: %s\t INSTR: %s\n", instr.InstrType(), instr.Instr())
		if instr.InstrType() == C_INSTR {
			fmt.Printf("\t DEST: %s   COMP: %s   JUMP: %s\n", instr.(*CInstruction).dest, instr.(*CInstruction).comp, instr.(*CInstruction).jump)
		}
	}
}

func TestPeek(t *testing.T) {
	input := tokens{
		{Tok: token.AT, Lit: "@"},
		{Tok: token.NUMBER, Lit: "16"},
		{Tok: token.D, Lit: "D"},
		{Tok: token.EQ, Lit: "="},
		{Tok: token.M, Lit: "M"},
		{Tok: token.SEMICOLON, Lit: ";"},
		{Tok: token.JMP, Lit: "JMP"},
	}
	p := New(input)
	got := p.peek()
	want := token.NUMBER
	if got != want {
		t.Errorf("expected %s, got %s\n", want, got)
	}
	p.advance()
	p.advance()
	got = p.peek()
	want = token.EQ
	if got != want {
		t.Errorf("expected %s, got %s\n", want, got)
	}
	p.advance()
	p.advance()
	p.advance()
	p.advance()
	got = p.peek()
	want = token.EOF
	if got != want {
		t.Errorf("expected %s, got %s\n", want, got)
	}
}
