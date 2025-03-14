package main

import (
	"runtime"
	"slices"
	"testing"
	"time"
)

var a = []*string{nil, nil, nil}

func TestDeleteRef(t *testing.T) {
	var str = "hello world"
	runtime.AddCleanup(&str, func(*struct{}) {
		t.Log("str GCed")
	}, nil)

	a[2] = &str
	a[2] = nil // 如果 slice a 没有GC, 且没有移除对 str 的所有引用, 则 str 不会被GC.

	s := a[:2:2]
	t.Log(s, cap(s))

	runtime.GC()
	time.Sleep(time.Second)
}

func TestClone(t *testing.T) {
	s := make([]int, 0, 5)
	d := []int{1, 2, 3}

	s = slices.Clone(d)
	t.Log(s)
	t.Log(cap(s))
}

func TestFind(t *testing.T) {
	d := []int{2, 3, 5, 8, 10}
	t.Log(slices.IndexFunc(d, func(e int) bool {
		return 7 < e
	}))

	t.Log(slices.IndexFunc(d, func(e int) bool {
		return 10 < e
	}))

	t.Log(slices.IndexFunc(d, func(e int) bool {
		return 12 < e
	}))
}
