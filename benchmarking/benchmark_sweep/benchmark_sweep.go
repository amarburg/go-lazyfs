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

type Bench struct {
  Iter int
  BufSize int
  Source,Store  string
  Duration  time.Duration
  HttpBytes  int
}

func (result Bench) Write( out io.Writer ) {

  io.WriteString( out, fmt.Sprintf("%s,%s,%d,%d,%d,%.1f\n",
    result.Source, result.Store,
    result.Iter,
    result.BufSize,
    result.Duration.Nanoseconds()/int64(result.Iter),
    float32(result.HttpBytes) / float32(result.Iter) ) )
}


func (bench *Bench) RunBenchmark( source lazyfs.FileSource )  {

  sz := source.FileSize()

  startTime := time.Now()
  for i := 0; i < bench.Iter; i++ {
    offset := rand.Intn( sz - bench.BufSize )

    buf := make([]byte,bench.BufSize)

    // Test ReadAt
    n,_ := source.ReadAt(buf, int64(offset))
    if n != bench.BufSize { panic("bad read") }
  }

  bench.Duration = time.Now().Sub( startTime )
}



func Iterate( benchFunc func( bench *Bench ) ) {

  bufsizes := []int{32,128,256,1024,4096}
  iterations := []int{1e2,1e4}

  for _,bufsize := range bufsizes {
    for _,iter := range iterations {
      for rep := 0; rep < 2; rep++ {

        bench := &Bench{
          BufSize: bufsize,
          Iter: iter,
        }

        benchFunc( bench )

        bench.Write( os.Stderr )
      }
    }
  }

}


func main() {

  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  Iterate( func( bench *Bench ) {
    source := lazyfs_benchmarking.MakeLocalHttpSource()
    bench.Source = "http"
    bench.Store = ""

    bench.RunBenchmark( source )
    bench.HttpBytes = source.Stats.ContentBytesRead
  })

  Iterate( func( bench *Bench ) {
    source := lazyfs_benchmarking.MakeLocalHttpSource()
    store  := lazyfs_benchmarking.MakeSparseStore( source )
    bench.Source = "http"
    bench.Store = "sparse"

    bench.RunBenchmark( store )
    bench.HttpBytes = source.Stats.ContentBytesRead
  })
}
