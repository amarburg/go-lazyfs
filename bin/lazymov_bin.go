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
  var indent = 0

  sz,_ := file.FileSize()
  ParseAtom( file, offset, sz, indent )



}


func ParseAtom( file lazyfs.FileSource, offset int64, length int64, indent int ) error {
  for offset < length {

    header_buf := make([]byte, quicktime.AtomHeaderLength )
    n,err := file.ReadAt(header_buf, offset)

    if err == io.EOF {
      fmt.Println("End of file")
      return err
    }

    if n != quicktime.AtomHeaderLength {
      panic(fmt.Sprintf("Tried to read an Atom header and got %d instead",n))
    }

    header,err := quicktime.ParseAtomHeader( header_buf )

    if err != nil {
      panic("Error parsing header")
    }

    for i := 0; i < indent; i++ { fmt.Printf("  ") }
    fmt.Printf("%v  %v\n",header.Type, header.Size )

    atom_end := offset + int64(header.Size)

    if( header.IsContainer() ) {
      err = ParseAtom( file, offset+quicktime.AtomHeaderLength, atom_end, indent+1 )
      if err != nil { return err }
    }

    offset = atom_end
  }

  return nil
}
