package assert

import (
	"reflect"
	"testing"
)

var v any

func TestSwitchType(t *testing.T) {
	v = 100
	t.Log(reflect.TypeOf(v))

	v = "abc"
	t.Log(reflect.TypeOf(v))

	switch v.(type) {
	case int:
		t.Log("int")
	case string:
		t.Log("str")
	default:
		t.Error("error")
	}
}

// 新加入 reflect.TypeAssert() 方法
func TestAssert(t *testing.T) {
	v = 100
	t.Log(reflect.TypeOf(v))

	// false
	a, ok := v.(int64)
	t.Log(a)
	t.Log(ok)

	// OK
	b, ok := reflect.TypeAssert[int](reflect.ValueOf(v))
	t.Log(b)
	t.Log(ok)

	// false
	c, ok := reflect.TypeAssert[int64](reflect.ValueOf(v))
	t.Log(c)
	t.Log(ok)
}
