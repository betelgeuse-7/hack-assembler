// produce instructions from a stream of tokens
package parsing

import (
	"assembler/token"
	"fmt"
)

type tokens []token.Token

func (t tokens) hasEq() (bool, int) {
	for i, v := range t {
		if v.Tok == token.EQ {
			return true, i
		}
	}
	return false, -1
}

// also return the index of semicolon, if it exists, otherwise, -1
func (t tokens) hasSemicolon() (bool, int) {
	for i, v := range t {
		if v.Tok == token.SEMICOLON {
			return true, i
		}
	}
	return false, -1
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
	tokens           tokens
}

func NewCInstruction() *CInstruction {
	return &CInstruction{}
}

func (c *CInstruction) InstrType() instrType {
	return C_INSTR
}

func (c *CInstruction) Instr() string {
	for _, v := range c.tokens {
		c.rawInstr += v.Lit
	}
	return c.rawInstr
}

func (c *CInstruction) appendToken(tok token.Token) {
	c.tokens = append(c.tokens, tok)
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

func (c *CInstruction) parseComp() {
	tokens := c.tokens
	if ok, ieq := tokens.hasEq(); ok {
		if ok, isem := tokens.hasSemicolon(); ok {
			for _, v := range tokens[ieq+1 : isem] {
				c.comp += v.Lit
			}
			return
		}
		for _, v := range tokens[ieq+1:] {
			c.comp += v.Lit
		}
		return
	}
	if ok, isem := tokens.hasSemicolon(); ok {
		for _, v := range tokens[:isem] {
			c.comp += v.Lit
		}
		return
	}
	for _, v := range tokens {
		c.comp += v.Lit
	}
}

func (c *CInstruction) parseDest() {
	tokens := c.tokens
	if ok, ieq := tokens.hasEq(); ok {
		for _, v := range tokens[:ieq] {
			c.dest += v.Lit
		}
	}
}

func (c *CInstruction) parseJump() {
	tokens := c.tokens
	if ok, isem := tokens.hasSemicolon(); ok {
		for _, v := range tokens[isem+1:] {
			c.jump += v.Lit
		}
	}
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

// return the next instruction
func (p *Parser) Next() Instruction {
	switch p.curTok.Tok {
	case token.AT:
		return p.parseAInstruction()
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

// bad code

func (p *Parser) parseCInstruction() *CInstruction {
	instr := NewCInstruction()
	instr.appendToken(p.curTok)
	for {
		p.advance()
		if p.peek() == token.AT {
			instr.appendToken(p.curTok)
			break
		}
		if p.curTok.Tok == token.SEMICOLON {
			instr.appendToken(p.curTok)
			p.advance()
			instr.appendToken(p.curTok)
			break
		} else if p.curTok.Tok == token.EOF {
			break
		}
		instr.appendToken(p.curTok)
	}
	instr.parseDest()
	instr.parseComp()
	instr.parseJump()
	p.advance()
	return instr
}
