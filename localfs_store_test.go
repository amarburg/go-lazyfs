package lazyfs

import "testing"
import "os"

func ClearLocalStoreRoot() {
  os.RemoveAll( LocalStoreRoot )
}

func TestLocalFSStore( t *testing.T ) {

  ClearLocalStoreRoot()

  fs, err := OpenLocalFSStore( LocalStoreRoot )

  if fs == nil || err != nil {
    t.Fatal("Couldn't create LocalFSStore", err.Error() )
  }

  source,_ := OpenLocalFileSource( LocalFilesRoot, AlphabetPath )

  alphabet,err := fs.Store( source )

  if alphabet == nil || err != nil {
    t.Fatal("Couldn't create FileStore from LocalFSStore", err.Error() )
  }

  dest,_ := os.Stat( LocalStoreRoot + AlphabetPath )

  if dest.Size() != AlphabetSize {
    t.Errorf("Storage file isn't the expected length %d != %d", dest.Size(), AlphabetSize )
  }

}
