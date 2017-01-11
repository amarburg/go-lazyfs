package main

import "testing"
import "github.com/amarburg/go-lazyfs-testfiles/http_server"
import "github.com/amarburg/go-lazyfs/benchmarking"

import "fmt"

func main() {

  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  bufsizes := []int{32,128,256,1024}
  iterations := []int{10}

  for _,bufsize := range bufsizes {
    for _,iter := range iterations {

      fmt.Printf("Testing for buffer size %d, %d iterations\n", bufsize, iter )

      results := make([]testing.BenchmarkResult, 0, iter )

      for i := 0; i < iter; i++  {
        result := testing.Benchmark( func ( b *testing.B ) {
          source := lazyfs_benchmarking.MakeLocalHttpSource()
          store  := lazyfs_benchmarking.MakeSparseStore( source )

           bench := lazyfs_benchmarking.LazyFSBenchmark{
             BufSize: bufsize,
           }
           bench.Run( store, b )
        } )

fmt.Println(result)

        results = append( results, result )
      }

    }
  }
}
