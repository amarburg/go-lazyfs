package lazyfs

import "testing"

import "fmt"

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



func TestHttpFSListing( t *testing.T ) {
  fs, err := OpenHttpFSSource( OOIRawDataRootURL )

  if fs == nil || err != nil {
    t.Fatal("Couldn't open HTTPFSSource to OOI Raw Data Server", err.Error() )
  }


  listing, err := fs.ReadHttpDir( "/" )

  if err != nil {
    t.Fatal("Couldn't list directory on OOI Raw Data Server", err.Error() )
  }

  fmt.Printf("root contains %d entries\n", len( listing.Children ))

}
