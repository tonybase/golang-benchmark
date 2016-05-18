package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

func main() {
	m := make(map[int64]int32)
	for i := 0; i < 100000; i++ {
		m[int64(i)] = int32(i % 10)
	}
	// encode
	now := time.Now()
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("bytes: %d\n", len(buff.Bytes()))
	fmt.Printf("encoder used: %s\n", time.Now().Sub(now))

	// clc
	now = time.Now()
	r := make(map[int32]int32)
	for _, v := range m {
		r[v]++
	}
	fmt.Println(time.Now().Sub(now))
	count := int32(0)
	for i := 0; i < 10; i++ {
		count += r[int32(i)]
	}
	fmt.Printf("counter: %s\n", time.Now().Sub(now))
	fmt.Println("", count)
}
