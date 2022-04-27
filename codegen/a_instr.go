package codegen

import (
	"assembler/parsing"
)

func CompileAInstruction(instr *parsing.AInstruction) string {
	intConstantStr := instr.Instr()[1:]
	bin := decToBin(intConstantStr)
	if binLen := len(bin); binLen < 16 {
		padding := 16 - binLen
		for padding > 0 {
			bin = "0" + bin
			padding--
		}
		// binLen >= 16
	} else {
		// take first 15 bits
		bin = bin[:15]
		// set first bit (opcode/MSB) to 0
		bin = "0" + bin
	}
	return bin
}
