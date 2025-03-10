// 约束中的并集 Union & 交集 Intersection

package gen_test

import "testing"

// 并集 Union
// type Union interface{ int | int64 | int32 | ... }

// 交集
// type Intersection interface { int; int64; int32; ...}

// 以下为示例:
// Union: Int --------------------------------------------------------------------------------------
type UnionA interface{ int | int8 | int64 }

type UnionB interface{ ~int | ~int8 | ~int64 }

type UnionC interface{ *int | *int8 | *int64 }

type UnionD interface{ ~*int | ~*int8 | ~*int64 }

// Intersection ------------------------------------------------------------------------------------
// int 和 int64 的交集, empty set.
type InterA interface {
	~int
	~int64
}

// 正确的使用方式:
type InterB interface {
	UnionA
	UnionB
}

func InterATest[T InterA](T) {}

func InterBTest[T InterB](T) {}

func TestIntersection(*testing.T) {
	// InterATest(1) // ERROR: cannot satisfy InterA (empty type set)
	InterBTest(1)
}

// 混合类型 ----------------------------------------------------------------------------------------
// 类型为并集, 同时必须有一个 Run() 的方法.
type MixA interface {
	~int | ~int8 | ~int64
	Run()
}

func mixTest[T MixA](T) {}

// 实现接口
type myMix int

func (myMix) Run() {}

func TestMix(*testing.T) {
	mixTest(myMix(0))
}
