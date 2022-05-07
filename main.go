package main

import (
	"assembler/cmd"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		panic("you must provide a .asm file")
	}

	asmFile := os.Args[1]
	bx, err := os.ReadFile(asmFile)
	if err != nil {
		panic("error while reading the input file, " + err.Error() + "\n")
	}
	out, err := os.Create(strings.TrimSuffix(asmFile, ".asm") + ".hack")
	if err != nil {
		panic("error while creating the output file, " + err.Error() + "\n")
	}
	fmt.Println(string(bx))
	/*
		lexer needs a newline at the very end of the input.
		so i am appending it at the end.
		i don't want to fix it.
		this is not a bug. this is a feature. i think...
	*/
	n, err := cmd.Assemble(string(bx)+"\n", out)
	if err != nil {
		panic("an error occured, " + err.Error() + "\n")
	}
	log.Printf("[✓] Wrote %d bytes\n", n)
}
