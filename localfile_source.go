package lazyfs

import "os"
import "io"
import "path"

type LocalFileSource struct {
	root, path string
	file       *os.File
}

// OpenLocalFile takes a local path and produces a
// pointer to a LocalFileSource, which implements the
// FileSource interface.
func OpenLocalFile(fpath string) (fsrc *LocalFileSource, err error) {
	root, file := path.Split(fpath)
	return OpenLocalFileSource(root, file)
}

// OpenLocalFileSource takes a two-term path (a "root" and an relative path)
// and produces a pointer to a LocalFileSource, which implements the
// FileSource interface.
func OpenLocalFileSource(root string, ff string) (fsrc *LocalFileSource, err error) {
	f, err := os.Open(path.Join(root, ff))

	// TODO:  Check for non-existent files

	return &LocalFileSource{root: root, path: ff, file: f}, err
}

// ReadAt implements the ReaderAt interface for LocalFileSource.
// Attempts to read len(p) bytes from offset off within the LocalFile.
// Returns the number of bytes read.
// ( it's really a wrapper around fs.file.ReadAt() )
func (fs *LocalFileSource) ReadAt(p []byte, off int64) (n int, err error) {
	return fs.file.ReadAt(p, off)
}

// FileSize returns the size of file underlying the LocalFileSource
func (fs *LocalFileSource) FileSize() (int64, error) {
	stat, _ := fs.file.Stat()
	return stat.Size(), nil
}

// Generate a Reader for the file pointed to by the LocalFileSource
func (fs *LocalFileSource) Reader() io.Reader {
	f, _ := os.Open(fs.root + fs.path)
	return f
}

// Returns the path to the LocalFileSource
func (fs *LocalFileSource) Path() string {
	return path.Join(fs.root, fs.path)
}
