package token

type TokenType string

const (
	EOF     TokenType = "EOF"
	ILLEGAL TokenType = "ILLEGAL"

	// address/data register
	A TokenType = "A"
	// main memory register
	M TokenType = "M"
	// data register
	D TokenType = "D"

	// jumps
	JGT TokenType = "JGT"
	JEQ TokenType = "JEQ"
	JGE TokenType = "JGE"
	JLT TokenType = "JLT"
	JNE TokenType = "JNE"
	JLE TokenType = "JLE"
	JMP TokenType = "JMP"

	// variables etc. @symbol
	SYMBOL TokenType = "SYMBOL"
	// goto labels
	LABEL  TokenType = "LABEL"
	NUMBER TokenType = "NUMBER"

	AT    TokenType = "AT"    // @
	EQ    TokenType = "EQ"    // =
	PLUS  TokenType = "PLUS"  // +
	MINUS TokenType = "MINUS" // -
	AND   TokenType = "AND"   // &
	OR    TokenType = "OR"    // |
	//LPAREN    TokenType = "LPAREN"    // (
	//RPAREN    TokenType = "RPAREN"    // )
	SEMICOLON TokenType = "SEMICOLON" // ;

	// ram address 0
	SP TokenType = "SP"
	// ram address 1
	LCL TokenType = "LCL"
	// ram address 2
	ARG TokenType = "ARG"
	// ram address 3
	THIS TokenType = "THIS"
	// ram addr 4
	THAT TokenType = "THAT"
	// base address of screen memory map (16384)
	// 16K bits
	SCREEN TokenType = "SCREEN"
	// base addr of keyboard memory map (24576)
	// this is 16 bits
	KBD TokenType = "KBD"

	// virtual registers
	R0  TokenType = "R0"
	R1  TokenType = "R1"
	R2  TokenType = "R2"
	R3  TokenType = "R3"
	R4  TokenType = "R4"
	R5  TokenType = "R5"
	R6  TokenType = "R6"
	R7  TokenType = "R7"
	R8  TokenType = "R8"
	R9  TokenType = "R9"
	R10 TokenType = "R10"
	R11 TokenType = "R11"
	R12 TokenType = "R12"
	R13 TokenType = "R13"
	R14 TokenType = "R14"
	R15 TokenType = "R15"
)

type Token struct {
	Tok TokenType
	Lit string
}

func (t Token) String() string {
	return "(" + string(t.Tok) + ", " + t.Lit + ")"
}
