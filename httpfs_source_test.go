package lazyfs

import "testing"


func TestHttpFSSource( t *testing.T ) {

  fs, err := OpenHttpFSSource( TestUrlRoot )

  if fs == nil || err != nil {
    t.Fatal("Couldn't create HttpFSSource", err.Error() )
  }


  alphabet,err := fs.Open( AlphabetPath )

  if alphabet == nil || err != nil {
    t.Fatal("Couldn't create FileSource from HttpFSSource", err.Error() )
  }


}
