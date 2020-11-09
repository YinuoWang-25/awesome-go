package built_in_collections

import "testing"

// 0. declare slice
// Note: splice is a view of array
func TestBasicSlice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log("arr[2:6] =", arr[2:6])
	t.Log("arr[:6] =", arr[:6])
	t.Log("arr[2:] =", arr[2:])
	t.Log("arr[:] =", arr[:])
}

// 1. update original array
func updateSlice(s []int) {
	s[0] = 100
}
func TestUpdate(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:]
	t.Log("s1 =", s1)
	s2 := arr[:]
	t.Log("s2 =", s2)

	t.Log("After updateSlice(s1)")
	updateSlice(s1)
	t.Log("s1 =", s1)
	t.Log(arr)
	t.Log("s2 =", s2)
}

// 2. reslice
func TestReslice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:]
	t.Log("s1 =", s1)
	s2 := s1[2:]
	t.Log("s2 =", s2)
}

// 3. extendable
// Note: slice has len and cap, so if the upper bound index < cap, there
// is no problem, else there will be an error
func TestExtend(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	t.Log("s1 =", s1)
	s2 := s1[3:5]
	t.Log("s2 =", s2)
	// Note: arr
	//s3 := s1[3:7]
	//t.Log("s3 =", s3)
	t.Logf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
}

// 4. add element to slice
func TestAddElement(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	t.Logf("s3=%v, s4=%v, s5=%v\n", s3, s4, s5)
	t.Log(arr)
}
