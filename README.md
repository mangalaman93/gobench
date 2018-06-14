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

# Single Threaded Read Benchmark
```
goos: darwin
goarch: amd64
pkg: github.com/mangalaman93/gobench

BenchmarkRead-8
2000000000	         0.41 ns/op

BenchmarkReadAtomic-8
2000000000	         0.42 ns/op

BenchmarkReadLock-8
20000000	        58.9 ns/op
```

# References
* http://wysocki.in/golang-concurrency-data-races/
* https://software.intel.com/en-us/blogs/2013/01/06/benign-data-races-what-could-possibly-go-wrong
