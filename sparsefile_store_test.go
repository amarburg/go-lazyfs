package lazyfs

import "testing"

func TestSparseFileStore(t *testing.T) {
  fs,err := OpenSparseFileStore( LocalAlphabetPath )

  if fs == nil {
    t.Fatal("SparseFileStore doesn't exist")
  }

  if err != nil {
    t.Fatal("Error opening SparseFileStore", err.Error() )
  }

}
