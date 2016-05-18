package bench

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

type JsonGroup struct {
	Type   string `json:"type"`
	Values string `json:"values"`
}

type JsonGroupArr struct {
	Type   string   `json:"type"`
	Values []string `json:"values"`
}

func Values() (tokens string) {
	var buffer bytes.Buffer

	n := 100
	for i := 0; i < n; i++ {
		buffer.WriteString("test")
		if i < n {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func JsonLine() {
	var group *JsonGroup
	group = &JsonGroup{
		Type:   "test",
		Values: Values(),
	}

	// 1
	b, _ := json.Marshal(group)

	// 2
	json.Unmarshal(b, group)

	// 3
	strings.Split(group.Values, ",")
}

func JsonArray() {
	tokens := Values()
	var group *JsonGroupArr
	group = &JsonGroupArr{
		Type: "test",
		// 1
		Values: strings.Split(tokens, ","),
	}

	// 2
	b, _ := json.Marshal(group)

	// 3
	json.Unmarshal(b, group)
}

// test line
func BenchmarkStringSplit(b *testing.B) {
	// 必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		JsonLine()
	}
}

// 测试并发效率
func BenchmarkStringSplitParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			JsonLine()
		}
	})
}

// test array
func BenchmarkJsonArray(b *testing.B) {
	// 必须循环 b.N 次 。 这个数字 b.N 会在运行中调整，以便最终达到合适的时间消耗。方便计算出合理的数据。 （ 免得数据全部是 0 ）
	for i := 0; i < b.N; i++ {
		JsonArray()
	}
}

// 测试并发效率
func BenchmarkJsonArrayParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			JsonArray()
		}
	})
}
