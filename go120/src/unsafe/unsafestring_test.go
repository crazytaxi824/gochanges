// unsafe.String() 代替 reflect.StringHeader

package unsafe_test

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestStringHeader(t *testing.T) {
	s := struct {
		A, B, C, D, E, F byte
	}{97, 98, 99, 100, 101, 102} // abcdef

	sh := &reflect.StringHeader{
		Data: uintptr(unsafe.Pointer(&s.A)), // slice first elem address
		Len:  3,
	}

	s1 := *(*string)(unsafe.Pointer(sh))
	t.Log(s1) // abc
}

func TestStringHeader2(t *testing.T) {
	s := struct {
		A, B, C, D, E, F byte
	}{97, 98, 99, 100, 101, 102} // abcdef

	sh := &struct {
		Data uintptr
		Len  int
	}{
		Data: uintptr(unsafe.Pointer(&s.A)), // slice first elem address
		Len:  3,
	}

	s1 := *(*string)(unsafe.Pointer(sh))
	t.Log(s1) // abc
}

func TestStringHeader3(t *testing.T) {
	s := struct {
		A, B, C, D, E, F byte
	}{97, 98, 99, 100, 101, 102} // abcdef

	// 使用 array 代替 SliceHeader 结构
	sh := [2]uintptr{uintptr(unsafe.Pointer(&s.A)), 3}
	s1 := *(*string)(unsafe.Pointer(&sh))
	t.Log(s1) // abc
}

// unsafe.String()
func TestUnsafeString(t *testing.T) {
	b := [3]byte{97, 98, 99} // abc

	str := unsafe.String(&b[0], 3)
	t.Log(str) // abc

	b[0] = 100
	t.Log(str) // dbc, 所以是 unsafe string.
}

// unsafe.StringData()
func TestUnsafeStringData(t *testing.T) {
	b := []byte{97, 98, 99} // abc

	str := unsafe.String(&b[0], 3)
	t.Log(str) // abc

	t.Log(unsafe.StringData(str)) // underlying []bytes first elem address
	t.Logf("%p\n", &b[0])         // []bytes first elem address
	t.Logf("%p\n", b)             // []bytes first elem address
	t.Logf("%p\n", &str)          // StringHeader struct address
}
