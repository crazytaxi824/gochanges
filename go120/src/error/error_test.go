package error_test

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

func TestErrorIs(t *testing.T) {
	wp := fmt.Errorf("foo %w", os.ErrNotExist)

	// wrap error 放在前面
	t.Log(errors.Is(wp, os.ErrNotExist)) // true
	t.Log(errors.Is(os.ErrNotExist, wp)) // false, wrap error 放在前面
	t.Log(wp == os.ErrNotExist)          // false
}

// Is() 的目的是比较两个 error 是否相同, 包括 wrap 的 error
func TestErrorIs2(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		// err 返回的是一个 fs.PathError, 属于 wrap error
		t.Log(err == fs.ErrNotExist) // false

		if errors.Is(err, fs.ErrNotExist) {
			t.Log("file does not exist")
		} else {
			t.Log(err)
		}
	}
}

// As() 的目的不是比较, 而是断言.
func TestErrorAs(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		// t.Log(err.Path) // 报错

		var perr *fs.PathError

		// NOTE: 将 error 断言成 PathError 类型, 同时传入 perr (指针)中
		// 便于使用 PathError.Path 属性.
		if errors.As(err, &perr) {
			t.Log("Failed at path:", perr.Path)
		} else {
			t.Log(err)
		}
	}
}

var (
	ErrDB   = errors.New("db error")
	ErrAuth = errors.New("auth error")
)

// Join 是将多个 error 添加到一起
func TestErrorJoin(t *testing.T) {
	// combine multiple errors
	err := errors.Join(ErrDB, ErrAuth)

	// print
	t.Log("Joined error:", err)

	// check membership
	if errors.Is(err, ErrDB) {
		t.Log("Contains ErrDB")
	}
	if errors.Is(err, ErrAuth) {
		t.Log("Contains ErrAuth")
	}
}

type ErrA struct{}

type ErrB struct{}

func (*ErrA) Error() string {
	return "ErrA"
}

func (*ErrB) Error() string {
	return "ErrB"
}

func TestErrorJoin2(t *testing.T) {
	err := errors.Join(&ErrA{}, &ErrB{})

	t.Log("Joined error:", err)

	// check error Type
	var myErrA *ErrA
	if errors.As(err, &myErrA) {
		t.Log("As ErrA:", myErrA.Error())
	}

	var myErrB *ErrB
	if errors.As(err, &myErrB) {
		t.Log("As ErrB:", myErrB.Error())
	}
}
