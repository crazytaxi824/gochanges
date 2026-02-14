package majorchanges_test

import "testing"

type Foo struct {
	A *string
	B *int
}

func TestNew(t *testing.T) {
	foo := Foo{
		A: new(string("abc")),
		B: new(int(24)),
	}
	t.Log(foo)
}
