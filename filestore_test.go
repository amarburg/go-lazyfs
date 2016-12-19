package lazyfs

import (
  "testing"
)

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

  h,err := fs.HasAt( buf, 0 )
  if err != nil {
    t.Errorf("Error on HasAt: %s", err.Error() )
  }

  if h != 10 {
    t.Errorf("Should have %d != 10 bytes at offset 0", h)
  }

}
