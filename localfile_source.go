package lazyfs

import "os"
import "io"
import "path"
//import "fmt"

type LocalFileSource struct {
  root, path string
  file *os.File
  //store FileStorage
}

func OpenLocalFile( fpath string ) (fsrc *LocalFileSource, err error ) {
  root,file := path.Split(fpath)
  return OpenLocalFileSource( root, file )
}

func OpenLocalFileSource( root string, ff string ) (fsrc *LocalFileSource, err error ) {
  f,err := os.Open( path.Join(root,ff) )

  // TODO:  Check for non-existent files

  return &LocalFileSource{ root:root, path: ff, file: f }, err
}

// func (fs *LocalFileSource) SetBackingStore( store FileStorage ) {
// 	fs.store = store
// }


func (fs *LocalFileSource) ReadAt( p []byte, off int64 ) (n int, err error) {
  // if fs.store != nil {
  //   //fmt.Println("Checking store")
  //
  //   if _,err := fs.store.HasAt(p,off); err == nil  {
  //     fmt.Println("Retrieving from store")
  //     return fs.store.ReadAt( p, off )
  //   } else {
  //     fmt.Println( "Need to update store")
  //     n,_ := fs.file.ReadAt(p,off)
  //     fs.store.WriteAt(p[:n], off)
  //
  //     return n, nil
  //   }
  // } else {
    return fs.file.ReadAt(p,off)
//  }
}

func (fs *LocalFileSource) FileSize() (int64,error) {
	stat,_ := fs.file.Stat()
	return stat.Size(), nil
}

func (fs *LocalFileSource) Reader() io.Reader {
  f,_ := os.Open( fs.root + fs.path )
  return f
}

func (fs *LocalFileSource) Path() string {
  return fs.path
}
