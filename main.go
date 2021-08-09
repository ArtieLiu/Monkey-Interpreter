package main

import (
	"fmt"
	"monkeyinterpreter/lexer"
)

func main() {
	fmt.Println("hello")
	input := `123abc`
	l := lexer.New(input)

	fmt.Println(l)
}
