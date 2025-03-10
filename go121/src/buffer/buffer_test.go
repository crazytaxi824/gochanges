package buffer_test

import (
	"bytes"
	"testing"
)

func TestBytesBuffer(t *testing.T) {
	buf := bytes.NewBuffer([]byte("abc"))
	t.Log(buf.Len(), buf.Cap(), buf.Available())

	// buf.Available() == 底层 cap([]byte)-len([]byte), 而不是 buf.Cap()-buf.Len()
	t.Log(string(buf.Next(1)))
	t.Log(buf.Len(), buf.Cap(), buf.Available())

	// Grow() 增长 cap, 而不是 len
	buf.Grow(10)
	t.Log(buf.Len(), buf.Cap(), buf.Available())

	buf.WriteString("def")
	t.Log(buf.Len(), buf.Cap(), buf.Available())
}

// AvailableBuffer() 返回一个 []byte, 用于填满整个buffer.
// cap([]byte) == buf.Available() == 底层 cap([]byte)-len([]byte)
func TestBuffer(t *testing.T) {
	var buf bytes.Buffer
	t.Log(buf.Len(), buf.Cap(), buf.Available())

	buf.WriteString("abc")
	t.Log(buf.Len(), buf.Cap(), buf.Available())

	// AvailableBuffer() 用于填满整个buffer, 而不自动 Grow()
	b := buf.AvailableBuffer()
	t.Log(len(b), cap(b)) // len(b) == 0, cap(b) == buf.Available()

	for i := 0; i < cap(b); i++ {
		b = append(b, 'z')
	}

	buf.Write(b)
	t.Log(buf.String())
	t.Log(buf.Len(), buf.Cap(), buf.Available())
}
