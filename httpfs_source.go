package lazyfs

import "fmt"

type HttpFSSource struct {
  url_root string
  store FSStorage
}

func OpenHttpFSSource( url_root string ) (*HttpFSSource, error) {
  fs := HttpFSSource{ url_root: url_root }
  return &fs, nil
}

func (fs *HttpFSSource) SetBackingStore( store FSStorage ) {
  fs.store = store
}

func (fs *HttpFSSource ) Open( path string ) (*HttpSource, error) {
  src,err := OpenHttpSource( fs.url_root, path )

  if fs.store != nil {
      st,err := fs.store.Store( src )

      if err != nil {
        panic(fmt.Sprintf("Couldn't create store for file %s",path))
      }

      src.SetBackingStore( st )
  }

  return src,err
}
