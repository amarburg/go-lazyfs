package lazyfs

import "os"
import "io"

import "fmt"
import "path/filepath"

type LocalFileStoreError struct {
	Err string
}

func (e LocalFileStoreError) Error() string {
	return e.Err
}

type LocalFileStore struct {
	source FileSource
	root   string
	file   *os.File
}

// OpenLocalFileStore makes a LocalFileStore which wraps around a FileSource
//  root specifies the local location for local file store cache.
// Returns a pointer to the LocalFileStore
func OpenLocalFileStore(source FileSource, root string) (*LocalFileStore, error) {
	fs := LocalFileStore{file: nil, root: root, source: source}
	return &fs, nil
}

// load copies the entirety of the file from the FileSource to the local
// location.
func (fs *LocalFileStore) load() error {
	if fs.file == nil {

		path := fs.root + fs.source.Path()

		os.MkdirAll(filepath.Dir(path), 0755)

		f, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		reader := fs.source.Reader()
		io.Copy(f, reader)

		f.Close()

		fs.file, _ = os.Open(path)

	}
	return nil
}

// ReadAt is first lazy-loads the wrapped FileSource to the local disk using
// load(), then reads from the local disk
func (fs *LocalFileStore) ReadAt(p []byte, off int64) (n int, err error) {
	if err := fs.load(); err != nil {
		return 0, err
	}

	return fs.file.ReadAt(p, off)
}

// LocalFileStore implements HasAt by first lazy-loading the file using
// load().  Once the file is available locally, the question is not
// whether the cache has the bytes, but whether the bytes exist in the
// file at all ... or are e.g., off the end of the file.  Returns the
// number of bytes available, which may be <= len(p) if successful.
// Or 0 if the bytes are not available.
func (fs *LocalFileStore) HasAt(p []byte, off int64) (n int, err error) {
	if err := fs.load(); err != nil {
		return 0, err
	}

	len := int64(cap(p))
	sz, _ := fs.FileSize()

	switch {
	case (off + len) < sz:
		return int(len), nil
	case off > sz:
		return 0, LocalFileStoreError{"Offset beyond end of file"}
	case (off + len) > sz:
		return int(sz - off), nil
	}

	return 0, LocalFileStoreError{"Shouldn't get here"}
}

// Returns the size of the file underlying the LocalFileStore
func (fs *LocalFileStore) FileSize() (int64, error) {
	stat, _ := fs.file.Stat()
	return stat.Size(), nil
}

// Returns a Reader to the LocalFileStore
func (fs *LocalFileStore) Reader() io.Reader {
	return fs.source.Reader()
}

// Returns the path to the LocalFileStore
func (fs *LocalFileStore) Path() string {
	return fs.source.Path()
}
