# rstats

[![Build Status](https://drone.io/github.com/lucindo/rstats/status.png)](https://drone.io/github.com/lucindo/rstats/latest)
[![Coverage Status](https://coveralls.io/repos/github/lucindo/rstats/badge.svg?branch=master)](https://coveralls.io/github/lucindo/rstats?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucindo/rstats)](https://goreportcard.com/report/github.com/lucindo/rstats)
[![GoDoc](https://godoc.org/github.com/lucindo/rstats?status.svg)](https://godoc.org/github.com/lucindo/rstats)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)

```go
package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/lucindo/rstats"
)

func hello(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	w.Write([]byte("Hello stats!"))
}

func handleStats(fn http.HandlerFunc, stats *rstats.Stats) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fn(w, r)
		stats.Add(time.Since(start).Seconds())
	}
}

func main() {
	stats := rstats.New()
	ticker := time.NewTicker(time.Millisecond * 1500)

	go func(stats *rstats.Stats) {
		for _ = range ticker.C {
			log.Println("http stats:", stats)
		}
	}(stats)

	http.HandleFunc("/", handleStats(hello, stats))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
