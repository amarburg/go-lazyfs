package main

import "fmt"
import "io"
import "github.com/amarburg/lazyfs"
import "github.com/amarburg/go-quicktime"

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

    var offset int64 = 0

  for {

    header_buf := make([]byte, quicktime.AtomHeaderLength )
    n,err := file.ReadAt(header_buf, offset)

    if err == io.EOF {
      fmt.Println("End of file")
      break
    }

    if n != quicktime.AtomHeaderLength {
      panic(fmt.Sprintf("Tried to read an Atom header and got %d instead",n))
    }

    header,err := quicktime.ParseAtomHeader( header_buf )

    if err != nil {
      panic("Error parsing header")
    }

    fmt.Println(header.Size, header.DataSize, header.Type )

    offset += int64(header.Size)

  }




}
