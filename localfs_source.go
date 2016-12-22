package lazyfs

import "fmt"


type LocalFSSource struct {
  root string
  store FSStorage
}

func OpenLocalFSSource( root string ) (*LocalFSSource, error) {
  fs := LocalFSSource{ root: root, store: nil }
  return &fs, nil
}

func (fs *LocalFSSource) SetBackingStore( store FSStorage ) {
  fs.store = store
}

func (fs *LocalFSSource) Open( path string ) (*LocalFileSource, error) {
  file,err := OpenLocalFileSource( fs.root, path )

  if fs.store != nil {
      st,err := fs.store.Store( file )

      if err != nil {
        panic(fmt.Sprintf("Couldn't create store for file %s",path))
      }

      file.SetBackingStore( st )
  }
  return file, err
}
