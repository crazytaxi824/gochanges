package slice_test

import (
	"cmp"
	"slices"
	"sort"
	"testing"
)

// shallow clone, 只 assign slice 所有的 elem 到另个一 silce 中.
// 如果 elem 本身是引用类型, 则不会继续深入 clone.
func TestSliceClone(t *testing.T) {
	src := []int{1, 2, 3}
	c := slices.Clone(src)
	t.Log(src, c) // [1 2 3] [1 2 3]
	src[0] = 99
	t.Log(src, c) // [99 2 3] [1 2 3]

	// shallow copy slice 内部元素.
	src2 := [][]int{{1, 2, 3}, {4, 5, 6}}
	c2 := slices.Clone(src2)
	t.Log(src2, c2) // [[1 2 3] [4 5 6]] [[1 2 3] [4 5 6]]
	src2[0][0] = 99
	t.Log(src2, c2) // [[99 2 3] [4 5 6]] [[99 2 3] [4 5 6]]
}

// Search
// NOTE: The slice must be sorted in increasing order.
func TestSliceSearch(t *testing.T) {
	s := []int{1, 3, 5, 7}           // sorted
	t.Log(slices.BinarySearch(s, 5)) // 2, true
	t.Log(slices.BinarySearch(s, 9)) // 4, false

	slices.Reverse(s)                // not sorted
	t.Log(slices.BinarySearch(s, 5)) // 4,false
}

// Index
func TestSliceIndex(t *testing.T) {
	s := []int{1, 5, 3, 7}
	t.Log(slices.BinarySearch(s, 5)) // 3, false
	t.Log(slices.Index(s, 5))        // 1
}

// Equal
// slice elem 必须是 same order.
func TestSliceEqual(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 2, 1}
	s3 := []int{1, 2, 3}
	t.Log(slices.Equal(s1, s2)) // false
	t.Log(slices.Equal(s1, s3)) // true

	// ERROR: slice 内的元素必须是 comparable
	// t.Log(slices.Equal([][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 2, 3}, {4, 5, 6}}))
}

// Sort
func TestSliceSort(t *testing.T) {
	s := []int{3, 2, 4, 6, 5, 1}
	slices.Sort(s)
	t.Log(s)

	s1 := []int{3, 2, 4, 6, 5, 1}
	slices.SortFunc(s1, cmp.Compare)
	// slices.SortFunc(s1, func(a, b int) int {
	// 	return cmp.Compare(a, b)
	// })
	t.Log(s1)

	s2 := []int{3, 2, 4, 6, 5, 1}
	sort.Slice(s2, func(i, j int) bool {
		// return s1[i] < s1[j]
		return cmp.Less(s2[i], s2[j])
	})
	t.Log(s2)
}

// Insert, clone 到新的 slice
func TestSliceInsert(t *testing.T) {
	s := []int{1, 2, 3}
	s1 := slices.Insert(s, 3, 99)
	t.Log(s, s1) // [1 2 3] [1 2 3 99]

	s2 := slices.Insert(s, 0, 99)
	t.Log(s, s2) // [1 2 3] [99 1 2 3]

	// ERROR: out of range
	s3 := slices.Insert(s, 9, 99)
	t.Log(s, s3) // Panic

	// Insert 的时候 Clone 了 slice
	s[1] = 77
	t.Log(s, s1, s2) // [1 77 3] [1 2 3 99] [99 1 2 3]
}

// Delete, 修改原 slice 底层数据.
// delete 方式是将后面的 elem copy 到前面的 index 中, 然后更改 len & cap.
func TestSliceDelete(t *testing.T) {
	s1 := []int{0, 1, 2, 3, 4, 5}
	s2 := slices.Delete(s1, 3, 4) // delete 之后 s1 数据已经改变, 无法再使用.

	t.Log(s1, s2) // [0 1 2 4 5 5] [0 1 2 4 5]
	s1[0] = 99
	t.Log(s1, s2) // [99 1 2 4 5 5] [99 1 2 4 5]
}

// Replace, clone 到新的 slice
// replace 更像是 Insert + Delete
func TestSliceReplace(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}

	// delete [1,2) 的数据, 插入新的数据
	s1 := slices.Replace(s, 1, 2, 88, 99, 100)
	t.Log(s, s1) // [1 2 3 4 5 6] [1 88 99 100 3 4 5 6]

	s[0] = 77
	t.Log(s, s1) // [77 2 3 4 5 6] [1 88 99 100 3 4 5 6]
}
