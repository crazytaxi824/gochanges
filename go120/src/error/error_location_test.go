package error

import (
	"errors"
	"fmt"
	"runtime"
	"testing"
)

type MyWrapError struct {
	err  error
	file string
	line int
}

func (me *MyWrapError) Error() string {
	return fmt.Sprintf("%s:%d: %s", me.file, me.line, me.err.Error())
}

// 实现 Is 之后可以使用 errors.Is() 来判断类型
func (*MyWrapError) Is(target error) bool {
	_, ok := target.(*MyWrapError)
	return ok
}

// helper: wraps error with <file:line> info
func withLocation(err error) error {
	if err == nil {
		return nil
	}

	// 判断 error 是否已经 wrap
	if errors.Is(err, &MyWrapError{}) {
		return err
	}

	_, file, line, ok := runtime.Caller(1) // 1 = caller of this function
	if !ok {
		return &MyWrapError{
			err: err,
		}
	}
	return &MyWrapError{
		err:  err,
		file: file,
		line: line,
	}
}

// root cause
var ErrDBConn = errors.New("db connection failed")

// lowest layer
func dbQuery() error {
	// simulate DB failure
	return withLocation(ErrDBConn)
}

// repo layer
func repoGetUser() error {
	if err := dbQuery(); err != nil {
		return withLocation(err)
	}
	return nil
}

// service layer
func serviceGetUser() error {
	if err := repoGetUser(); err != nil {
		return withLocation(err)
	}
	return nil
}

// top layer (handler)
func handler() error {
	if err := serviceGetUser(); err != nil {
		return withLocation(err)
	}
	return nil
}

func TestErrorLocation(t *testing.T) {
	err := handler()
	t.Log(err)
}
