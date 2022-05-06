package cmd

import (
	"assembler/codegen"
	"assembler/lexer"
	"assembler/parsing"
	"fmt"
	"io"
)

// return number of bytes written and any error encountered
func Assemble(input string, output io.Writer) (int, error) {
	bin := ""
	l := lexer.New(input)
	p := parsing.NewWithLexer(l)
	for {
		if p.CurTokIsEOF() {
			break
		}
		parsed := p.Next()
		switch parsed.InstrType() {
		case parsing.A_INSTR:
			bin += codegen.CompileAInstruction(parsed.(*parsing.AInstruction))
		case parsing.C_INSTR:
			bin += codegen.CompileCInstruction(parsed.(*parsing.CInstruction))
		default:
			panic("invalid instr '" + parsed.Instr() + "'  instr type -> " + string(parsed.InstrType()+"\n"))
		}
	}
	return fmt.Fprint(output, bin)
}
