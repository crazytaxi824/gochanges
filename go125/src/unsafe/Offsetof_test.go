package unsafepkg

import (
	"testing"
	"unsafe"
)

// unsafe.Offsetof() 的核心用途是获取 struct 内字段的偏移量
func TestStructOffsetof(t *testing.T) {
	s2 := S2{
		a: 1,
		b: 2,
		c: 3,
	}

	// 找到各属性偏移量
	t.Log(unsafe.Offsetof(s2.a))
	t.Log(unsafe.Offsetof(s2.b))
	t.Log(unsafe.Offsetof(s2.c))

	// 获取各属性数据
	ptr := unsafe.Pointer(&s2) // Struct 内存起始位置
	pb := (*int32)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(s2.b)))
	t.Log(*pb)
}

func TestSliceOffsetof(t *testing.T) {
	s := []int{1, 2, 3}
	v := unsafe.Sizeof(s[0])

	ptr := unsafe.Pointer(&s[0]) // Slice 内存起始位置
	p := (*int)(unsafe.Pointer(uintptr(ptr) + v))
	t.Log(*p)
}
