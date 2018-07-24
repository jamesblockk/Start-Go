package json

import (
	"testing"

	"encoding/json"

	jsoniter "github.com/json-iterator/go"
)

const (
	jstr = `{"req":[{"uid":"1","url":"testUrl1"},{"uid":"2","url":"testUrl2"}]}`
)

type MyJsonName struct {
	Req []struct {
		UID string `json:"uid"`
		URL string `json:"url"`
	} `json:"req"`
}

func Benchmark_array_by_json(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := make([]int, 0, 10)
		json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
	}
}

func Benchmark_array_by_jsoniter(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		sample := make([]int, 0, 10)
		jsoniter.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9]`), &sample)
	}
}

func Benchmark_decode_by_json(b *testing.B) {
	var jsonstr MyJsonName
	for n := 0; n < b.N; n++ {
		json.Unmarshal([]byte(jstr), &jsonstr)
	}
}

func Benchmark_decode_by_jsoniter(b *testing.B) {
	var jsonstr MyJsonName
	for n := 0; n < b.N; n++ {
		jsoniter.Unmarshal([]byte(jstr), &jsonstr)
	}
}
