package main

import (
	"fmt"
	"math"
)

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

// 5. built-in data types
/*
- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
- byte, rune
- float32, float64, complex64, complex128
 */

// 6. data type conversion: need to do it explicitly
func triangle() {
	var a, b = 3, 4
	var c int = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println("triangle:", c)
}


func main() {
	variable()
	variableWithInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(k, bol)
	triangle();
}
