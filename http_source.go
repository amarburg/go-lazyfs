package lazyfs

import "net/http"
import "fmt"
import "io"


type HttpSource struct {
  url string
  path string
  client http.Client
}

func OpenHttpSource( url_root string, path string ) (hsrc *HttpSource, err error ) {
  h := HttpSource{ url: url_root + path, path: path }
  return &h, nil
}

func (fs *HttpSource) ReadAt( p []byte, off int64 ) (n int, err error) {
  request, err := http.NewRequest("GET", fs.url, nil)

  range_str := fmt.Sprintf("bytes=%d-%d", off, off+int64(cap(p)))
  request.Header = map[string][]string{
                    "Range": { range_str },
                  }

  response, err := fs.client.Do( request )

  //TODO: Check status

  defer response.Body.Close()
  n, err = response.Body.Read(p)

  return n, err
}


func (fs *HttpSource) Reader() io.Reader {
  request, _ := http.NewRequest("GET", fs.url, nil)
  response, _ := fs.client.Do( request )

  return response.Body
}


func (fs *HttpSource) Path() string {
  return fs.path
}
