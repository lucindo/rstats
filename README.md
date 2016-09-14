# rstats

[![Build Status](https://drone.io/github.com/lucindo/rstats/status.png)](https://drone.io/github.com/lucindo/rstats/latest)
[![Coverage Status](https://coveralls.io/repos/github/lucindo/rstats/badge.svg?branch=master)](https://coveralls.io/github/lucindo/rstats?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucindo/rstats)](https://goreportcard.com/report/github.com/lucindo/rstats)
[![GoDoc](https://godoc.org/github.com/lucindo/rstats?status.svg)](https://godoc.org/github.com/lucindo/rstats)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.png?v=103)](https://opensource.org/licenses/mit-license.php)

`rstats` is a simple helper library to calculate running statistics. It's memory efficient using only 7 internal variables and update statistics values on the fly every time you push data to it (no arrays or slices involved).

### Credits

This is a Go port of C++ code by [John D. Cook](http://www.johndcook.com/). Please check his great [blog post](http://www.johndcook.com/blog/standard_deviation/).

The implementation uses the Welford method presented in Donald Knuth’s Art of Computer Programming, Vol 2, page 232, 3rd edition. Please refer to John D. Cook's post for a detailed explanation and references.

### How to use

Get a new instance and push `float64` values to it with `Add` method:

```go
    stats := rstats.New()
    // ...
    stats.Add(/*some value*/)
```

At any time you can call the methods to get statiscs:

 * `Count`
 * `Min`
 * `Max`
 * `Mean`
 * `StandardDeviation`
 * `Variance`
 * `Skewness`
 * `Kurtosis`

`Stats` also implements `Stringer` and returns an string like this:

```
count 1893209 min 0.01 max 1.80 mean 0.51 (std dev 0.289 variance 0.08)
```

You can reset the state calling `Reset` method.

### HTTP Statiscs

I use this package mostly to measure requests/connections on my projects. Here is a dummy implementation collecting execution time statistics of a HTTP handler, and periodically logging it.

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
		defer stats.Add(time.Since(start).Seconds())
		fn(w, r)
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

This package also provides an auxiliary struct to serialize the statistics:

```go
    stats := rstats.New()
    // ...
    stats.Add(/*some value*/)
    // ...
	statsStruct := new(rstats.StatsStruct)
	rstats.GetStatsStruct(statsStruct, stats)
    // serialize statsStruct
```

I use it to expose statistics in `expvar` this way:

```go
import (
	"expvar"

	"github.com/lucindo/rstats"
)

func statsexpvar(stats *rstats.Stats) expvar.Func {
	return func() interface{} {
		statsStruct := new(rstats.StatsStruct)
		rstats.GetStatsStruct(statsStruct, stats)
		return *statsStruct
	}
}

func main() {
	stats := rstats.New()
	expvar.Publish("MyStats", expvar.Func(statsexpvar(stats)))

    //...
}
```
