package main

import (
	"sync"
	"testing"
)

type Bucket struct {
	lock sync.RWMutex
	V    []int
}

func singleLock() {
	b := &Bucket{}
	for i := 0; i < 100000; i++ {
		b.lock.Lock()
		b.V = append(b.V, i)
		b.lock.Unlock()
	}
}

func multiLock() {
	var (
		bs []*Bucket
		b  *Bucket
	)
	for i := 0; i < 64; i++ {
		bs = append(bs, &Bucket{})
	}
	for i := 0; i < 100000; i++ {
		b = bs[i%64]
		b.lock.Lock()
		b.V = append(b.V, i)
		b.lock.Unlock()
	}
}

func BenchmarkSLock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			singleLock()
		}
	})
}

func BenchmarkMLock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			multiLock()
		}
	})
}
