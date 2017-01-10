package lazyfs_benchmarking

import "testing"
import "github.com/amarburg/go-lazyfs"
import "github.com/amarburg/go-lazyfs-testfiles"
import "net/url"
import "fmt"

var TestUrlRoot = "https://raw.githubusercontent.com/amarburg/go-lazyfs-testfiles/master/"
var AlphabetUrl,_ = url.Parse( TestUrlRoot + lazyfs_testfiles.AlphabetFile )

var HttpSourceSparseStore = "test_cache/github_benchmark/"

func BenchmarkGithubHttpSource( b *testing.B ) {

  source,err := lazyfs.OpenHttpSource( *AlphabetUrl )

  if err != nil {
    panic("Couldn't open HttpFSSource")
  }

  b.ResetTimer()
  for i := 0; i < b.N; i++ {
      buf := make([]byte,BufSize)

      // Test ReadAt
      n,err := source.ReadAt(buf, 0)
      if n != BufSize || err != nil { panic("bad read")}
  }

  if b.N > 1 {
  fmt.Printf("Read %d bytes of content over HTTP\n", source.Stats.ContentBytesRead)
}
}

func BenchmarkGithubHttpSourceSparseStore( b *testing.B ) {

  source,err := lazyfs.OpenHttpSource( *AlphabetUrl )
  if err != nil {
    panic("Couldn't open HttpFSSource")
  }

  store,err :=  lazyfs.OpenSparseFileStore( source, HttpSourceSparseStore )
  buf := make([]byte,BufSize)


  b.ResetTimer()
  for i := 0; i < b.N; i++ {
      n,err := store.ReadAt(buf, 0)
      if n != BufSize || err != nil { panic("bad read")}
  }

  if b.N > 1  {
  fmt.Printf("Read %d bytes of content over HTTP\n", source.Stats.ContentBytesRead)
}

}
