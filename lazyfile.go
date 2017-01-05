package lazyfs

import "io"

type FileSource interface {
  FileSize() (int64,error)
  Reader() (io.Reader)
  Path() string
  io.ReaderAt
}

type FileStorage interface {
  io.ReaderAt
  //io.WriterAt
  HasAt( p []byte, off int64 ) (int, error)
  FileSize() (int64,error)
}

// type LazyFile struct {
//   storage FileStorage
//   source FileSource
// }
//
// func (fs *LazyFile) ReadAt( p []byte, off int64 ) (n int, err error) {
//   n,err = fs.storage.HasAt( p, off )
//
// 	if err == nil {
//     return fs.storage.ReadAt( p, off )
//   } else {
//     if fs.source == nil { return 0, LazyFSError { "Source not defined" } }
//
//     n,err = fs.source.ReadAt( p, off )
//
//     if cap(p) != n { return n, LazyFSError { "Short read!"}}
//
//     _,err = fs.storage.WriteAt( p, off )
//
//     return n,err
//   }
// }
