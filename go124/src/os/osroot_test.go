// 新增 os.Root 用于防止错误的读取文件, 将文件读取限制在指定的目录下.
package os_test

import (
	"os"
	"testing"
)

func TestOsRoot(t *testing.T) {
	root, err := os.OpenRoot("/path/to/dir")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(root.Name())

	// open file under root dir
	f, err := root.Open("test.txt")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(f.Fd())

	// get file info under root dir
	finfo, err := root.Stat("test.txt")
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v", finfo)

	// root.Mkdir()
	// root.OpenFile()
	// ...
}
