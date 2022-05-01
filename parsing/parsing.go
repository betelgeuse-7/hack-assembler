// produce instructions from a stream of tokens
package parsing

import (
	"assembler/token"
	"fmt"
	"strings"
)

type tokens []token.Token

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

type CInstruction struct {
	rawInstr string
	// populate these after parsing a C-instr
	dest, comp, jump string
}

func NewCInstruction(instr string) *CInstruction {
	return &CInstruction{
		rawInstr: instr,
	}
}

func (c *CInstruction) InstrType() instrType {
	return C_INSTR
}

func (c *CInstruction) Instr() string {
	return c.rawInstr
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

func (c *CInstruction) parseDest() {
	instr := c.rawInstr
	if strings.Contains(instr, "=") {
		c.dest = strings.Split(instr, "=")[0]
	}
}

func (c *CInstruction) parseJump() {
	instr := c.rawInstr
	if strings.Contains(instr, ";") {
		c.jump = strings.Split(instr, ";")[1]
	}
}

func (c *CInstruction) parseComp() {
	instr := c.rawInstr
	hasDest := strings.Contains(instr, "=")
	hasJump := strings.Contains(instr, ";")
	if hasDest && hasJump {
		destIdx := strings.Index(instr, "=")
		jumpIdx := strings.Index(instr, ";")
		c.comp = instr[destIdx+1 : jumpIdx]
	} else if hasDest {
		destIdx := strings.Index(instr, "=")
		c.comp = instr[destIdx+1:]
	} else if hasJump {
		jumpIdx := strings.Index(instr, ";")
		c.comp = instr[:jumpIdx]
	} else {
		c.comp = instr
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
	case token.A_INSTR:
		return p.parseAInstruction()
	default:
		return p.parseCInstruction()
	}
}

func (p *Parser) parseAInstruction() *AInstruction {
	instr := NewAInstruction(p.curTok.Lit)
	p.advance()
	return instr
}

func (p *Parser) parseCInstruction() *CInstruction {
	instr := NewCInstruction(p.curTok.Lit)
	instr.parseDest()
	instr.parseComp()
	instr.parseJump()
	p.advance()
	return instr
}
