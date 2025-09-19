package unsafepkg

import (
	"testing"
	"unsafe"
)

type S1 struct {
	a byte
	b int32
	c int64
}

type S2 struct {
	b int32
	c int64
	a byte
}

// unsafe.Alignof() 显示内存最大对齐值
// unsafe.Sizeof()  显示内存占用数量
func TestAlignof(t *testing.T) {
	t.Log("Alignof(byte): ", unsafe.Alignof(byte(0)))  // 1
	t.Log("Alignof(int32):", unsafe.Alignof(int32(0))) // 4
	t.Log("Alignof(int64):", unsafe.Alignof(int64(0))) // 8
	t.Log("Alignof(S):", unsafe.Alignof(S1{}))         // 8 (取 struct 中最大对齐值)

	t.Log(unsafe.Sizeof(S1{})) // 16 struct 的整体内存占用 16 而不是 1+4+8=13
}

// 内存对齐方式
func TestAlignof2(t *testing.T) {
	// S1, S2 都是按照 8 来对齐内存, 但是因为顺序不同, 两个 struct 内存排列如下:
	// S1: [a,0,0,0,b,b,b,b, c,c,c,c,c,c,c,c]  总共占 16-Bytes
	// S2: [b,b,b,b,0,0,0,0, c,c,c,c,c,c,c,c, a,0,0,0,0,0,0,0,]  总共占 24-Bytes
	t.Log(unsafe.Alignof(S1{})) // 8
	t.Log(unsafe.Sizeof(S1{}))  // 16

	t.Log(unsafe.Alignof(S2{})) // 8
	t.Log(unsafe.Sizeof(S2{}))  // 24
}
