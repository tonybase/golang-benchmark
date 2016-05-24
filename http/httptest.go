package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const MAX = 10000

func main() {
	now := time.Now()
	for i := 0; i < MAX; i++ {
		resp, err := http.Get("http://s1.tony.wiki:8001/")
		if err != nil {
			log.Fatal(err)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("MAX: ", MAX, " Time:", time.Now().Sub(now))
}
