package main

import "fmt"
//import "io"
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

  // var offset int64 = 0
  // var indent = 0

  sz,_ := file.FileSize()
//  ParseAtom( file, offset, sz, indent )

  tree := quicktime.BuildTree( file, sz )


  DumpTree( file, tree )

}


// func ParseAtom( file lazyfs.FileSource, offset int64, length int64, indent int ) error {
//   for offset < length {
//
//     // header_buf := make([]byte, quicktime.AtomHeaderLength )
//     // n,err := file.ReadAt(header_buf, offset)
//     //
//     // if err == io.EOF {
//     //   fmt.Println("End of file")
//     //   return err
//     // }
//     //
//     // if n != quicktime.AtomHeaderLength {
//     //   panic(fmt.Sprintf("Tried to read an Atom header and got %d instead",n))
//     // }
//     //
//     // header,err := quicktime.ParseAtomHeader( header_buf )
//
//     header,err := quicktime.ParseAtomHeaderAt( file, offset )
//
//     if err != nil {
//       panic("Error parsing header")
//     }
//
//     PrintAtom( file, header, indent )
//
//     atom_end := offset + int64(header.Size)
//
//     if( header.IsContainer() ) {
//       err = ParseAtom( file, offset+quicktime.AtomHeaderLength, atom_end, indent+1 )
//       if err != nil { return err }
//     }
//
//     offset = atom_end
//   }
//
//   return nil
// }


func DumpTree( file lazyfs.FileSource, tree []quicktime.AtomHeader ) {
  indent := 0
  for _,atom := range tree  {
    PrintAtom( file, atom, indent )
  }
}

func PrintAtom( file lazyfs.FileSource, header quicktime.AtomHeader, indent int ){

    for i := 0; i < indent; i++ { fmt.Printf("..") }
    fmt.Printf("%v  %v",header.Type, header.Size )

    switch header.Type {
    case "ftyp":
      atom,_ :=  quicktime.ReadAtomAt( file, header )
      ftyp,_ := quicktime.ParseFTYP( atom )

      fmt.Printf(" (%08x %08x) ", ftyp.MajorBrand, ftyp.MinorVersion )
    case "stco":
      atom,_ :=  quicktime.ReadAtomAt( file, header )
      stco,_ := quicktime.ParseSTCO( atom )

      fmt.Printf(" (%d entries) ", len(stco.ChunkOffsets) )
    case "stsz":
      atom,_ :=  quicktime.ReadAtomAt( file, header )
      stsz,_ := quicktime.ParseSTSZ( atom )

      fmt.Printf(" (%d entries, default size %d)", len(stsz.SampleSizes), stsz.SampleSize )
    case "stsc":
      atom,_ :=  quicktime.ReadAtomAt( file, header )
      stsz,_ := quicktime.ParseSTSC( atom )

      fmt.Printf(" (%d entries): ", len(stsz.Entries) )
      for _,e := range stsz.Entries  {
        fmt.Printf(" (%d,%d,%d) ", e.FirstChunk, e.SamplesPerChunk, e.SampleId )
      }
    }


    fmt.Printf("\n")


    for _,child := range header.Children  {
      PrintAtom( file, child, indent+1 )
    }

}
