package lazyfs

import "os"
import "fmt"

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
	has map[int64]bool
}

func OpenSparseFileStore( name string ) (*SparseFileStore,error) {
	f,err := os.OpenFile(name, os.O_RDWR, 0644 )
	fs := SparseFileStore{ file: f }

	sz,err := fs.FileSize()
	fs.has = make( map[int64]bool, sz )

	return &fs, err
}

//=====================================
func (fs *SparseFileStore) ReadAt( p []byte, off int64) (n int, err error) {
	// Check validity
	n,err =  fs.HasAt( p, off )
	if err != nil {
		return 0, SparseFileStoreError{"ReadAt: Don't have all of the requested bytes"}
	}

	return fs.file.ReadAt( p, off )
}

func (fs *SparseFileStore) WriteAt(p []byte, off int64) (n int, err error) {
	n,err =  fs.file.WriteAt( p, off )

	for i:=0; i < n; i++ {
		fs.has[off + int64(i)] = true;
	}

	return n, err
}

func (fs *SparseFileStore) HasAt( p []byte, off int64 ) (n int, err error) {
	sz,err := fs.FileSize();
	if off+int64(cap(p)) > sz {
		n = int(sz - off)
	} else {
		n = cap(p)
	}

	for i:= 0; i < n; i++ {
		fmt.Println(off+int64(i),"=",fs.has[off+int64(i)])
		if fs.has[off+int64(i)] == false {
			return 0, SparseFileStoreError{"HasAt: Don't have all of the requested bytes"}
		}
	}

	return n, nil
}

func (fs *SparseFileStore) FileSize() (int64,error) {
	stat,_ := fs.file.Stat()
	return stat.Size(),nil
}
