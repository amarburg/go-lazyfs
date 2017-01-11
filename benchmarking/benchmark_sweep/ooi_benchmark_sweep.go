package main

import "net/url"
import "github.com/amarburg/go-lazyfs/benchmarking"
import "github.com/amarburg/go-lazyfs"



func main() {

  lazyfs_benchmarking.Iterate( func( bench *lazyfs_benchmarking.Bench ) {
    source := lazyfs_benchmarking.MakeOOIHttpSource()

    bench.Source = "http"
    bench.Store = ""

    bench.RunBenchmark( source )
    bench.HttpBytes = source.Stats.ContentBytesRead
  })

  lazyfs_benchmarking.Iterate( func( bench *lazyfs_benchmarking.Bench ) {
    source := lazyfs_benchmarking.MakeOOIHttpSource()
    store  := lazyfs_benchmarking.MakeSparseStore( source )
    bench.Source = "http"
    bench.Store = "sparse"

    bench.RunBenchmark( store )
    bench.HttpBytes = source.Stats.ContentBytesRead
  })
}
