package lazyfs

import "os"
import "io"

type LocalFileSource struct {
  root, path string
  file *os.File
}

func OpenLocalFileSource( root string, path string ) (fsrc *LocalFileSource, err error ) {
  f,err := os.Open(root+path)

  // TODO:  Check for non-existent files

  return &LocalFileSource{ root:root, path: path, file: f }, err
}

func (fs *LocalFileSource) ReadAt( p []byte, off int64 ) (n int, err error) {
  return fs.file.ReadAt(p,off)
}

func (fs *LocalFileSource) FileSize() (int64,error) {
	stat,_ := fs.file.Stat()
	return stat.Size(), nil
}

func (fs *LocalFileSource) Reader() io.Reader {
  f,_ := os.Open( fs.root + fs.path)
  return f
}

func (fs *LocalFileSource) Path() string {
  return fs.path
}
