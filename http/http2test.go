package main

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const MAX = 10000

func main() {
	tr := &http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	now := time.Now()
	for i := 0; i < MAX; i++ {
		resp, err := client.Get("https://s1.tony.wiki:8002/")
		if err != nil {
			log.Fatal(err)
		}
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Max: ", MAX, " Time:", time.Now().Sub(now))
}
