package lazyfs

import "os"


type LocalFileStoreError struct {
	Err string
}

func (e LocalFileStoreError) Error() string {
	return e.Err
}

type LocalFileStore struct {
	file *os.File
}

func OpenLocalFileStore( name string ) (*LocalFileStore, error) {
	f,err := os.Open(name)
	fs := LocalFileStore{ file: f }
	return &fs, err
}

func (fs *LocalFileStore) ReadAt( p []byte, off int64) (n int, err error) {
		return fs.file.ReadAt( p, off )
}

func (fs *LocalFileStore) WriteAt(p []byte, off int64) (n int, err error) {
	return 0,nil
}

func (fs *LocalFileStore) HasAt( p []byte, off int64 ) (n int, err error) {
	len := int64(cap( p ))
	sz,_ := fs.FileSize()

	switch {
		case (off + len) < sz: return int(len), nil
		case off > sz: return 0, LocalFileStoreError{"Offset beyond end of file"}
		case (off + len) > sz: return int(sz - off), nil
	}

	return 0, LocalFileStoreError{"Shouldn't get here"}
}

func (fs *LocalFileStore) FileSize() (int64,error) {
	stat,_ := fs.file.Stat()
	return stat.Size(),nil
}
