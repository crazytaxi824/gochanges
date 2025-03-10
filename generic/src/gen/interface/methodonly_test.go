package interface_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 定义一般接口
type MethodsOnly interface {
	Eat(s string)
}

type MethodsOnly2[T any] interface {
	Value() T
}

// 实现接口
type mo struct {
	name string
}

func (m *mo) Eat(s string) {
	fmt.Printf("%s eats %s\n", m.name, s)
}

type mo2[T any] struct {
	value T
}

func (m *mo2[T]) Value() T {
	return m.value
}

// 测试接口
func methodsOnlyAsTypeParam[T MethodsOnly](t T) {
	fmt.Println(reflect.TypeOf(t))
	fmt.Printf("%T\n", t)
	t.Eat("food")
}

func methodsOnlyAsArgs(t MethodsOnly) {
	fmt.Println(reflect.TypeOf(t))
	fmt.Printf("%T\n", t)
	t.Eat("food")
}

// 测试接口2
func methodsOnly2AsTypeParam[T MethodsOnly2[E], E any](t T) {
	fmt.Println(reflect.TypeOf(t))
	fmt.Printf("%T\n", t)
	fmt.Println(t.Value())
}

func methodsOnly2AsArgs[T any](t MethodsOnly2[T]) {
	fmt.Println(reflect.TypeOf(t))
	fmt.Printf("%T\n", t)
	fmt.Println(t.Value())
}

func TestMethodOnly(*testing.T) {
	m := mo{
		name: "kk",
	}
	methodsOnlyAsTypeParam(&m)
	methodsOnlyAsArgs(&m)

	m2 := mo2[string]{
		value: "gg",
	}
	methodsOnly2AsTypeParam(&m2)
	methodsOnly2AsArgs(&m2)
}

// switch types
func swithTypeParam[T any](t T) {
	switch any(t).(type) {
	case int:
		fmt.Println("int")

	default:

	}
}

func switchArgsType(t any) {
	switch t.(type) {
	case int:
		fmt.Println("int")

	default:

	}
}

func TestSwitchType(*testing.T) {
	swithTypeParam(1)
	switchArgsType(1)
}
