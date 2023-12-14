# go-pointers-benchmark
Benchmarking by value and by pointer in Go

## How to run

```
$ go test -gcflags=-N -bench=. -count 1
```

Sample output

```
goos: linux
goarch: amd64
pkg: github.com/skartikey/go-pointers-benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-11850H @ 2.50GHz
BenchmarkDirectValues-16             669           1920348 ns/op
BenchmarkUsingPointers-16            560           1814547 ns/op
BenchmarkPassByValue-16         474365810                2.479 ns/op
BenchmarkPassByPointer-16       396106653                3.074 ns/op
PASS
ok      github.com/skartikey/go-pointers-benchmark      5.681s
```
