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

	if isDigit(ch) {
		return l.lexNumber()
	} else if isAsciiLetter(ch) {
		return l.lexWord()
	} else if isWhitespace(ch) {
		l.eatWhitespace()
		return l.Lex()
	}

	switch ch {
	case '@':
		l.advance()
		return token.Token{
			Tok: token.AT,
			Lit: "@",
		}
	case '=':
		l.advance()
		return token.Token{
			Tok: token.EQ,
			Lit: "=",
		}
	case '+':
		l.advance()
		return token.Token{
			Tok: token.PLUS,
			Lit: "+",
		}
	case '-':
		l.advance()
		return token.Token{
			Tok: token.MINUS,
			Lit: "-",
		}
	case '&':
		l.advance()
		return token.Token{
			Tok: token.AND,
			Lit: "&",
		}
	case '|':
		l.advance()
		return token.Token{
			Tok: token.OR,
			Lit: "|",
		}
	case ';':
		l.advance()
		return token.Token{
			Tok: token.SEMICOLON,
			Lit: ";",
		}
	case '(':
		return l.lexLabel()
	}

	l.advance()
	return token.Token{
		Tok: token.ILLEGAL,
		Lit: string(ch),
	}
}

func (l *Lexer) lexNumber() token.Token {
	res := string(l.ch)
	l.advance()
	for isDigit(l.ch) {
		res += string(l.ch)
		l.advance()
	}
	return token.Token{
		Tok: token.NUMBER,
		Lit: res,
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

	switch res {
	case "A":
		return token.Token{
			Tok: token.A,
			Lit: "A",
		}
	case "M":
		return token.Token{
			Tok: token.M,
			Lit: "M",
		}
	case "D":
		return token.Token{
			Tok: token.D,
			Lit: "D",
		}
	case "JGT":
		return token.Token{
			Tok: token.JGT,
			Lit: res,
		}
	case "JEQ":
		return token.Token{
			Tok: token.JEQ,
			Lit: res,
		}
	case "JGE":
		return token.Token{
			Tok: token.JGE,
			Lit: res,
		}
	case "JLT":
		return token.Token{
			Tok: token.JLT,
			Lit: res,
		}
	case "JNE":
		return token.Token{
			Tok: token.JNE,
			Lit: res,
		}
	case "JLE":
		return token.Token{
			Tok: token.JLE,
			Lit: res,
		}
	case "JMP":
		return token.Token{
			Tok: token.JMP,
			Lit: res,
		}
	case "SP":
		return token.Token{
			Tok: token.SP,
			Lit: res,
		}
	case "LCL":
		return token.Token{
			Tok: token.LCL,
			Lit: res,
		}
	case "ARG":
		return token.Token{
			Tok: token.ARG,
			Lit: res,
		}
	case "THIS":
		return token.Token{
			Tok: token.THIS,
			Lit: res,
		}
	case "THAT":
		return token.Token{
			Tok: token.THAT,
			Lit: res,
		}
	case "SCREEN":
		return token.Token{
			Tok: token.SCREEN,
			Lit: res,
		}
	case "KBD":
		return token.Token{
			Tok: token.KBD,
			Lit: res,
		}
	case "R":
		if isDigit(l.ch) {
			lexed := l.lexVirtualRegister()
			return token.Token{
				Tok: lexed.Tok,
				Lit: res + lexed.Lit,
			}
		}
		return token.Token{
			Tok: token.ILLEGAL,
			Lit: res,
		}
	}

	return token.Token{
		Tok: token.SYMBOL,
		Lit: res,
	}
}

func (l *Lexer) lexVirtualRegister() (tok token.Token) {
	res := string(l.ch)
	l.advance()
	ch := l.ch
	if isDigit(ch) {
		res += string(ch)
		l.advance()
	}
	switch res {
	case "0":
		tok = token.Token{
			Tok: token.R0,
			Lit: res,
		}
	case "1":
		tok = token.Token{
			Tok: token.R1,
			Lit: res,
		}
	case "2":
		tok = token.Token{
			Tok: token.R2,
			Lit: res,
		}
	case "3":
		tok = token.Token{
			Tok: token.R3,
			Lit: res,
		}
	case "4":
		tok = token.Token{
			Tok: token.R4,
			Lit: res,
		}
	case "5":
		tok = token.Token{
			Tok: token.R5,
			Lit: res,
		}
	case "6":
		tok = token.Token{
			Tok: token.R6,
			Lit: res,
		}
	case "7":
		tok = token.Token{
			Tok: token.R7,
			Lit: res,
		}
	case "8":
		tok = token.Token{
			Tok: token.R8,
			Lit: res,
		}
	case "9":
		tok = token.Token{
			Tok: token.R9,
			Lit: res,
		}
	case "10":
		tok = token.Token{
			Tok: token.R10,
			Lit: res,
		}
	case "11":
		tok = token.Token{
			Tok: token.R11,
			Lit: res,
		}
	case "12":
		tok = token.Token{
			Tok: token.R12,
			Lit: res,
		}
	case "13":
		tok = token.Token{
			Tok: token.R13,
			Lit: res,
		}
	case "14":
		tok = token.Token{
			Tok: token.R14,
			Lit: res,
		}
	case "15":
		tok = token.Token{
			Tok: token.R15,
			Lit: res,
		}
	}
	return
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

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.ch) {
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
	return ch == ' ' || ch == '\n' || ch == '\t' || ch == '\r'
}
