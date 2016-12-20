package lazyfs


type LazyFSError struct {
	Err string
}

func (e LazyFSError) Error() string {
	return e.Err
}
