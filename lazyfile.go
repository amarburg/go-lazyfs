package lazyfs

import "io"

type FileSource interface {
	io.ReaderAt
	FileSize() (int64, error)
	Reader() io.Reader
	Path() string
}
