package lazyfs_benchmarking

import "testing"
import "github.com/amarburg/go-lazyfs"
import "net/url"

var TestUrlRoot = "https://raw.githubusercontent.com/amarburg/go-lazyfs-testfiles/master/"
var AlphabetPath = "alphabet.fs"
var AlphabetUrl,_ = url.Parse( TestUrlRoot + AlphabetPath )

var HttpSourceSparseStore = "test_cache/httpsource_sparsestore/"

var BufSize = 10

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

  //

  //
  // for _,test := range test_pairs {
  //
  //   buf := make([]byte,BufSize)
  //
  //   // Test ReadAt
  //   n,err := store.ReadAt(buf, test.offset)
  //
  //   if err != nil && err != io.EOF {
  //     t.Errorf("Error on read: %s", err.Error() )
  //   }
  //
  //   if n != test.length {
  //     t.Error("Expected",test.length,"bytes, got",n)
  //   }
  //
  //   buf = buf[:n]
  //
  //   if !CheckTestFile(buf, test.offset) {
  //     t.Errorf("\"%s\" doesn't match test file at %d", n, 0)
  //   }
  //
  // }

}
