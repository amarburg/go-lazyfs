package lazyfs

import "os"
import "errors"


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


func (fs *SparseFileFSStore ) Open( path string ) (*SparseFileStore, error) {
  return OpenSparseFileStore( fs.root + path )
}
