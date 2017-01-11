package main

import "github.com/amarburg/go-lazyfs/benchmarking"

import "flag"

var repsFlag = flag.Int("reps", 10, "Number of reps")

func main() {

  flag.Parse()

  setRepsFunc := func( conf *lazyfs_benchmarking.IterationConfig ) { conf.Reps = *repsFlag }
  setIterationsFunc := func( conf *lazyfs_benchmarking.IterationConfig ) { conf.Iterations = []int{100} }

  lazyfs_benchmarking.Iterate( func( bench *lazyfs_benchmarking.Bench ) {
    source := lazyfs_benchmarking.MakeOOIHttpSource()
    bench.Source = "http"
    bench.Store = ""

    bench.RunBenchmark( source )
    bench.HttpBytes = source.Stats.ContentBytesRead
  }, setRepsFunc, setIterationsFunc )

  lazyfs_benchmarking.Iterate( func( bench *lazyfs_benchmarking.Bench ) {
    source := lazyfs_benchmarking.MakeOOIHttpSource()
    store  := lazyfs_benchmarking.MakeSparseStore( source )
    bench.Source = "http"
    bench.Store = "sparse"

    bench.RunBenchmark( store )
    bench.HttpBytes = source.Stats.ContentBytesRead
  }, setRepsFunc , setIterationsFunc )
}
