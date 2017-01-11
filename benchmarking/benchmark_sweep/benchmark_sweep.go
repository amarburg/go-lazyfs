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

    io.WriteString( out, fmt.Sprintf("http,sparse,%d,%d,%d,%.1f\n",
                result.Iter,
                result.BufSize,
                result.Duration.Nanoseconds()/int64(result.Iter),
                float32(result.HttpBytes) / float32(result.Iter) ) )
}


func (bench *Bench) RunBenchmark( source lazyfs.FileSource )  {

    startTime := time.Now()
    for i := 0; i < bench.Iter; i++ {
      offset := rand.Intn( lazyfs_testfiles.TenMegFileLength - bench.BufSize )

        buf := make([]byte,bench.BufSize)

        // Test ReadAt
        n,_ := source.ReadAt(buf, int64(offset))
        if n != bench.BufSize { panic("bad read") }
    }

    bench.Duration = time.Now().Sub( startTime )
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

        bench := Bench{
          BufSize: bufsize,
          Source: "http",
          Store:  "sparse",
          Iter: iter,
        }

        bench.RunBenchmark( store )
        bench.HttpBytes = source.Stats.ContentBytesRead

        bench.Write( os.Stderr )
      }
    }
  }
}
