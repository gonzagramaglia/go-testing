### Benchmarking

Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

```
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```
You'll see the code is very similar to a test.
The testing.B gives you access to the cryptically named b.N.
When the benchmark code is executed, it runs b.N times and measures how long it takes.
The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.
To run the benchmarks do `go test -bench=.`

NOTE by default Benchmarks are run sequentially.