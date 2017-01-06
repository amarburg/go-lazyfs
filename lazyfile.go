package lazyfs

import "io"

type File interface {
  io.ReaderAt
  FileSize() (int64,error)
}

type FileSource interface {
  File
  Reader() (io.Reader)
  Path() string
}

type FileStorage interface {
  File
  //HasAt( p []byte, off int64 ) (int, error)
}
