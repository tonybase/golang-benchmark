package bench

import (
	"crypto/tls"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

// go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out http_test.go
// http 1.1
func BenchmarkHttpParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := http.Get("http://s1.tony.wiki:8001/")
			if err != nil {
				log.Fatal(err)
			}
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			//b.Log(string(body))
		}
	})
}

// http 2.0
func BenchmarkHttp2Parallel(b *testing.B) {
	tr := &http2.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := client.Get("https://s1.tony.wiki:8002/")
			if err != nil {
				log.Fatal(err)
			}
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			//b.Log(string(body))
		}
	})
}
