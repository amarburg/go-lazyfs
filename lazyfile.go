package lazyfs

import "io"

const Version = "v0.1.0"

type FileSource interface {
	io.ReaderAt
	FileSize() (int64, error)
	Reader() io.Reader
	Path() string
}
