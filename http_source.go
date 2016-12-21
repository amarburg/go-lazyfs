package lazyfs

import "net/http"
import "fmt"


type HttpSource struct {
  url string
  client http.Client
}

func OpenHttpSource( url string ) (hsrc HttpSource, err error ) {
  return HttpSource{ url: url }, nil
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
