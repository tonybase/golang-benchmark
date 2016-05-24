package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const MAX = 3000

var (
	now         = time.Now()
	count int32 = 0
	q           = make(chan int, 1)
)

func main() {
	for i := 0; i <= MAX; i++ {
		go req()
	}
	<-q
}

func req() {
	resp, err := http.Get("http://127.0.0.1:8001/")
	if err != nil {
		panic(err)
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	atomic.AddInt32(&count, 1)
	if count == MAX {
		log.Println("MAX: ", MAX, " Time:", time.Now().Sub(now))
		q <- 1
	}
}
