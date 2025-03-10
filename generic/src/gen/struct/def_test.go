package struct_test

import "testing"

type DefStructA interface {
	~struct {
		key   string
		value any
	}
}

type DefStructB interface {
	struct {
		key   string
		value any
	}
}

func structConstrainATest[T DefStructA](s T) {}

func structConstrainBTest[T DefStructB](s T) {}

func TestStructConstrain(t *testing.T) {
	structConstrainATest(struct {
		key   string
		value any
	}{})
	structConstrainBTest(struct {
		key   string
		value any
	}{})

	type foo struct {
		key   string
		value any
	}
	structConstrainATest(*new(foo))
	// structConstrainBTest(*new(foo)) // ERROR: foo does not satisfy DefStructB

	type bar struct {
		key   string
		value int
	}
	// structConstrainATest(*new(bar)) // ERROR: foo2 does not satisfy DefStructA
	// structConstrainBTest(*new(bar)) // ERROR: foo2 does not satisfy DefStructB
}
