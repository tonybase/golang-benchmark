package main

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const MAX = 3000

var (
	now          = time.Now()
	count  int32 = 0
	q            = make(chan int, 1)
	client *http.Client
)

func main() {
	tr := &http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	for i := 0; i < MAX; i++ {
		go reqHttp2()
	}
	<-q
}

func reqHttp2() {
	resp, err := client.Get("https://127.0.0.1:8002/")
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
