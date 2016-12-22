package lazyfs


type LocalFSSource struct {
  root string
}

func OpenLocalFSSource( root string ) (*LocalFSSource, error) {
  fs := LocalFSSource{ root: root }
  return &fs, nil
}

func (fs *LocalFSSource ) Open( path string ) (*FileSource, error) {
  return OpenLocalFileSource( fs.root, path )
}
