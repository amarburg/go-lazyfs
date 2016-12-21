package lazyfs

import "os"


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

func (fs *FileStore) HasAt( p []byte, off int64 ) (n int, err error) {
	len := int64(cap( p ))
	sz,_ := fs.FileSize()

	switch {
		case (off + len) < sz: return int(len), nil
		case off > sz: return 0, FileStoreError{"Offset beyond end of file"}
		case (off + len) > sz: return int(sz - off), nil
	}

	return 0, FileStoreError{"Shouldn't get here"}
}

func (fs *FileStore) FileSize() (int64,error) {
	stat,_ := fs.file.Stat()
	return stat.Size(),nil
}
