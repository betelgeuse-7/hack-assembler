package symbol

import (
	"assembler/token"
	"fmt"
)

type tokens []token.Token

const (
	r0     = "@0"
	r1     = "@1"
	r2     = "@2"
	r3     = "@3"
	r4     = "@4"
	r5     = "@5"
	r6     = "@6"
	r7     = "@7"
	r8     = "@8"
	r9     = "@9"
	r10    = "@10"
	r11    = "@11"
	r12    = "@12"
	r13    = "@13"
	r14    = "@14"
	r15    = "@15"
	sp     = "@0"
	lcl    = "@1"
	arg    = "@2"
	this   = "@3"
	that   = "@4"
	screen = "@16384"
	kbd    = "@24576"
)

func Resolve(tokens tokens) tokens {
	tokens = resolveBuiltins(tokens)
	tokens = resolveLabels(tokens)
	//tokens = resolveVars(tokens)
	return tokens
}

func resolveLabels(tokens tokens) tokens {
	tt := []token.Token{}
	tokens, env := translateLabels(tokens)
	fmt.Println("resolveLabels -> env = ", env)
	fmt.Println("resolveLabels -> tt = ", tt)
	for _, tok := range tokens {
		if tok.Lit[0] == '@' && isUpperCase(tok.Lit) {
			// get rid of @ and add paranthesis
			key := fmt.Sprintf("(%s)", tok.Lit[1:])
			fmt.Println("key-> ", key)
			tok.Lit = "@" + env[key]
		}
		tt = append(tt, tok)
	}
	fmt.Println("tt.end -> ", tt)
	return tt
}

func translateLabels(tokens tokens) (tokens, map[string]string) {
	env := map[string]string{}
	tt := []token.Token{}
	for i, tok := range tokens {
		if tok.Tok == token.LABEL {
			env[tok.Lit] = fmt.Sprintf("%d", i)
			continue
		}
		// do not append label declarations
		tt = append(tt, tok)
	}
	return tt, env
}

func resolveBuiltins(tokens tokens) tokens {
	tt := []token.Token{}
	for _, tok := range tokens {
		switch tok.Lit {
		case "@R0":
			tok.Lit = r0
		case "@R1":
			tok.Lit = r1
		case "@R2":
			tok.Lit = r2
		case "@R3":
			tok.Lit = r3
		case "@R4":
			tok.Lit = r4
		case "@R5":
			tok.Lit = r5
		case "@R6":
			tok.Lit = r6
		case "@R7":
			tok.Lit = r7
		case "@R8":
			tok.Lit = r8
		case "@R9":
			tok.Lit = r9
		case "@R10":
			tok.Lit = r10
		case "@R11":
			tok.Lit = r11
		case "@R12":
			tok.Lit = r12
		case "@R13":
			tok.Lit = r12
		case "@R14":
			tok.Lit = r14
		case "@R15":
			tok.Lit = r15
		case "@SP":
			tok.Lit = sp
		case "@ARG":
			tok.Lit = arg
		case "@THIS":
			tok.Lit = this
		case "@THAT":
			tok.Lit = that
		case "@LCL":
			tok.Lit = lcl
		case "@SCREEN":
			tok.Lit = screen
		case "@KBD":
			tok.Lit = kbd
		}
		tt = append(tt, tok)
	}
	return tt
}

func isNotABuiltinVar(str string) bool {
	builtins := []string{
		"@R0", "@R1", "@R2", "@R3", "@R4", "@R5", "@R6", "@R7", "@R8", "@R9", "@R10", "@R11",
		"@R12", "@R13", "@R14", "@R15", "@SP", "@LCL", "@ARG", "@THIS", "@THAT", "@SCREEN", "@KBD",
	}
	for _, v := range builtins {
		if str == v {
			return false
		}
	}
	return true
}

func isAsciiLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isUpperCase(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, ch := range str {
		if !(('A' <= ch && ch <= 'Z') || ch == '_') {
			return false
		}
	}
	return true
}

func isNotAllDigits(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, ch := range str {
		if isAsciiLetter(byte(ch)) {
			return true
		}
	}
	return false
}
