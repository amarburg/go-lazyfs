package lazyfs

import "testing"
import "os"
import "io"

func ClearSparseStoreRoot() {
  os.RemoveAll( SparseStoreRoot )
}

func TestSparseFileFSStore( t *testing.T ) {

  ClearSparseStoreRoot()
  os.MkdirAll( SparseStoreRoot, 0755 )

  fs, err := OpenSparseFileFSStore( SparseStoreRoot )

  if fs == nil || err != nil {
    t.Fatal("Couldn't create SparseFileFSStore", err.Error() )
  }

  source,err := OpenLocalFileSource( LocalFilesRoot, AlphabetPath )
  alphabet,err := fs.Store( source )

  if alphabet == nil  {
    if err != nil {
      t.Fatal("Couldn't create SparseFileStore from SparseFileFSStore", err.Error() )
    } else {
      t.Fatal("Couldn't create SparseFileStore from SparseFileFSStore; no error")
    }
  }


    dest,err := os.Stat( SparseStoreRoot + AlphabetPath )

    if err != nil {
      t.Fatal("Couldn't find sparse file", SparseStoreRoot + AlphabetPath )
    }

    if dest.Size() != AlphabetSize {
      t.Errorf("Storage file isn't the expected length %d != %d", dest.Size(), AlphabetSize )
    }


}


func TestZeroReader( t *testing.T ) {
  z := ZeroReader{ size: 36 }
  p := make( []byte, 16)

  l,err := z.Read( p )
  if err != nil || l != 16 {
    t.Error("Didn't work (first)",l,err)
  }

  l,err = z.Read( p )
  if err != nil || l != 16 {
    t.Error("Didn't work (second)",l,err)
  }

  l,err = z.Read( p )
  if err != io.EOF || l != 4 {
    t.Error("Didn't work (third)",l,err)
  }
}


// func TestSparseFileFSStoreBadPath( t *testing.T ) {
//
//   _, err := OpenSparseFileFSStore( LocalBadPath )
//
//   if err == nil {
//     t.Error("Expected error" )
//   }
//
//
//   // This is an existing file
//   _, err = OpenSparseFileFSStore( LocalAlphabetPath )
//
//   if err == nil {
//     t.Error("Expected error")
//   }
//
// }
