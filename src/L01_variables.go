package main

import "fmt"

// 0. variable declared with default value
// Note: when printing empty string, use a combination of Printf and %q (as quote)
func variable() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

// 1. Declare variables with initial values
// Note: can declare multiple variables in one statement
// Note: no values needed
func variableWithInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println( a, b, s)
}

// 2. Declare variables with initial values
func variableTypeDeduction() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println( a, b, s)
}


func main() {
	variable()
	variableWithInitialValue()
	variableTypeDeduction()
}
