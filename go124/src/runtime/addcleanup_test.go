// runtime.AddCleanup() replace runtime.SetFinalizer()
package runtime_test

import (
	"runtime"
	"testing"
	"time"
	"unsafe"
)

func TestAddCleanup(t *testing.T) {
	i := make([]int, 4)
	t.Logf("1st slice addr: %p, first elem addr: %p\n", &i, i)

	for k := range i {
		t.Logf("%d elem addr: %p\n", k, &i[k])

		// 在 func 内不能使用 unsafe.Pointer() 会触发 fatal error: found pointer to free object
		// 这是因为 object 已经被释放, 使用 unsafe.Pointer() 时会导致指针指向已经无效的内存地址。
		// runtime.AddCleanup(&i[k], func(ptr uintptr) {
		// 	t.Logf("AddCleanup GC %d addr: %p", k, unsafe.Pointer(ptr))
		// }, uintptr(unsafe.Pointer(&i[k])))

		runtime.AddCleanup(&i[k], func(ptr uintptr) {
			t.Logf("1 AddCleanup GC elem %d addr: %x", k, ptr)
		}, uintptr(unsafe.Pointer(&i[k])))

		// runtime.AddCleanup() 的 func() 可以不传递任何数据.
		// runtime.AddCleanup() 可以多次添加到同一个对象上, runtime.SetFinalizer() 对同一个对象只能添加一次.
		// runtime.SetFinalizer() 允许 object 复活, 导致很多问题,同时需要至少2次 GC 才能释放.
		runtime.AddCleanup(&i[k], func(*struct{}) {
			t.Logf("2 AddCleanup GC elem")
		}, nil)
	}

	runtime.AddCleanup(&i, func(ptr uintptr) {
		t.Logf("1 AddCleanup GC slice addr: %x", ptr)
	}, uintptr(unsafe.Pointer(&i)))

	runtime.AddCleanup(&i, func(*struct{}) {
		t.Logf("2 AddCleanup GC slice")
	}, nil)

	runtime.GC()
	t.Log("GC & Sleep")
	time.Sleep(time.Second)

	t.Log("end")
}
