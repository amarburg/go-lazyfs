package main

import "github.com/amarburg/go-lazyfs-testfiles/http_server"
import "github.com/amarburg/go-lazyfs/benchmarking"

import "flag"

var repsFlag = flag.Int("reps", 10, "Number of reps")

func main() {
  flag.Parse()

  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  setRepsFunc := func( conf *lazyfs_benchmarking.IterationConfig ) { conf.Reps = *repsFlag }

  lazyfs_benchmarking.Iterate( func( bench *lazyfs_benchmarking.Bench ) {
    source := lazyfs_benchmarking.MakeLocalHttpSource()
    bench.Source = "http"
    bench.Store = ""

    bench.RunBenchmark( source )
    bench.HttpBytes = source.Stats.ContentBytesRead
  }, setRepsFunc )

  lazyfs_benchmarking.Iterate( func( bench *lazyfs_benchmarking.Bench ) {
    source := lazyfs_benchmarking.MakeLocalHttpSource()
    store  := lazyfs_benchmarking.MakeSparseStore( source )
    bench.Source = "http"
    bench.Store = "sparse"

    bench.RunBenchmark( store )
    bench.HttpBytes = source.Stats.ContentBytesRead
  }, setRepsFunc )
}
