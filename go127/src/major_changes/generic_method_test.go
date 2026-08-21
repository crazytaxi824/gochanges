package generic_test

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"testing"
)

// std pkg example
func TestExample(t *testing.T) {
	t.Log(rand.N(int(100)))
	t.Log(rand.N(int32(100)))
	t.Log(rand.N(int64(100)))
}

// 实现
type Obj struct{}

func (m *Obj) Foo[T comparable](t T) {
	fmt.Println(reflect.TypeOf(t).String())
}

func TestGenericMethod(t *testing.T) {
	var m Obj
	m.Foo(int16(1))
}
