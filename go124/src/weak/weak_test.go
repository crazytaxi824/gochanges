// 新增 weak 包, 用于优化内存.
// 弱指针是为了不阻止 GC 回收 object. 用于释放不常用内存. 类似冷热数据储存.
package weak_test

import (
	"runtime"
	"testing"
	"time"
	"unsafe"
	"weak"
)

func TestWeak(t *testing.T) {
	f := []byte("file content ...") // 模拟大型文件的内容

	fwp := weak.Make(&f)    // 将 slice 的地址存入弱指针
	ewp := weak.Make(&f[0]) // 将底层 array 内存地址存入弱指针

	// []byte slice 的内存地址
	t.Logf("%p %p %p", &f, unsafe.Pointer(&f), fwp.Value())
	// 底层 array 的内存地址
	t.Logf("%p %p %p", &f[0], unsafe.Pointer(&f[0]), ewp.Value())

	// *weak.Value() 是获取 f 的值.
	t.Log(string(*fwp.Value())) // file content ...

	// weak Pointer is comparable
	t.Log(fwp == weak.Make(&f)) // true

	// 对 slice 的回收检查
	runtime.AddCleanup(&f, func(ptr uintptr) {
		t.Logf("GC Slice %x", ptr)
	}, uintptr(unsafe.Pointer(&f)))

	// 对底层 array 的回收检查
	runtime.AddCleanup(&f[0], func(ptr uintptr) {
		t.Logf("GC Array %x", ptr)
	}, uintptr(unsafe.Pointer(&f[0])))

	t.Logf("GC sleep...")
	runtime.GC()
	time.Sleep(time.Second)

	t.Log(fwp.Value()) // nil, slice 已被回收
	t.Log(ewp.Value()) // nil, 底层 array 已被回收
}
