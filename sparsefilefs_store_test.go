package lazyfs

import "testing"
import "os"


func TestSparseFileFSStore( t *testing.T ) {

  os.Mkdir( SparseFileRoot, 0755 )

  fs, err := OpenSparseFileFSStore( SparseFileRoot )

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
