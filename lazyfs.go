package lazyfs


import ( "io" )

type FSHandler interface {
  io.ReaderAt
}

type LazyFS struct {
  handler FSHandler
}

func (fs *LazyFS) ReadAt( p []byte, off int64) (n int, err error) {
	return fs.handler.ReadAt( p, off )
}
