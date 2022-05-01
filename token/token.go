package token

type TokenType string

const (
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"

	A_INSTR TokenType = "A_INSTR"
	C_INSTR TokenType = "C_INSTR"

	// variables etc. @symbol
	SYMBOL TokenType = "SYMBOL"
	// goto labels
	LABEL TokenType = "LABEL"
)

type Token struct {
	Tok TokenType
	Lit string
}

func (t Token) String() string {
	return "(" + string(t.Tok) + ", " + t.Lit + ")"
}
