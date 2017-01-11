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

type IterationConfig struct {
  BufSizes    []int
  Iterations  []int
  Reps        int
}


func Iterate( benchFunc func( bench *Bench ),
              confFuncs ...func( conf *IterationConfig ) ){

  config := IterationConfig{
    BufSizes: []int{32,128,256,1024,4096},
    Iterations: []int{1e2,1e4},
    Reps: 10,
  }

  for _,f := range confFuncs {
    f( &config )
  }

  for _,bufsize := range config.BufSizes {
    for _,iter := range config.Iterations {
      for rep := 0; rep < config.Reps; rep++ {

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
