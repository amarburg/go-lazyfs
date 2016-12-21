package lazyfs

import "testing"


func TestLocalFSStore( t *testing.T ) {

  fs, err := OpenLocalFSStore( LocalFilesRoot )

  if fs == nil || err != nil {
    t.Fatal("Couldn't create LocalFSStore", err.Error() )
  }


  alphabet,err := fs.Open( AlphabetPath )

  if alphabet == nil || err != nil {
    t.Fatal("Couldn't create FileStore from LocalFSStore", err.Error() )
  }


}
