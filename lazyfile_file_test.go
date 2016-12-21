package lazyfs

import "testing"
import "io"

// This is actually a null test because FileStore is always complete (never
//  needs to query the source)
func TestLazyFileFileSourceAndStorage( t *testing.T ) {
  store := OpenFileStore(LocalAlphabetPath)
  if store == nil {
    t.Fatal("Couldn't open FileStore")
  }

  source,err := OpenFileSource(LocalAlphabetPath)
  if source == nil || err != nil {
    t.Fatal("Couldn't open FileSource")
  }

  lazyfs := LazyFile { storage: store, source: source }

  for _,test := range test_pairs {

    buf := make([]byte,BufSize)
    n,err := lazyfs.ReadAt(buf, test.offset)

    if err != nil && err != io.EOF {
      t.Errorf("Error on read: %s", err.Error() )
    }

    if n != test.length {
      t.Error("Expected",test.length,"bytes, got",n)
    }

    buf = buf[:n]

    if !CheckTestFile(buf,test.offset) {
      t.Errorf("\"%s\" doesn't match test file at %d", n, 0)
    }

  }

}
