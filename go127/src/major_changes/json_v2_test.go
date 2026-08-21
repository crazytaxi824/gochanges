package generic_test

import (
	v1 "encoding/json"
	"testing"

	v2 "encoding/json/v2"
)

// v1: nil slice & map -> null
// v2: nil slice & map -> [] / {}
func TestJson(t *testing.T) {
	type jsonData struct {
		Slice []int          `json:"slice"`
		Map   map[int]string `json:"map"`
	}
	data := jsonData{}

	b1, _ := v1.Marshal(data)
	t.Log(string(b1)) // {"slice":null,"map":null}

	b2, _ := v2.Marshal(data)
	t.Log(string(b2)) // {"slice":[],"map":{}}
}

// Omitempty diff
func TestJson2(t *testing.T) {
	type Profile struct {
		Age  int    `json:",omitempty"`
		Note string `json:",omitempty"`
	}

	type User struct {
		Name string   `json:",omitempty"`
		Info Profile  `json:",omitempty"` // 嵌套结构体
		Tags []string `json:",omitempty"` // 空切片
	}
	u := User{
		Name: "",
		Info: Profile{Age: 0, Note: ""}, // 内部全为零值
		Tags: []string{},                // 已初始化但长度为 0
	}

	b1, _ := v1.Marshal(u)
	t.Log(string(b1)) // {"Info":{}}

	b2, _ := v2.Marshal(u)
	t.Log(string(b2)) // {"Info":{"Age":0}}
}
