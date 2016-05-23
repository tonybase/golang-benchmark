package bench

import (
	"sync"
	"testing"
)

func TMap() {
	count := 100000
	r := sync.RWMutex{}
	q := make(chan int, 1)
	m := make(map[int]int, count)
	for i := 0; i <= count; i++ {
		n := i
		go func() {
			r.Lock()
			m[n] = n * 2
			r.Unlock()
			if n == count {
				q <- 1
			}
		}()
	}
	<-q
}

func TChan() {
	c := make(chan int32, 10000)
	for i := int32(0); i < 100000; i++ {
		go func() {
			c <- i
		}()
	}
	<-c
}

func BenchmarkMapParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TMap()
		}
	})
}

func BenchmarkChanParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TChan()
		}
	})
}
