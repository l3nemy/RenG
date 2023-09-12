// Start Point
//     역할
//      1.실행시 메인 게임 엔진 메인화면 코드가 있는 파일 주소를 인터프리터한테 넘김
//      2.파일이 제대로 작동할 준비가 되었는지 확인함
//

/*
func main() {
	if runtime.GOOS == "windows" {
		root, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("core\\RenG", "-r", fmt.Sprintf("%s\\RenGLauncher", root))
		cmd.Run()
	}
}
*/

/*
func main() {
	if runtime.GOOS == "windows" {
		root, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("core\\RenG", "-r", fmt.Sprintf("%s\\game", root))
		cmd.Run()
	}
}

func main() {
	s := system.Init("RenG", 1280, 720, "core\\RenG\\cursor.png", nil,nil)
	s.WindowStart()
}
*/

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

	rf := file.CreateFile(os.Args[1])

	RgoCode := rf.Read()
	rf.CloseFile()

	l := lexer.New(RgoCode)
	p := parser.New(l)
	program := p.ParseProgram()
	fmt.Println(program.String())
	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			io.WriteString(os.Stdout, err+"\n\n")
		}
		return
	}

	comp := compiler.New()
	err := comp.CompileGlobal(program)
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
