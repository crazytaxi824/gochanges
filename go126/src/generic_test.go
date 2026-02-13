package src_test

import (
	"fmt"
	"testing"
)

// 递归类型实例化
// 可以实现真正的"流式接口"
type INode[T INode[T]] interface {
	Children() []T // T 还是 INode 类型
	Name() string
}

func rangeNodes[P INode[P]](n P) {
	fmt.Println(n.Name())
	c := n.Children()
	if len(c) < 1 {
		return // stop
	} else {
		for _, v := range c {
			rangeNodes(v)
		}
	}
}

type UserNode struct {
	children []*UserNode // 实际存储数据的字段
	UserName string
}

// 实现接口方法
func (u *UserNode) Children() []*UserNode {
	return u.children
}

func (u *UserNode) Name() string {
	return u.UserName
}

func TestGeneric(t *testing.T) {
	leaf := UserNode{
		UserName: "leaf",
	}

	n1 := UserNode{
		children: []*UserNode{&leaf},
		UserName: "n1",
	}

	n2 := UserNode{
		children: []*UserNode{&n1},
		UserName: "n2",
	}

	rangeNodes(&n2)
}
