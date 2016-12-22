package lazyfs

import "testing"
import "fmt"

func TestLazyFileSparseStorage( t *testing.T ) {

  store,err := OpenSparseFileFSStore(SparseStoreRoot)
  if store == nil {
    t.Fatal("Couldn't open SparesFileFSStore")
  }

  source,err := OpenLocalFSSource(LocalFilesRoot)
  if source == nil || err != nil {
    t.Fatal("Couldn't open FileSource")
  }

  lastfs := struct {
    storage  *SparseFileFSStore
    source   *LocalFSSource
  } { storage: store, source: source }


  fmt.Println( lastfs )

}

//
//   store.Store( source )
//
//   lazyfs := CreateLazyFile { storage: store, source: source }
//
//   for _,test := range test_pairs {
//
//     buf := make([]byte,BufSize)
//     n,err := lazyfs.ReadAt(buf, test.offset)
//
//     if err != nil && err != io.EOF {
//       t.Errorf("Error on read: %s", err.Error() )
//     }
//
//     if n != test.length {
//       t.Error("Expected",test.length,"bytes, got",n)
//     }
//
//     buf = buf[:n]
//
//     if !CheckTestFile(buf,test.offset) {
//       t.Errorf("\"%s\" doesn't match test file at %d", n, 0)
//     }
//
//   }
//
// }
