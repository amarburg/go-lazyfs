package lazyfs


import ( "io" )

type FSStorage interface {
  io.ReaderAt
  io.WriterAt
  HasAt
}

type FSSource interface {
  io.ReaderAt
}

type LazyFS struct {
  storage FSStorage
  source FSSource
}

func (fs *LazyFS) ReadAt( p []byte, off int64 ) (n int, err error) {
  n,err = fs.storage.HasAt( p, off )

	if err == nil {
    return fs.storage.ReadAt( p, off )
  } else {
    if fs.source == nil { return 0, LazyFSError { "Source not defined" } }

    n,err = fs.source.ReadAt( p, off )

    if cap(p) != n { return n, LazyFSError { "Short read!"}}

    _,err = fs.storage.WriteAt( p, off )

    return n,err
  }
}
