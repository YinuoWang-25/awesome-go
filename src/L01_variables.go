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
	fmt.Println(a, b, s)
}

// 2. Declare variables with initial values
// Note: in this way, variables in different types can declared in one statement
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 3. Declare variables in short
// Note: assignment can not use := again
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 4. Declare variables out of functions
// Note: always need var key word
// Note: scope: in package
var (
	k = "kk"
	bol = true
)

func main() {
	variable()
	variableWithInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(k, bol)
}
