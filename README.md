# gobench
Benchmark for atomic instruction vs acquiring a read lock vs no synchronization for golang

# Single Threaded Write Benchmark
```
goos: darwin
goarch: amd64
pkg: github.com/mangalaman93/gobench

BenchmarkIncr-8
2000000000	         1.96 ns/op

BenchmarkIncrAtomic-8
300000000	         5.96 ns/op

BenchmarkIncrLock-8
100000000	        32.0 ns/op
```
