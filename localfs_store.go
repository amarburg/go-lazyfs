package lazyfs


type LocalFSStore struct {
  root  string
}

func OpenLocalFSStore( root string ) (*LocalFSStore, error) {
  fs := LocalFSStore{ root: root }
  return &fs, nil
}


func (fs *LocalFSStore ) Open( path string ) (*FileStore, error) {
  return OpenFileStore( fs.root + path ), nil
}
