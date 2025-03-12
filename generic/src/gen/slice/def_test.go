package slice_test

import "testing"

type SliceA interface{ []int | []int8 | []int64 }

type SliceB interface{ ~[]int | ~[]int8 | ~[]int64 }

type SliceC interface{ []*int | []*int8 | []*int64 }

type SliceD interface{ ~[]*int | ~[]*int8 | ~[]*int64 }

// type SliceE interface{ []~int | []~int8 | []~int64 } // ERROR: syntax: expected type, found '~'
// type SliceF interface{ ~[]~int | ~[]~int8 | ~[]~int64 }  // ERROR: syntax: expected type, found '~'

// type param 中定义 slice
type IntA interface{ int | int8 | int64 }

type IntB interface{ ~int | ~int8 | ~int64 }

// slice 一般定义在 Type param 中定义, 而不是直接用 interface 定义
// type SliceTP interface{ ~[]IntB } // ERROR: cannot use type IntB outside a type constraint

// 只能在 type param 中定义 ~[]~int, []~int
// ~[]~int | ~[]~int64
func sliceTP1[S ~[]E, E interface{ ~int | ~int64 }](S) {} // elem 是 ~int | ~int64, slice 是 ~[]~int | ~[]~int64

// []~int | []~int64
func sliceTP2[S []E, E interface{ ~int | ~int64 }](S) {} // elem 是 ~int | ~int64, slice 是 []~int | []~int64

// ~[]int | ~[]int64
func sliceArg1[T interface{ ~[]int | ~[]int64 }](T) {} // 相当于 func sliceArg[T SliceB](T) {}

func sliceArg2[S ~[]E, E interface{ int | int64 }](S) {}

func TestSliceArgs(*testing.T) {
	type myInt int
	sliceArg1([]int{1})
	// sliceArg([]myInt{1}) // ERROR: []myInt missing in ~[]int | ~[]int8 | ~[]int64

	sliceArg2([]int{1})
	// sliceArg2([]myInt{1}) // ERROR: myInt does not satisfy interface{int | int64}
}
