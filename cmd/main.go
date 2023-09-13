package main

import (
	"internal/compiler/core/compiler"
	"internal/compiler/core/lexer"
	"internal/compiler/core/parser"
	"internal/compiler/file"
	"internal/compiler/vm"

	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [Rgo file]\n", os.Args[0])
		return
	}

	rf, err := file.CreateFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "File open error: %s\n", err)
		return
	}

	RgoCode := rf.Read()
	rf.CloseFile()

	l := lexer.New(RgoCode)
	p := parser.New(l)
	program := p.ParseProgram()
	fmt.Println(program.String())
	if len(p.Errors()) != 0 {
		for _, parse_err := range p.Errors() {
			_, err = io.WriteString(os.Stdout, parse_err+"\n\n")
			if err != nil {
				panic(err)
			}
		}
		return
	}

	comp := compiler.New()
	err = comp.CompileGlobal(program)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Compile failed:\n %s\n\n", err)
		return
	}

	err = comp.ReplaceSymbol()
	if err != nil {
		fmt.Println(err)
		return
	}

	// save it or execute it immediately
	machine := vm.New(comp.Bytecode())
	err = machine.Run()
	if err != nil {
		fmt.Fprintf(os.Stdout, "Bytecode execution failure:\n %s\n\n", err)
		return
	}
}
