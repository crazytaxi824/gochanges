package slice_test

import "testing"

// 研究 Slice type params 定义的区别
func S1[S ~[]E, E interface{ ~int | ~int64 }](S) {} // elem 是 ~int | ~int64, slice 是 ~[]~int | ~[]~int64

func S2[S ~[]E, E interface{ int | int64 }](S) {} // elem 是 int | int64, slice 是 ~[]int | ~[]int64

func S3[S []E, E interface{ ~int | ~int64 }](S) {} // elem 是 ~int | ~int64, slice 是 []~int | []~int64

func S4[S []E, E interface{ int | int64 }](S) {} // elem 是 int | int64, slice 是 []int | []int64

// 实现
type myInt int

type mySliceInt []int

type mySliceMyInt []myInt

func TestSliceTypeParam3(*testing.T) {
	S1([]int{1})        // []int{}
	S1([]myInt{1})      // []~int{}
	S1(mySliceInt{1})   // ~[]int{}
	S1(mySliceMyInt{1}) // ~[]~int{}

	S2([]int{1}) // []int{}
	// S2([]myInt{1})      // []~int{}, ERROR: myInt does not satisfy interface{int | int64}
	S2(mySliceInt{1}) // ~[]int{}
	// S2(mySliceMyInt{1}) // ~[]~int{}, ERROR: myInt does not satisfy interface{int | int64}

	S3([]int{1})   // []int{}
	S3([]myInt{1}) // []~int{}
	// S3(mySliceInt{1})   // ~[]int{}, ERROR: mySliceInt does not satisfy []int
	// S3(mySliceMyInt{1}) // ~[]~int{}, ERROR: mySliceMyInt does not satisfy []myInt

	S4([]int{1}) // []int{}
	// S4([]myInt{1})      // []~int{}  // ERROR: myInt does not satisfy interface{int | int64}
	// S4(mySliceInt{1})   // ~[]int{}  // ERROR: mySliceInt does not satisfy []int
	// S4(mySliceMyInt{1}) // ~[]~int{} // ERROR: mySliceMyInt does not satisfy []myInt
}
