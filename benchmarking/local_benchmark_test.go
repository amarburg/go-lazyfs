package lazyfs_benchmarking

import "testing"
import "github.com/amarburg/go-lazyfs-testfiles/http_server"

func BenchmarkLocalHttpSource( b *testing.B ) {
  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  source := MakeLocalHttpSource()
  bench :=  LazyFSBenchmark{
    BufSize: 1024,
  }

  bench.Run( source, b )
}

func BenchmarkLocalHttpSourceSparseStore( b *testing.B ) {
  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  source := MakeLocalHttpSource()
  store  := MakeSparseStore( source )
  bench :=  LazyFSBenchmark{
    BufSize: 1024,
  }

  bench.Run( store, b )
}
