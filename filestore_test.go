package lazyfs

import (
  "testing"
  "reflect"
  "fmt"
  "io"
)



var BufSize = 10

var test_pairs = []struct {
  offset int64
  length int
}{
  {0,  BufSize},
  {10, BufSize},
  {20, 6},
}


func CheckTestFile( buf []byte, off int64 ) bool {
  // I'm sure there's a beautiful idiomatic Go way to do this
  test_string := [26]byte{ 65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
    75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
    85, 86, 87, 88, 89, 90 }

    l := int(off)+len(buf)
    return reflect.DeepEqual(buf, test_string[int(off):l] )
  }

  func TestFileStore(t *testing.T) {
    fs := OpenFileStore("test_files/foo.fs")

    if fs == nil {
      t.Error("fs doesn't exist")
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


  func TestFileStoreInLazyFile( t *testing.T ) {
    fs := OpenFileStore("test_files/foo.fs")
    if fs == nil { t.Errorf("Couldn't open FileStore") }

    lazyfs := LazyFile { storage: fs }

    buf := make([]byte,10)
    n,err := lazyfs.ReadAt(buf, 0)

    if err != nil {
      t.Errorf("Error on read: %s", err.Error() )
    }

    if n != 10 {
      t.Errorf("Couldn't read %d != 10 bytes from test file", n)
    }

    if !CheckTestFile(buf,0) {
      t.Errorf("\"%s\" doesn't match test file at %d", n, 0)
    }

  }
