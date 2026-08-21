package generic_test

import (
	jsonv1 "encoding/json"
	"testing"

	jsonv2 "encoding/json/v2"
)

type Address struct {
	City     string `json:"city"`
	State    string `json:"state"`
	Postcode string `json:"postcode"`
}

type User struct {
	ID       int64             `json:"id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Active   bool              `json:"active"`
	Age      int               `json:"age"`
	Score    float64           `json:"score"`
	Tags     []string          `json:"tags"`
	Metadata map[string]string `json:"metadata"`
	Address  Address           `json:"address"`
}

var testUser = User{
	ID:     123456789,
	Name:   "Ray",
	Email:  "ray@example.com",
	Active: true,
	Age:    43,
	Score:  98.75,
	Tags: []string{
		"golang",
		"swift",
		"python",
		"javascript",
	},
	Metadata: map[string]string{
		"language": "en",
		"timezone": "Australia/Sydney",
		"device":   "macOS",
		"role":     "developer",
	},
	Address: Address{
		City:     "Sydney",
		State:    "NSW",
		Postcode: "2000",
	},
}

var jsonData []byte

func init() {
	var err error
	jsonData, err = jsonv1.Marshal(testUser)
	if err != nil {
		panic(err)
	}
}

// Prevent compiler from optimizing the result away.
var (
	resultBytes []byte
	resultUser  User
)

func BenchmarkJSONV1Marshal(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		data, err := jsonv1.Marshal(testUser)
		if err != nil {
			b.Fatal(err)
		}
		resultBytes = data
	}
}

func BenchmarkJSONV2Marshal(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		data, err := jsonv2.Marshal(testUser)
		if err != nil {
			b.Fatal(err)
		}
		resultBytes = data
	}
}

func BenchmarkJSONV1Unmarshal(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		var user User
		if err := jsonv1.Unmarshal(jsonData, &user); err != nil {
			b.Fatal(err)
		}
		resultUser = user
	}
}

func BenchmarkJSONV2Unmarshal(b *testing.B) {
	b.ReportAllocs()

	for b.Loop() {
		var user User
		if err := jsonv2.Unmarshal(jsonData, &user); err != nil {
			b.Fatal(err)
		}
		resultUser = user
	}
}
