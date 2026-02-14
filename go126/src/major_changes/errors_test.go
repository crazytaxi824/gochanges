package majorchanges_test

import (
	"errors"
	"fmt"
	"testing"
)

// AppError 是我们的自定义错误类型
type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// 模拟返回一个包装过的 AppError
func DatabaseQuery() error {
	return fmt.Errorf("database access failed: %w", &AppError{
		Code:    404,
		Message: "User not found",
	})
}

// 再次包装错误
func ServiceLayer() error {
	if err := DatabaseQuery(); err != nil {
		return fmt.Errorf("service layer error: %w", err)
	}
	return nil
}

// switch err.(type) 无法识别通过 fmt.Errorf("...: %w", err) 包装过的错误.
func TestErrorsAsType(t *testing.T) {
	err := ServiceLayer()
	if err != nil {
		// 使用 errors.AsType[*AppError](err)
		// 相比旧的 errors.As(&target, err)，它直接返回转换后的类型和布尔值
		if target, ok := errors.AsType[*AppError](err); ok {
			t.Logf("HTTP 状态码: %d\n", target.Code)
			t.Logf("具体信息: %s\n", target.Message)
		} else {
			t.Error("这是一个未知类型的普通错误")
		}
	}
}
