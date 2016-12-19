package lazyfs

import (
  "testing"
)

func TestFileStore(t *testing.T) {
  fs := OpenFileStore("testfile/foo.fs")

  if fs == nil {
    t.Error("fs doesn't exist")
  }

}
