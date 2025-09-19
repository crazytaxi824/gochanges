package unsafepkg

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestStringToBytes(t *testing.T) {
	s := "abc"

	// 找到 string 起始内存地址
	pb := unsafe.StringData(s)
	t.Log(*pb)

	bs := unsafe.Slice(pb, len(s))
	t.Log(bs)
	t.Log(reflect.TypeOf(bs))

	// Panic
	bs[0] = 100
	t.Log(bs)
}

func TestBytesToString(t *testing.T) {
	byt := []byte{97, 98, 99}

	s := unsafe.String(&byt[0], len(byt))
	t.Log(s) // abc

	byt[0] = 100
	t.Log(s) // dbc
}
