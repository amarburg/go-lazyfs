package lazyfs

import "testing"
import "net/url"
import "io"

func TestSparseFileStore(t *testing.T) {
  src,_ := OpenLocalFileSource( LocalFilesRoot, AlphabetPath )
  fs,err := OpenSparseFileStore( src, SparseStoreRoot )

  if fs == nil {
    t.Fatal("SparseFileStore doesn't exist")
  }

  if err != nil {
    t.Fatal("Error opening SparseFileStore", err.Error() )
  }

  for _,test := range test_pairs {

    buf := make([]byte,BufSize)

    // Test ReadAt
    n,err := fs.ReadAt(buf, test.offset)

    if err != nil && err != io.EOF {
      t.Errorf("Error on read: %s", err.Error() )
    }

    if n != test.length {
      t.Error("Expected",test.length,"bytes, got",n)
    }

    buf = buf[:n]

    if !CheckTestFile(buf, test.offset) {
      t.Errorf("\"%s\" doesn't match test file at %d", n, 0)
    }

  }

}


func TestHttpSourceSparseStore( t *testing.T ) {

  src_url,_ := url.Parse( TestUrlRoot+AlphabetPath)
  source,err := OpenHttpSource( *src_url )
  if err != nil {
    t.Fatal("Couldn't open HttpFSSource")
  }

  store,err :=  OpenSparseFileStore( source , HttpSourceSparseStore )


  for _,test := range test_pairs {

    buf := make([]byte,BufSize)

    // Test ReadAt
    n,err := store.ReadAt(buf, test.offset)

    if err != nil && err != io.EOF {
      t.Errorf("Error on read: %s", err.Error() )
    }

    if n != test.length {
      t.Error("Expected",test.length,"bytes, got",n)
    }

    buf = buf[:n]

    if !CheckTestFile(buf, test.offset) {
      t.Errorf("\"%s\" doesn't match test file at %d", n, 0)
    }

  }

}
