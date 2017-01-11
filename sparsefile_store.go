package lazyfs

import "os"
import "fmt"
import "path/filepath"
import "io"

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
	source FileSource
	has map[int64]bool
}

func OpenSparseFileStore( source FileSource, root string ) (*SparseFileStore,error) {
	sz,_ := source.FileSize()
	f,err := InitializeSparsefile( root + source.Path(), sz )
	if err != nil { return nil,err }

	fs := SparseFileStore{ file: f, source: source, has: make( map[int64]bool, sz ) }

	return &fs, err
}

func InitializeSparsefile( sparsefile string, sz int64 ) (*os.File, error) {
	// Fill file with with null

	fileinfo,err := os.Stat( sparsefile )
	if err != nil || fileinfo.Size() != sz {

		fmt.Println("Creating sparsefile size", sparsefile, sz)
		os.MkdirAll( filepath.Dir(sparsefile), 0755 )
		dest,err := os.Create( sparsefile )

		if err != nil {
			panic(fmt.Sprintf("Couldn't create sparsefile %s", sparsefile) )
		}

		zero := &ZeroReader{ size: sz }
		io.Copy( dest, zero )
	}

	file,err := os.OpenFile( sparsefile, os.O_RDWR, 0644 )

	return file,err
}



type ZeroReader struct {
  size int64
}

func (w *ZeroReader) Read( p []byte) (n int, err error) {
if int64(cap(p)) > w.size {
	n = int(w.size)
} else {
	n = cap(p)
}

for i := 0; i < n; i++ { p[i] = 0 }

w.size -= int64(n)

if w.size == 0 {
	err = io.EOF
	} else {
		err = nil
	}

	return n,err
}


//=====================================
func (fs *SparseFileStore) ReadAt( p []byte, off int64) (n int, err error) {
	// Check validity
  if _,err := fs.HasAt(p,off); err == nil  {
    //fmt.Println("Retrieving from store")
    return fs.file.ReadAt( p, off )
  }

  //fmt.Println( "Need to update store")
  n,_ = fs.source.ReadAt(p,off)
  fs.WriteAt(p[:n], off)

  return n, nil

  // }
	//
	// n,err =  fs.HasAt( p, off )
	// if err != nil {
	// 	return 0, SparseFileStoreError{"ReadAt: Don't have all of the requested bytes"}
	// }
	//
	// return fs.file.ReadAt( p, off )
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
		//fmt.Println(off+int64(i),"=",fs.has[off+int64(i)])
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

func (fs *SparseFileStore) Reader() (io.Reader) {
	return fs.source.Reader()
}

func (fs *SparseFileStore) Path() string {
	return fs.source.Path()
}
