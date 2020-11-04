package basic

import "testing"

// 0. pointer
// We can not do operation on pointer like C
func TestPointer (t *testing.T) {
	a := 2
	var pa *int = &a
	*pa = 3
	t.Log(a)
}