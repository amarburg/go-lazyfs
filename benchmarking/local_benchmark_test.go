package lazyfs_benchmarking

import "testing"
import "github.com/amarburg/go-lazyfs"
import "github.com/amarburg/go-lazyfs-testfiles"
import "github.com/amarburg/go-lazyfs-testfiles/http_server"

import "math/rand"
import "net/url"
//import "fmt"


func BenchmarkLocalHttpSource( b *testing.B ) {
  BufSize := 1024
  AlphabetUrl,_ := url.Parse( "http://localhost:4567/" + lazyfs_testfiles.TenMegBinaryFile )
  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  source,err := lazyfs.OpenHttpSource( *AlphabetUrl )


  if err != nil {
    panic("Couldn't open HttpFSSource")
  }

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    offset := rand.Intn( lazyfs_testfiles.TenMegFileLength - BufSize )

      buf := make([]byte,BufSize)
      bytesPrev := source.Stats.ContentBytesRead

      // Test ReadAt
      n,_ := source.ReadAt(buf, int64(offset))
      //fmt.Println(n,err)
      if n != BufSize { panic("bad read")}

      b.SetBytes( int64(source.Stats.ContentBytesRead - bytesPrev) )

  }
  b.StopTimer()

  //if b.N > 1 {
  //  fmt.Printf("Read %d bytes of content over HTTP\n", source.Stats.ContentBytesRead)
  //}


}

func BenchmarkLocalHttpSourceSparseStore( b *testing.B ) {
  BufSize := 1024
  HttpSourceSparseStore := "test_cache/local_benchmark/"
  SourceUrl,_ := url.Parse( "http://localhost:4567/" + lazyfs_testfiles.TenMegBinaryFile )
  srv := lazyfs_testfiles_http_server.HttpServer( 4567 )
  defer srv.Stop()

  source,err := lazyfs.OpenHttpSource( *SourceUrl )
  if err != nil {
    panic("Couldn't open HttpFSSource")
  }

  store,err :=  lazyfs.OpenSparseFileStore( source, HttpSourceSparseStore )
  buf := make([]byte,BufSize)

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    bytesPrev := source.Stats.ContentBytesRead

    offset := rand.Intn( lazyfs_testfiles.TenMegFileLength - BufSize )

      n,_ := store.ReadAt(buf, int64(offset) )
      if n != BufSize { panic("bad read")}

      b.SetBytes( int64(source.Stats.ContentBytesRead - bytesPrev) )

  }
  b.StopTimer()


  //if b.N > 1  {
  //  fmt.Printf("Read %d bytes of content over HTTP\n", source.Stats.ContentBytesRead)
  //}

}
