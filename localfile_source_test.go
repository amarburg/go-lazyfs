package lazyfs

import "testing"
import "io"


func TestLocalFileSource(t *testing.T) {
  fs,err := OpenLocalFileSource( LocalFilesRoot, AlphabetPath )

  //if fs == nil {
  //  t.Error("FileStore doesn't exist")
  //}

  if err != nil {
    t.Fatal("Error on opening LocalFileSource for", LocalFilesRoot+AlphabetPath)
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
