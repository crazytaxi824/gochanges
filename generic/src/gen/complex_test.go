package gen_test

import (
	"fmt"
	"testing"
)

// 定义一个约束接口，可以执行某种操作
type Operable[T any] interface {
	~int | ~struct{ v int } // 约束可以实现接口的类型
	Add(T) T                // 传入什么类型就返回什么类型
}

// 定义一个结构体, 约束类型是必须有 Add(T)T 方法, 这里最难理解的是 [T Operable[T]]
// MyNum[T Operable[T]] 是一个递归的类型定义:
//   - 左侧 T 是结构体 MyNum 的类型参数
//   - 右侧的约束 Operable[T] 是表示 MyNum 的类型参数 T 必须实现 Add(T) T 方法,
//     即 T 必须能够接受自己类型的值作为参数, 并返回自己类型的值.
//
// MyNum[T Operable[T]] 效果
//   - 约束了 Add(T)T 中入参和出参的类型必须和 value 的类型相同.
//   - 避免了不同类型之间的混合操作（比如不允许 Integer 和 Float 直接相加）
//
// 所以顺序是:
//  1. 实现接口 Operable
//  2. 将该类型作为 MyNum.v 的类型
type MyNum[T Operable[T]] struct {
	value T
}

// MyNumber 方法使用 Operable 类型来调用 Add
func (n MyNum[T]) AddToValue(val T) T {
	// 这里的 T 是 Operable 类型, 即 MyNumber.value 是 Operable 类型
	return n.value.Add(val)
}

// 实现接口的 int 类型, MyNum.v 的类型为 MyInt
type MyInt int

func (a MyInt) Add(b MyInt) MyInt {
	return a + b
}

// 实现接口的 struct 类型, MyNum.v 的类型为 MyStruct
type MyStruct struct {
	v int
}

func (a MyStruct) Add(b MyStruct) MyStruct {
	return MyStruct{
		v: a.v + b.v,
	}
}

func TestComplex(*testing.T) {
	// MyInt 类型的实现
	num := MyNum[MyInt]{value: 10}
	result := num.AddToValue(5)
	fmt.Println(result) // 15

	// struct 类型实现
	st := MyNum[MyStruct]{value: MyStruct{v: 1}}
	r := st.AddToValue(MyStruct{v: 2})
	fmt.Println(r.v) // 3
}
