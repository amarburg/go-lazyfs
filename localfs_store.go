package lazyfs

import "path/filepath"
import "os"
import "io"

type LocalFSStore struct {
  root  string
}

func OpenLocalFSStore( root string ) (*LocalFSStore, error) {
  fs := LocalFSStore{ root: root }
  return &fs, nil
}

func (fs *LocalFSStore) Store( source FileSource ) (*LocalFileStore, error) {
  file := fs.root + source.Path()

  _,err := os.Stat( file )
  if err != nil {

    os.MkdirAll( filepath.Dir( file ), 0755 )

    // Copy file over
    dest,_ := os.Create( file )
    io.Copy( dest, source.Reader() )

  }

  return OpenLocalFileStore( file )
}
