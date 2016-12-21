package lazyfs


type HttpFSSource struct {
  url_root string
}

func OpenHttpFSSource( url_root string ) (*HttpFSSource, error) {
  fs := HttpFSSource{ url_root: url_root }
  return &fs, nil
}

func (fs *HttpFSSource ) Open( path string ) (*HttpSource, error) {
  return OpenHttpSource( fs.url_root + path )
}
