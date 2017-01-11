package main

import "io"
import "os"
import "time"
import "math/rand"

import "github.com/amarburg/go-lazyfs-testfiles/http_server"
import "github.com/amarburg/go-lazyfs-testfiles"
import "github.com/amarburg/go-lazyfs/benchmarking"
import "github.com/amarburg/go-lazyfs"

import "fmt"

type BenchmarkResult struct {
  Iter int
  BufSize int
  Source,Store  string
  Duration  time.Duration
  HttpBytes  int
}

func (result BenchmarkResult) Write( out io.Writer ) {

    io.WriteString( out, fmt.Sprintf("http,sparse,%d,%d,%d,%.1f\n",
                result.Iter,
                result.BufSize,
                result.Duration.Nanoseconds()/int64(result.Iter),
                float32(result.HttpBytes) / float32(result.Iter) ) )
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
    return BenchmarkResult{ Iter: N,
                            Duration: time.Now().Sub( startTime ) }
}




func main() {

  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  bufsizes := []int{32,128,256,1024,4096}
  iterations := []int{1e2,1e4,1e6}

  for _,bufsize := range bufsizes {
    for _,iter := range iterations {
      for rep := 0; rep < 2; rep++ {

        source := lazyfs_benchmarking.MakeLocalHttpSource()
        store  := lazyfs_benchmarking.MakeSparseStore( source )

        result := RunBenchmark( store, iter, bufsize )

        result.BufSize = bufsize
        result.Source = "http"
        result.Store  = "sparse"
        result.HttpBytes = source.Stats.ContentBytesRead

        result.Write( os.Stderr )
      }
    }
  }
}
