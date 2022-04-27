package codegen

import (
	"assembler/parsing"
)

// https://raw.githubusercontent.com/aalhour/Assembler.hack/master/assets/c_instructions_reference.png

var compRefAIsZero = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"D+A": "000010",
	"D-A": "010011",
	"A-D": "000111",
	"D&A": "000000",
	"D|A": "010101",
}

var compRefAIsOne = map[string]string{
	"M":   "110000",
	"!M":  "110001",
	"-M":  "110011",
	"M+1": "110111",
	"M-1": "110010",
	"D+M": "000010",
	"D-M": "010011",
	"M-D": "000111",
	"D&M": "000000",
	"D|M": "010101",
}

var destRef = map[string]string{
	"null": "000",
	"M":    "001",
	"D":    "010",
	"MD":   "011",
	"A":    "100",
	"AM":   "101",
	"AD":   "110",
	"AMD":  "111",
}

var jumpRef = map[string]string{
	"null": "000",
	"JGT":  "001",
	"JEQ":  "010",
	"JGE":  "011",
	"JLT":  "100",
	"JNE":  "101",
	"JLE":  "110",
	"JMP":  "111",
}

func CompileCInstruction(instruction *parsing.CInstruction) string {
	// 1 1 1 a c1 c2 c3 c4 c5 c6 d1 d2 d3 j1 j2 j3
	bin := "111"
	comp := instruction.Comp()
	compRefVal, ok := compRefAIsZero[comp]
	if ok {
		bin += "0" + compRefVal
	}
	compRefVal, ok = compRefAIsOne[comp]
	if !ok {
		panic("invalid comp, " + comp + "\n")
	}
	bin += "1" + compRefVal
	// check if there is a dest
	if dest := instruction.Dest(); len(dest) != 0 {
		destRefVal, ok := destRef[dest]
		if !ok {
			panic("invalid dest, " + dest + "\n")
		}
		bin += destRefVal
	}
	// check jump
	if jump := instruction.Jump(); len(jump) != 0 {
		jumpRefVal, ok := jumpRef[jump]
		if !ok {
			panic("invalid jump, " + jump + "\n")
		}
		bin += jumpRefVal
	}
	// padding
	if binLen := len(bin); binLen < 16 {
		padding := 16 - binLen
		for padding > 0 {
			bin = bin + "0"
			padding--
		}
	}
	return bin
}
