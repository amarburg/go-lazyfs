package lazyfs

import "testing"

func TestSparseFileStore(t *testing.T) {
  src,_ := OpenLocalFileSource( LocalFilesRoot, AlphabetPath )
  fs,err := OpenSparseFileStore( src, SparseStoreRoot )

  if fs == nil {
    t.Fatal("SparseFileStore doesn't exist")
  }

  if err != nil {
    t.Fatal("Error opening SparseFileStore", err.Error() )
  }

  // Should fail
  buf := make([]byte,10)
  n,err := fs.HasAt( buf, 0 )

  if n != 0 || err == nil {
    t.Error("HasAt should have failed, but didn't")
  }

  n,err = fs.ReadAt( buf, 0 )

  if n != 0 || err == nil {
    t.Error("ReadAt should have failed, but didn't")
  }

}
