package codegen

import (
	"assembler/lexer"
	"assembler/parsing"
	"fmt"
	"testing"
)

func TestCompileCInstruction(t *testing.T) {
	input := `
		D=M
		D;JGE
		A=D;JMP
		D
		A;JNE
		`
	want := []string{
		"1111110000010000",
		"1110001100000011",
		"1110001100100111",
		"1110001100000000",
		"1110110000000101",
	}
	l := lexer.New(input)
	p := parsing.NewWithLexer(l)
	fmt.Println("codegen: test: got parser")
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
		if bl := len(bin); bl != 16 {
			t.Errorf("bin length is not 16, but %d\n", bl)
		}
		i++
	}
}
