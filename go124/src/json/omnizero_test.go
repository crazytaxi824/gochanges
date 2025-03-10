// this is a json tag 'omnizero' test
package json_test

import (
	"encoding/json"
	"testing"
)

// omitzero
type Person struct {
	Name    string `json:"name,omitzero"`
	Age     int    `json:"age,omitzero"`
	Sex     bool   `json:"sex,omitzero"`
	Contact struct {
		Email string `json:"email,omitzero"`
		Phone string `json:"phone,omitzero"`
	} `json:"contact,omitzero"`
}

// omitempty
type Person2 struct {
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Sex     bool   `json:"sex,omitempty"`
	Contact struct {
		Email string `json:"email,omitempty"`
		Phone string `json:"phone,omitempty"`
	} `json:"contact,omitempty"`
}

// omitzero with IsZero()
type MyInt int

func (m MyInt) IsZero() bool {
	// 当 IsZero() 返回 false 时, omitzero 不会忽略.
	return m < 0
}

type Person3 struct {
	Name    string `json:"name,omitzero"`
	Age     MyInt  `json:"age,omitzero"`
	Sex     bool   `json:"sex,omitzero"`
	Contact struct {
		Email string `json:"email,omitzero"`
		Phone string `json:"phone,omitzero"`
	} `json:"contact,omitzero"`
}

func TestOmitzero(t *testing.T) {
	// omitzero
	b, _ := json.Marshal(Person{Age: 0})
	t.Log(string(b)) // {}

	// omitempty
	b, _ = json.Marshal(Person2{Age: 0})
	t.Log(string(b)) // {"contact":{}}

	// omitzero with IsZero()
	b, _ = json.Marshal(Person3{Age: 0})
	t.Log(string(b)) // {"age":0}, 因为 0<0 is false, 所以不会被忽略.
}
