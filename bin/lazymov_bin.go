package main

import "fmt"
//import "io"
import "github.com/amarburg/lazyfs"

var TestUrlRoot = "http://localhost:8080/files/"
var TestMovPath = "CamHD_Vent_Short.mov"

var SparseHttpStoreRoot = "test_files/httpsparse/"

func main() {

  source,err := lazyfs.OpenHttpFSSource(TestUrlRoot)
  if err != nil {
    panic("Couldn't open HttpFSSource")
  }

  store,err := lazyfs.OpenSparseFileFSStore( SparseHttpStoreRoot )
  if store == nil {
    panic("Couldn't open SparesFileFSStore")
  }

  source.SetBackingStore( store )

  file,err := source.Open( TestMovPath )
  if err != nil {
    panic("Couldn't open AlphabetPath")
  }

  buf := make( []byte, 10 )
  n,err := file.ReadAt( buf, 0 )

  fmt.Printf("Read %d characters: %s\n", n, buf)


}
