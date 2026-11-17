package majorchanges_test

import (
	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
	"fmt"
	"testing"
)

// 根据 Type 决定如何二次解析 Payload
func TestJsonTextValue(t *testing.T) {
	// jsontext.Value
	type Event struct {
		Type string `json:"type"`

		// 暂不解析，原样保留原始 JSON 字节, jsontext.Value 本质是 []byte
		Payload jsontext.Value `json:"payload"`
	}

	type UserCreated struct {
		Username string `json:"username"`
	}

	raw := []byte(`{"type": "user_created", "payload": {"username": "alice"}}`)

	var ev Event
	if err := jsonv2.Unmarshal(raw, &ev); err != nil {
		panic(err)
	}

	// 根据 Type 决定如何二次解析 Payload
	if ev.Type == "user_created" {
		var u UserCreated
		// jsontext.Value 本质是 []byte，可直接传给 jsonv2.Unmarshal
		if err := jsonv2.Unmarshal(ev.Payload, &u); err != nil {
			panic(err)
		}
		t.Logf("Created user: %s\n", u.Username)
	}
}

// jsontext.Token 零分配的自定义编解码接口(MarshalerTo / UnmarshalerFrom)
// Point 在 Go 中是结构体，但在 JSON 中希望序列化为紧凑数组 [10, 20]
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// 实现 jsonv2.MarshalerTo：直接向 Encoder 写入 Token
func (p Point) MarshalJSONTo(enc *jsontext.Encoder) error {
	// 写入 '['
	if err := enc.WriteToken(jsontext.BeginArray); err != nil {
		return err
	}

	// 写入 X, Y
	if err := enc.WriteToken(jsontext.Int(int64(p.X))); err != nil {
		return err
	}
	if err := enc.WriteToken(jsontext.Int(int64(p.Y))); err != nil {
		return err
	}

	// 写入 ']'
	return enc.WriteToken(jsontext.EndArray)
}

// 实现 jsonv2.UnmarshalerFrom：直接从 Decoder 读取 Token
func (p *Point) UnmarshalJSONFrom(dec *jsontext.Decoder) error {
	// 读取 '['
	tok, err := dec.ReadToken()
	if err != nil {
		return err
	}
	if tok.Kind() != '[' {
		return fmt.Errorf("expected array for Point, got %v", tok.Kind())
	}

	// 读取 X 和 Y
	// xTok, yTok 是 jsontext.Token 类型
	xTok, err := dec.ReadToken()
	if err != nil {
		return err
	}
	if xTok.Kind() != '0' {
		return fmt.Errorf("expected number for X, got %v", xTok.Kind())
	}
	// VVI: jsontext.Token 为了做到极致零内存分配，它内部并不拷贝字符串或数字的原始字节，
	// 而是直接持有指向 jsontext.Decoder 内部读缓冲区（Buffer）的指针引用. 对同一个 Decoder
	// 发起了下一次读取调用. 如再次调用 dec.ReadToken()、dec.ReadValue() 或 dec.SkipValue()
	// 前一个 Token 就会立即失效, 所以在下次使用 decoder 读取数据之前, 先将数据处理好, 或者缓存下来.
	pX, err := xTok.Int()
	if err != nil {
		return err
	}

	// 读取 Y
	yTok, err := dec.ReadToken()
	if err != nil {
		return err
	}
	if yTok.Kind() != '0' {
		return fmt.Errorf("expected number for Y, got %v", yTok.Kind())
	}
	pY, err := yTok.Int()
	if err != nil {
		return err
	}

	// 读取 ']'
	endTok, err := dec.ReadToken()
	if err != nil {
		return err
	}
	if endTok.Kind() != ']' {
		return fmt.Errorf("expected end of array, got %v", endTok.Kind())
	}

	// 给 p 赋值
	p.X = int(pX)
	p.Y = int(pY)

	return nil
}

func TestJsonTextToken(t *testing.T) {
	b, err := jsonv2.Marshal(Point{X: 10, Y: 20})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(b))

	var p Point
	err = jsonv2.Unmarshal(b, &p)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v", p)
}
