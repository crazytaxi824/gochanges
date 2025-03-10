package interface_test

// type constrain
type DefInts interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// methods
type DefMethods interface {
	Run()
	Walk()
}

type DefMix interface {
	~int | ~int64
	Print()
}

type DefMix2[T interface{ ~int | ~int64 }] interface {
	Print()
}

// constrain only interface 只能作为 Type Param.
// func constrainAsArgs(s MyInts) {} // ERROR: cannot use type MyInts outside a type constraint
func constrainTypeParam[T DefInts](T) {}

// 只要有 constrain 就无法当作 Args Type 使用
// func mixAsArgs(MyMix) {} // ERROR: cannot use type MyMix outside a type constraint
func mixTypeParam[T DefMix](T) {}

// method only interface 可以作为 Type Param 也可以作为 Args Type 使用
func methodsTypeParam[T DefMethods](T) {}

func methodsAsArgs(DefMethods) {}

// 带 constrain 的 method only interface 可以作为 Type Param 也可以作为 Args Type 使用
func methodsTypeParam2[T DefMix2[int]](T) {}

func methodsAsArgs2(DefMix2[int]) {}
