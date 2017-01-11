package lazyfs_benchmarking

import "testing"
import "github.com/amarburg/go-lazyfs"
import "github.com/amarburg/go-lazyfs-testfiles"

import "math/rand"
import "net/url"
import "fmt"

type LazyFSBenchmark struct {
  BufSize     int
}

func (bench *LazyFSBenchmark) Run( source lazyfs.FileSource, b *testing.B ) {

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
      offset := rand.Intn( lazyfs_testfiles.TenMegFileLength - BufSize )

        buf := make([]byte,BufSize)

        // Test ReadAt
        n,_ := source.ReadAt(buf, int64(offset))
        if n != BufSize { panic("bad read") }

        b.SetBytes( int64(n) )

    }
    b.StopTimer()

}


func MakeLocalHttpSource() (*lazyfs.HttpSource ) {
  var SourceUrl,_ = url.Parse( "http://localhost:4567/" + lazyfs_testfiles.TenMegBinaryFile )
  source,err := lazyfs.OpenHttpSource( *SourceUrl )
  if err != nil {
    panic("Couldn't open HttpFSSource")
  }

  return source
}

func MakeSparseStore( in lazyfs.FileSource ) (*lazyfs.SparseFileStore) {
  HttpSourceSparseStore := "test_cache/local_benchmark/"
  store,err :=  lazyfs.OpenSparseFileStore( in, HttpSourceSparseStore )
  if store == nil || err != nil {
    if err == nil {
      panic("Couldn't open SparseFileStore")
    } else {
      panic(fmt.Sprintf("Couldn't open SparseFileStore: %s", err.Error() ))
    }
  }
  return store
}
