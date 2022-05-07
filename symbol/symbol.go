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

// resolve all symbols, variables to a bunch of a-instructions with constant decimal values
func Resolve(tokens tokens) tokens {
	tokens = resolveBuiltins(tokens)
	tokens = resolveLabels(tokens)
	tokens = resolveVars(tokens)
	return tokens
}

// substitute label references (e.g @LOOP), with their constant values using the env map returned from translateLabels
func resolveLabels(tokens tokens) tokens {
	tt := []token.Token{}
	tokens, env := translateLabels(tokens)
	for _, tok := range tokens {
		if isNotABuiltinVar(tok.Lit) && tok.Lit[0] == '@' && isUpperCase(tok.Lit[1:]) {
			// get rid of @ and add paranthesis
			key := fmt.Sprintf("(%s)", tok.Lit[1:])
			tok.Lit = "@" + env[key]
		}
		tt = append(tt, tok)
	}
	return tt
}

// return tokens without label declarations, and a map that stores values matching label declarations
// with their constant values.
// 	e.g
//	...
//	5   (LOOP)   // this is stored in env (map) like that: (LOOP): 5
//	6	0;JMP
//	7	@LOOP	 // in resolveLabels, this is translated to @5, using the env map returned from this function
//	8
//	...
func translateLabels(tokens tokens) (tokens, map[string]string) {
	env := map[string]string{}
	tt := []token.Token{}
	for i, tok := range tokens {
		if tok.Tok == token.LABEL {
			i -= len(env)
			env[tok.Lit] = fmt.Sprintf("%d", i)
			continue
		}
		// do not append label declarations
		tt = append(tt, tok)
	}
	return tt, env
}

// substitute built-in symbols with a-instructions that have constant decimal values
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

// substitute variables with their values using env map
// just like resolveLabels
func resolveVars(tokens tokens) tokens {
	tt := []token.Token{}
	tokens, env := translateVars(tokens)
	for _, tok := range tokens {
		key := tok.Lit[1:]
		if tok.Lit[0] == '@' && !isUpperCase(key) && isNotAllDigits(tok.Lit) {
			tok.Lit = "@" + env[key]
		}
		tt = append(tt, tok)
	}
	return tt
}

// very similar to translateLabels, but this is giving user-defined variables a constant decimal value
// first variable gets 16, the second gets 17, and so on...
// when we first encounter a new variable (user-defined variables must be not-all-uppercase, and they must include at least one ASCII character)
// , we set a key in env, with its constant value.
//
// returns tokens, and the env map
func translateVars(tokens tokens) (tokens, map[string]string) {
	env := map[string]string{}
	tt := []token.Token{}
	for _, tok := range tokens {
		key := tok.Lit[1:]
		if _, alreadyThere := env[key]; !alreadyThere {
			if tok.Tok == token.A_INSTR && !isUpperCase(tok.Lit) {
				env[key] = fmt.Sprintf("%d", 16+len(env))
			}
		}
		tt = append(tt, tok)
	}
	return tt, env
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
