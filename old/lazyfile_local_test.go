package lazyfs

import "testing"
//import "fmt"
import "io"

func TestLazyFileLocalStorage( t *testing.T ) {

  source,err := OpenLocalFSSource(LocalFilesRoot)
  if err != nil {
    t.Fatal("Couldn't open FileSource")
  }

  store,err := OpenLocalFSStore( LocalStoreRoot )
  if store == nil {
    t.Fatal("Couldn't open LocalFSStore")
  }

  source.SetBackingStore( store )

  file,err := source.Open( AlphabetPath )
  if err != nil {
    t.Fatal("Couldn't open AlphabetPath")
  }

  for _,test := range test_pairs {

    buf := make([]byte,BufSize)

    // Test ReadAt
    n,err := file.ReadAt(buf, test.offset)

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
