package majorchanges_test

import (
	jsonv1 "encoding/json"
	jsonv2 "encoding/json/v2"
	"testing"
)

// v1: nil slice & map -> null
// v2: nil slice & map -> [] / {}
func TestJson(t *testing.T) {
	type jsonData struct {
		Slice []int          `json:"slice"`
		Map   map[int]string `json:"map"`
	}
	data := jsonData{}

	b1, _ := jsonv1.Marshal(data)
	t.Log(string(b1)) // {"slice":null,"map":null}

	b2, _ := jsonv2.Marshal(data)
	t.Log(string(b2)) // {"slice":[],"map":{}}
}

// omitempty 标签语义在 V2 中会忽略编码结果为 null、""、[] 或 {} 的字段
func TestJson2(t *testing.T) {
	type Profile struct {
		Age  int    `json:",omitempty"`
		Note string `json:",omitempty"`
	}

	type User struct {
		Name    string            `json:",omitempty"`
		Info    Profile           `json:",omitempty"` // 嵌套结构体
		Tags    []string          `json:",omitempty"` // 空切片
		Contact map[string]string `json:",omitempty"` // 空 map
	}
	u := User{
		Name:    "",
		Info:    Profile{Age: 0, Note: ""}, // 0 不会被忽略, "" 会被忽略
		Tags:    []string{},                // [] 会被 omitempty 忽略
		Contact: map[string]string{},       // {} 会被 omitempty 忽略
	}

	b1, _ := jsonv1.Marshal(u)
	t.Log("v1:", string(b1)) // {"Info":{}}

	b2, _ := jsonv2.Marshal(u)
	t.Log("v2:", string(b2)) // {"Info":{"Age":0}}
}
