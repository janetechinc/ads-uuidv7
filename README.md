# uuid7

Documentation:
https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-03

## Example

```go
package main

import (
    "log"

    uuidv7 "github.com/janetechinc/ads-uuidv7"
)

func main() {
    u := uuidv7.New()

    log.Println(u.Next().String())
}
```

## Benchmark

```
% go test -bench=. -test.benchmem
goos: darwin
goarch: arm64
pkg: github.com/janetechinc/ads-uuidv7
BenchmarkNext-8         21308299                53.93 ns/op            0 B/op          0 allocs/op
BenchmarkString-8       44208190                26.25 ns/op           48 B/op          1 allocs/op
BenchmarkParse-8        25415180                46.92 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/janetechinc/ads-uuidv7       4.701s
```