package lazyfs_benchmarking

import "io"
import "os"
import "time"
import "math/rand"
import "fmt"

import "github.com/amarburg/go-lazyfs"

var TenMegs int64 = 10485760

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

  sz,_ := source.FileSize()
  if sz > TenMegs { sz = TenMegs }

  startTime := time.Now()
  for i := 0; i < bench.Iter; i++ {
    offset := rand.Intn( int(sz - int64(bench.BufSize)) )

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
