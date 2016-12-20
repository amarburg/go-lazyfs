package lazyfs

import (
	"os"
)

type FileStoreError struct {
	Err string
}

func (e FileStoreError) Error() string {
	return e.Err
}

type FileStore struct {
	file *os.File
}

func OpenFileStore( name string ) *FileStore {
	f,err := os.Open(name)
	if err != nil {
		return nil
	}
	fs := FileStore{ file: f }
	return &fs
}

func (fs *FileStore) ReadAt( p []byte, off int64) (n int, err error) {
	return fs.file.ReadAt( p, off )
}

func (fs *FileStore) WriteAt(p []byte, off int64) (n int, err error) {
	return 0,nil
}




type HasAt interface {
	HasAt( p []byte, off int64 ) (n int, err error)
}

func (fs *FileStore) HasAt( p []byte, off int64 ) (n int, err error) {
	len := cap( p )
	sz := fs.FileSize()

	switch {
		case (off + int64(len)) < sz: return len, nil
		case off > sz: return 0, FileStoreError{"Offset beyond end of file"}
		case off+int64(len) > sz: return int(off+int64(len) - sz), nil
	}

	return 0, FileStoreError{"Shouldn't get here"}
}

func (fs *FileStore) FileSize() int64 {
	stat,_ := fs.file.Stat()
	return stat.Size()
}
