# Should I Use a Pointer instead of a Copy of my Struct?

```
go get golang.org/x/perf/benchstat
go install
```

```
go test ./... -bench=BenchmarkMemoryHeap -benchmem -run=^$ -count=10 > heap.txt && benchstat head.txt
go test ./... -bench=BenchmarkMemoryStack -benchmem -run=^$ -count=10 > stack.txt && benchstat stack.txt
```

https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963