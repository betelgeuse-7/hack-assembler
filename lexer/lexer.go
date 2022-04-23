package lexer

import (
	"assembler/token"
	"fmt"
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
	}

	switch ch {
	case 'A':
		l.advance()
		return token.Token{
			Tok: token.A,
			Lit: "A",
		}
	case 'M':
		l.advance()
		return token.Token{
			Tok: token.M,
			Lit: "M",
		}
	case 'D':
		l.advance()
		return token.Token{
			Tok: token.D,
			Lit: "D",
		}
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
	}

	// ???
	if ch == 'J' {
		l.lexJump()
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

func (l *Lexer) lexJump() token.Token {
	res := string(l.ch)
	l.advance()
	// second char must be G, E, L, N, or M
	secondCh := l.ch
	if secondCh == 'G' || secondCh == 'E' || secondCh == 'L' || secondCh == 'N' || secondCh == 'M' {
		res += string(secondCh)
		l.advance()
		// third char must be either T, Q, E, or P
		thirdCh := l.ch
		if thirdCh == 'T' || thirdCh == 'Q' || thirdCh == 'E' || thirdCh == 'P' {
			res += string(thirdCh)
			fmt.Println("geasg", res)
			l.advance()
			switch res {
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
			}
		}
	}
	// ! return illegal for now
	return token.Token{
		Tok: token.ILLEGAL,
		Lit: res,
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
