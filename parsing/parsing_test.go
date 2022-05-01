package parsing

import (
	"assembler/lexer"
	"assembler/token"
	"testing"
)

func TestParserNext(t *testing.T) {
	input := `
		@16
		D=A;JEQ
		D=M;JMP
		D=D+1
		@61
		D
		M=M+1;JNE
		@128
		D=!D
	`
	l := lexer.New(input)
	lexed := tokens{}
	for {
		tok := l.Lex()
		if tok.Tok == token.EOF {
			break
		}
		lexed = append(lexed, tok)
	}
	p := New(lexed)

	want := []Instruction{
		&AInstruction{rawInstr: "@16"},
		&CInstruction{rawInstr: "D=A;JEQ", dest: "D", comp: "A", jump: "JEQ"},
		&CInstruction{rawInstr: "D=M;JMP", dest: "D", comp: "M", jump: "JMP"},
		&CInstruction{rawInstr: "D=D+1", dest: "D", comp: "D+1", jump: ""},
		&AInstruction{rawInstr: "@61"},
		&CInstruction{rawInstr: "D", dest: "", comp: "D", jump: ""},
		&CInstruction{rawInstr: "M=M+1;JNE", dest: "M", comp: "M+1", jump: "JNE"},
		&AInstruction{rawInstr: "@128"},
		&CInstruction{rawInstr: "D=!D", dest: "D", comp: "!D", jump: ""},
	}

	i := 0
	for {
		if p.curTok.Tok == token.EOF {
			break
		}
		instr := p.Next()
		if instr.InstrType() == A_INSTR {
			curWantInstr := want[i].(*AInstruction)
			curGotInstr := instr.(*AInstruction)
			if wantInstr := curWantInstr.rawInstr; wantInstr != curGotInstr.rawInstr {
				t.Errorf("*AInstruction.rawInstr not the same. want %s, got %s\n", wantInstr, curGotInstr.rawInstr)
			}
		}
		if instr.InstrType() == C_INSTR {
			curWantInstr := want[i].(*CInstruction)
			curGotInstr := instr.(*CInstruction)
			if wantInstr := curWantInstr.rawInstr; wantInstr != curGotInstr.rawInstr {
				t.Errorf("*CInstruction.rawInstr not the same. want %s, got %s\n", wantInstr, curGotInstr.rawInstr)
			}
			if wantDest := curWantInstr.dest; wantDest != curGotInstr.dest {
				t.Errorf("*CInstruction.dest not the same. want %s, got %s\n", wantDest, curGotInstr.dest)
			}
			if wantComp := curWantInstr.comp; wantComp != curGotInstr.comp {
				t.Errorf("*CInstruction.comp not the same. want %s, got %s\n", wantComp, curGotInstr.comp)
			}
			if wantJump := curWantInstr.jump; wantJump != curGotInstr.jump {
				t.Errorf("*CInstruction.jump not the same. want %s, got %s\n", wantJump, curGotInstr.jump)
			}
		}
		i++
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
