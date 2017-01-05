package lazyfs

import "os"
import "errors"
import "io"
import "fmt"
import "path/filepath"

type SparseFileFSStore struct {
  root  string
}

func OpenSparseFileFSStore( root string ) (*SparseFileFSStore, error) {
  fs := SparseFileFSStore{ root: root }

  // Check for validity of root
  f,err := os.Stat( root )

  if err != nil {
    return &fs, err
  }

  if !f.IsDir() {
    return &fs, errors.New("Is not a directory")
  }

  return &fs, nil
}


func (fs *SparseFileFSStore ) Store( source FileSource ) (FileStorage, error) {
  sparsefile := fs.root + source.Path()

  fmt.Println("Checking sparsefile ",sparsefile)

  // Fill file with with null
  sz,_ := source.FileSize()

  fileinfo,err := os.Stat( sparsefile )
  if err != nil || fileinfo.Size() != sz {

    fmt.Println("Creating sparsefile size", sparsefile, sz)
    os.MkdirAll( filepath.Dir(sparsefile), 0755 )
    dest,err := os.Create( sparsefile )

    if err != nil {
      panic(fmt.Sprintf("Couldn't create sparsefile %s", sparsefile) )
    }

    zero := &ZeroReader{ size: sz }
    io.Copy( dest, zero )
  }

  file,err := OpenSparseFileStore( sparsefile )

  return file,err
}



type ZeroReader struct {
  size int64
}

func (w *ZeroReader) Read( p []byte) (n int, err error) {
  if int64(cap(p)) > w.size {
    n = int(w.size)
  } else {
    n = cap(p)
  }

  for i := 0; i < n; i++ { p[i] = 0 }

  w.size -= int64(n)

  if w.size == 0 {
    err = io.EOF
    } else {
      err = nil
    }

    return n,err
}
