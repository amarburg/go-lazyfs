package main

import "time"
import "math/rand"

import "github.com/amarburg/go-lazyfs-testfiles/http_server"
import "github.com/amarburg/go-lazyfs-testfiles"
import "github.com/amarburg/go-lazyfs/benchmarking"
import "github.com/amarburg/go-lazyfs"

import "fmt"

type BenchmarkResult struct {
  Duration  time.Duration
}


func RunBenchmark( source lazyfs.FileSource, N int, bufSize int ) BenchmarkResult {

    startTime := time.Now()
    for i := 0; i < N; i++ {
      offset := rand.Intn( lazyfs_testfiles.TenMegFileLength - bufSize )

        buf := make([]byte,bufSize)

        // Test ReadAt
        n,_ := source.ReadAt(buf, int64(offset))
        if n != bufSize { panic("bad read") }

        //b.SetBytes( int64(n) )

    }
    return BenchmarkResult{ Duration: time.Now().Sub( startTime ) }
}




func main() {

  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  bufsizes := []int{32,128,256,1024,4096}
  iterations := []int{10}

  for _,bufsize := range bufsizes {
    for _,iter := range iterations {

      source := lazyfs_benchmarking.MakeLocalHttpSource()
      store  := lazyfs_benchmarking.MakeSparseStore( source )

      result := RunBenchmark( store, iter, bufsize )

      httpBytes := float32(source.Stats.ContentBytesRead) / float32(iter)

      fmt.Printf("http,sparse,%d,%d,%d,%.1f\n",
                  iter, bufsize,
                  result.Duration.Nanoseconds()/int64(iter),
                  httpBytes )
    }
  }
}
