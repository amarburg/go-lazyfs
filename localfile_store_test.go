package lazyfs

import (
  "testing"
  "io"
)

// Test actually uses the source files...
func TestLocalFileStore(t *testing.T) {
  fs,_ := OpenLocalFileStore( LocalFilesRoot + AlphabetPath )

  if fs == nil {
    t.Fatal("LocalFileStore doesn't exist")
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

    // Test HasAt
    h,err := fs.HasAt( buf, test.offset )
    if err != nil {
      t.Errorf("Error on HasAt: %s", err.Error() )
    }

    if h != test.length {
      t.Error("Expected",test.length,"bytes, got",h)
    }

  }

}
