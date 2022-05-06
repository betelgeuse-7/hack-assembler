package lexer

import (
	"assembler/token"
)

type Lexer struct {
	input string
	pos   int
	ch    byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
		pos:   0,
	}
	l.ch = l.input[l.pos]
	return l
}

// TODO add symbol and label support
func (l *Lexer) Lex() token.Token {
	ch := l.ch
	if ch == 0 {
		return token.Token{
			Tok: token.EOF,
			Lit: "<<<EOF>>>",
		}
	}
	if isWhitespace(ch) {
		l.eatWhitespace()
		return l.Lex()
	}

	switch ch {
	case '@':
		return l.lexAInstr()
	case '(':
		return l.lexLabel()
	case 'A', 'M', 'D', '-', '!', '0', '1':
		return l.lexCInstr()
	case '/':
		l.eatComment()
		return l.Lex()
	}
	l.advance()
	return token.Token{
		Tok: token.ILLEGAL,
		Lit: string(ch),
	}
}

// jumps, symbol names, label names, etc.
func (l *Lexer) lexWord() token.Token {
	res := string(l.ch)
	l.advance()
	for isAsciiLetter(l.ch) {
		res += string(l.ch)
		l.advance()
	}
	return token.Token{
		Tok: token.SYMBOL,
		Lit: res,
	}
}

func (l *Lexer) lexLabel() token.Token {
	res := string(l.ch)
	l.advance()
	ch := l.ch
	labelName := ""
	if isAsciiLetter(ch) {
		labelName = l.lexWord().Lit
	}
	res += labelName
	res += string(l.ch)
	l.advance()
	return token.Token{
		Tok: token.LABEL,
		Lit: res,
	}
}

func (l *Lexer) lexAInstr() token.Token {
	res := string(l.ch)
	l.advance()
	for l.ch != '\n' && l.ch != ' ' && l.ch != '/' {
		res += string(l.ch)
		l.advance()
	}
	return token.Token{
		Tok: token.A_INSTR,
		Lit: res,
	}
}

func (l *Lexer) lexCInstr() token.Token {
	res := string(l.ch)
	l.advance()
	for l.ch != '\n' && l.ch != '@' && l.ch != ' ' && l.ch != '/' {
		res += string(l.ch)
		l.advance()
	}
	return token.Token{
		Tok: token.C_INSTR,
		Lit: res,
	}
}

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.ch) {
		l.advance()
	}
}

func (l *Lexer) eatComment() {
	for l.ch != '\n' {
		l.advance()
	}
}

func (l *Lexer) advance() {
	if l.pos == len(l.input)-1 {
		l.ch = 0
		return
	}
	l.pos++
	l.ch = l.input[l.pos]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAsciiLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}
