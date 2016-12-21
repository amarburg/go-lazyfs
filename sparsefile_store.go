package lazyfs

import "os"


//=====================================
type SparseFileStoreError struct {
	Err string
}

func (e SparseFileStoreError) Error() string {
	return e.Err
}


//=====================================
type SparseFileStore struct {
	file *os.File
}

func OpenSparseFileStore( name string ) (*SparseFileStore,error) {
	f,err := os.Open(name)
	fs := SparseFileStore{ file: f }

	return &fs, err
}

//=====================================
func (fs *SparseFileStore) ReadAt( p []byte, off int64) (n int, err error) {
	return fs.file.ReadAt( p, off )
}

func (fs *SparseFileStore) WriteAt(p []byte, off int64) (n int, err error) {
	return 0,nil
}

func (fs *SparseFileStore) HasAt( p []byte, off int64 ) (n int, err error) {
	len := int64(cap( p ))
	sz,_ := fs.FileSize()

	switch {
		case (off + len) < sz: return int(len), nil
		case off > sz: return 0, SparseFileStoreError{"Offset beyond end of file"}
		case (off + len) > sz: return int(sz - off), nil
	}

	return 0, SparseFileStoreError{"Shouldn't get here"}
}

func (fs *SparseFileStore) FileSize() (int64,error) {
	stat,_ := fs.file.Stat()
	return stat.Size(),nil
}
