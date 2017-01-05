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


func TestHttpSourceSparseStore( t *testing.T ) {

  source,err := OpenHttpSource(TestUrlRoot, AlphabetPath)
  if err != nil {
    t.Fatal("Couldn't open HttpFSSource")
  }

  store,err :=  OpenSparseFileStore( src, HttpSourceSparseStore )



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
