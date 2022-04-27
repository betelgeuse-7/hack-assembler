// produce instructions from a stream of tokens
package parsing

import (
	"assembler/token"
	"fmt"
)

type tokens []token.Token

func (t tokens) hasEq() (int, bool) {
	for i, v := range t {
		if v.Tok == token.EQ {
			return i, true
		}
	}
	return -1, false
}

// also return the index of semicolon, if it exists, otherwise, -1
func (t tokens) hasSemicolon() (int, bool) {
	for i, v := range t {
		if v.Tok == token.SEMICOLON {
			return i, true
		}
	}
	return -1, false
}

type instrType string

const (
	A_INSTR instrType = "A_INSTR"
	C_INSTR instrType = "C_INSTR"
)

type Instruction interface {
	InstrType() instrType
	Instr() string
}

type AInstruction struct {
	rawInstr string
}

func NewAInstruction(rawInstr string) *AInstruction {
	return &AInstruction{
		rawInstr: rawInstr,
	}
}

func (a *AInstruction) InstrType() instrType {
	return A_INSTR
}

func (a *AInstruction) Instr() string {
	return a.rawInstr
}

// append a new string to a.rawInstr
func (a *AInstruction) appendRawInstr(instr string) {
	a.rawInstr += instr
}

type CInstruction struct {
	rawInstr string
	// populate these after parsing a C-instr
	dest, comp, jump string
}

func NewCInstruction(rawInstr string) *CInstruction {
	return &CInstruction{
		rawInstr: rawInstr,
	}
}

func (c *CInstruction) InstrType() instrType {
	return C_INSTR
}

func (c *CInstruction) Instr() string {
	return c.rawInstr
}

func (c *CInstruction) appendRawInstr(instr string) {
	c.rawInstr += instr
}

func (c *CInstruction) Dest() string {
	return c.dest
}

func (c *CInstruction) Comp() string {
	return c.comp
}

func (c *CInstruction) Jump() string {
	return c.jump
}

func (c *CInstruction) String() string {
	return fmt.Sprintf("(dest: %s, comp: %s, jump: %s)", c.dest, c.comp, c.jump)
}

type Parser struct {
	tokenStream tokens
	pos         int
	curTok      token.Token
}

func New(tokenStream tokens) *Parser {
	p := &Parser{
		tokenStream: tokenStream,
		pos:         0,
	}
	p.curTok = tokenStream[p.pos]
	return p
}

func (p *Parser) CurTokIsEOF() bool {
	return p.curTok.Tok == token.EOF
}

// p.pos++, p.curTok = p.tokenStream[p.pos]
func (p *Parser) advance() {
	if ahead := p.peek(); ahead == token.EOF {
		p.curTok.Tok = ahead
		p.curTok.Lit = "<<<EOF>>>"
		return
	}
	p.pos++
	p.curTok = p.tokenStream[p.pos]
}

func (p *Parser) peek() token.TokenType {
	if p.pos == len(p.tokenStream)-1 {
		return token.EOF
	}
	return p.tokenStream[p.pos+1].Tok
}

func (p *Parser) peekAssert(expected token.TokenType) bool {
	if ahead := p.peek(); ahead != expected {
		return false
	}
	return true
}

// return the next instruction
func (p *Parser) Next() Instruction {
	switch p.curTok.Tok {
	case token.AT:
		return p.parseAInstruction()
		// C-instr
	default:
		return p.parseCInstruction()
	}
}

func (p *Parser) parseAInstruction() *AInstruction {
	// since we assume that the given .asm file is error-free
	// we don't assert that we have a valid integer constant after token.AT
	instr := NewAInstruction(p.curTok.Lit)
	p.advance()
	instr.appendRawInstr(p.curTok.Lit)
	// we parsed an A-instruction
	p.advance()
	return instr
}

// TODO FIX THIS
func (p *Parser) parseCInstruction() *CInstruction {
	instr := NewCInstruction(p.curTok.Lit)
	var cInstrTokens tokens
	// go until token.AT (A instructions start with token.AT)
	for {
		if p.curTok.Tok == token.EOF || p.curTok.Tok == token.AT {
			break
		}
		cInstrTokens = append(cInstrTokens, p.curTok)
		p.advance()
	}
	// do we have a dest?
	if _, ok := cInstrTokens.hasEq(); ok {
		var destTokens tokens
		i := 0
		for {
			curCInstrTok := cInstrTokens[i]
			if curCInstrTok.Tok == token.EQ {
				break
			}
			destTokens = append(destTokens, curCInstrTok)
			i++
		}
		dest := ""
		for _, v := range destTokens {
			dest += v.Lit
		}
		instr.dest = dest
	}
	// do we have a jump
	if index, ok := cInstrTokens.hasSemicolon(); ok {
		index++
		curCInstrTok := cInstrTokens[index]
		instr.jump = curCInstrTok.Lit
	}
	// * set comp
	comp := ""
	if ieq, ok := cInstrTokens.hasEq(); ok {
		if isem, ok := cInstrTokens.hasSemicolon(); ok {
			compTokens := cInstrTokens[ieq+1 : isem]
			for _, v := range compTokens {
				comp += v.Lit
			}
		}
		compTokens := cInstrTokens[ieq+1:]
		for _, v := range compTokens {
			comp += v.Lit
		}
	}
	if isem, ok := cInstrTokens.hasSemicolon(); ok {
		compTokens := cInstrTokens[0:isem]
		for _, v := range compTokens {
			comp += v.Lit
		}
	}
	instr.comp = comp
	return instr
}
