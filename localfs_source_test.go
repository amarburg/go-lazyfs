package lazyfs

import "testing"


func TestLocalFSSource( t *testing.T ) {

  fs, err := OpenLocalFSSource( LocalFilesRoot )

  if fs == nil || err != nil {
    t.Fatal("Couldn't create LocalFSSource", err.Error() )
  }


  alphabet,err := fs.Open( AlphabetPath )

  if alphabet == nil || err != nil {
    t.Fatal("Couldn't create FileSource from LocalFSSource", err.Error() )
  }


}
