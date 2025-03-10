// unsafe.Slice() 代替 reflect.SliceHeader

package unsafe_test

import (
	"reflect"
	"testing"
	"unsafe"
)

// SliceHeader 用法, Deprecated
func TestSliceHeader(t *testing.T) {
	s := struct {
		A, B, C, D, E int
	}{3, 2, 1, 4, 6}

	sh := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&s.A)), // slice first elem address
		Len:  3,
		Cap:  3,
	}

	s1 := *(*[]int)(unsafe.Pointer(sh))
	t.Log(s1) // [3 2 1]
}

// 可以自定义 SliceHeader 结构体, 虽然 reflect.SliceHeader 被弃用, 但是 Slice 本身的数据结构没有变.
func TestSliceHeader2(t *testing.T) {
	s := struct {
		A, B, C, D, E int
	}{3, 2, 1, 4, 6}

	// 变相替代 SliceHeader
	sh := &struct {
		D uintptr
		L int
		C int
	}{
		D: uintptr(unsafe.Pointer(&s.A)),
		L: 3,
		C: 3,
	}

	s1 := *(*[]int)(unsafe.Pointer(sh))
	t.Log(s1) // [3 2 1]
}

// 简化 SliceHeader, 将 struct => array.
func TestSliceHeader3(t *testing.T) {
	s := struct {
		A, B, C, D, E int
	}{3, 2, 1, 4, 6}

	// 使用 array 代替 SliceHeader 结构
	sh := [3]uintptr{uintptr(unsafe.Pointer(&s.A)), 3, 3}
	s1 := *(*[]int)(unsafe.Pointer(&sh))
	t.Log(s1) // [3 2 1]
}

// unsafe.Slice() 替代 *(*[]int)(unsafe.Pointer(&sh))
// func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
func TestUnsafeSlice(t *testing.T) {
	s := struct {
		A, B, C, D, E int
	}{3, 2, 1, 4, 6}
	s1 := unsafe.Slice(&s.A, 3) // slice first elem address
	t.Log(s1)                   // [3 2 1]
}

// unsafe.SliceData()
func TestUnsafeSliceData(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log(unsafe.SliceData(s)) // slice first elem address 替代 reflect.SliceHeader.Data
	t.Logf("%p\n", s)          // slice first elem address
	t.Logf("%p\n", &s[0])      // slice first elem address
	t.Logf("%p\n", &s)         // SliceHeader struct address
}
