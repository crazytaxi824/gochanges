package unique_test

import (
	"testing"
	"unique"
)

func TestUniq(t *testing.T) {
	str1 := "hello"
	str2 := "hello"
	t.Logf("%p %p", &str1, &str2) // 0x14000106300 0x14000106310

	h1 := unique.Make(str1)
	h2 := unique.Make(str2)

	// comparable
	t.Log(h1, h2)   // {0x14000106350} {0x14000106350}
	t.Log(h1 == h2) // true
}
