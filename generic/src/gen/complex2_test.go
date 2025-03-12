package gen_test

import "testing"

type Op[T any] interface {
	Set(T)
	Get() T
}

// 实现 Op
type MyStr string

func (s MyStr) Get() MyStr {
	return s
}

func (s MyStr) Set(v MyStr) {
	s = v
}

// 递归类型, 约束了 Set(T) & Get() T 类型必须和 value 类型相同
type MyComplex[T Op[T]] struct {
	value T
}

func (m MyComplex[T]) get() T {
	return m.value.Get()
}

func (m *MyComplex[T]) set(v T) {
	m.value = v
}

func TestComplex2(t *testing.T) {
	m := MyComplex[MyStr]{value: "kk"}
	t.Log(m.get())

	m.set("gg")
	t.Log(m.get())
}
