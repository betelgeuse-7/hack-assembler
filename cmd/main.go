package cmd

import (
	"assembler/codegen"
	"assembler/lexer"
	"assembler/parsing"
	"assembler/symbol"
	"assembler/token"
	"fmt"
	"io"
)

// return number of bytes written and any error encountered
func Assemble(input string, output io.Writer) (int, error) {
	bin := ""
	l := lexer.New(input)
	tt := []token.Token{}
	for {
		tok := l.Lex()
		if tok.Tok == token.EOF {
			tt = append(tt, tok)
			break
		}
		tt = append(tt, tok)
	}
	tt = symbol.Resolve(tt)
	p := parsing.New(tt)
	for {
		if p.CurTokIsEOF() {
			break
		}
		parsed := p.Next()
		switch parsed.InstrType() {
		case parsing.A_INSTR:
			bin += codegen.CompileAInstruction(parsed.(*parsing.AInstruction)) + "\n"
		case parsing.C_INSTR:
			bin += codegen.CompileCInstruction(parsed.(*parsing.CInstruction)) + "\n"
		default:
			panic("invalid instr '" + parsed.Instr() + "'  instr type -> " + string(parsed.InstrType()+"\n"))
		}
	}
	return fmt.Fprint(output, bin)
}
