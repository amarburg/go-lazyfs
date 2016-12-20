package lazyfs

import (
  "testing"
  "reflect"
)


func CheckTestFile( buf []byte, off int64 ) bool {
  // I'm sure there's a beautiful idiomatic Go way to do this
  str := [26]byte{ 65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
                   75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
                   85, 86, 87, 88, 89, 90 }
  return reflect.DeepEqual(buf, str[off:off+int64(cap(buf))] )
}

func TestFileStore(t *testing.T) {
  fs := OpenFileStore("test_files/foo.fs")

  if fs == nil {
    t.Error("fs doesn't exist")
  }

  buf := make([]byte,10)
  n,err := fs.ReadAt(buf, 0)

  if err != nil {
    t.Errorf("Error on read: %s", err.Error() )
  }

  if n != 10 {
    t.Errorf("Couldn't read %d != 10 bytes from test file", n)
  }

  if !CheckTestFile(buf,0) {
    t.Errorf("\"%s\" doesn't match test file at %d", n, 0)
  }

  h,err := fs.HasAt( buf, 0 )
  if err != nil {
    t.Errorf("Error on HasAt: %s", err.Error() )
  }

  if h != 10 {
    t.Errorf("Should have %d != 10 bytes at offset 0", h)
  }

  if !CheckTestFile(buf,0) {
    t.Errorf("\"%s\" doesn't match test file at %d", h, 0)
  }

}


func TestFileStoreInLazyFS( t *testing.T ) {
  fs := OpenFileStore("test_files/foo.fs")
  if fs == nil { t.Errorf("Couldn't open FileStore") }

  lazyfs := LazyFS { handler: fs }

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
