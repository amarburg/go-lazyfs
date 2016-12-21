package lazyfs

import "os"

type FileSource struct {
  path string
  file *os.File
}

func OpenFileSource( path string ) (fsrc FileSource, err error ) {
  f,err := os.Open(path)

// TODO:  Check for non-existent files

  return FileSource{ path: path, file: f }, err
}

func (fs *FileSource) ReadAt( p []byte, off int64 ) (n int, err error) {
  return fs.file.ReadAt(p,off)
}

func (fs *FileSource) FileSize() (int64,error) {
	stat,_ := fs.file.Stat()
	return stat.Size(), nil
}
