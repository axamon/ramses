# stringset

[![GoDoc](https://godoc.org/github.com/axamon/stringset?status.svg)](https://godoc.org/github.com/axamon/stringset)
[![Build Status](https://travis-ci.org/axamon/stringset.svg?branch=master)](https://travis-ci.org/axamon/stringset)
[![Go Report Card](https://goreportcard.com/badge/github.com/axamon/stringset)](https://goreportcard.com/report/github.com/axamon/stringset)
[![codecov](https://codecov.io/gh/axamon/stringset/branch/master/graph/badge.svg)](https://codecov.io/gh/axamon/stringset)
[![Maintainability](https://api.codeclimate.com/v1/badges/5738c2943ca85e95975d/maintainability)](https://codeclimate.com/github/axamon/stringset/maintainability)
[![HitCount](http://hits.dwyl.io/axamon/stringset.svg)](http://hits.dwyl.io/axamon/stringset)

stringset creates sets for strings in golang that are concurrency safe



## Installation
```go get -u github.com/axamon/stringset```

## Usage
```go

package main

import (
    "fmt"
    "github.com/axamon/stringset"
)

func main() {
    testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")
      
	slice := testSet.Strings()
	sort.Strings(slice)
	for _, element := range slice {
		fmt.Println(element)
	}
	// Output:
	// paperino
	// pippo
	// pluto
}
```

# Benchmarks

```
BenchmarkAdd-16          	 3000000	       497 ns/op

BenchmarkDelete-16       	 3000000	       539 ns/op

BenchmarkIntersect-16    	 1000000	      2168 ns/op

BenchmarkUnion-16        	 2000000	      1825 ns/op
```


