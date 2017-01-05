package lazyfs

import "net/http"
import "fmt"
import "io"
import "strings"
import "strconv"
import "net/url"

type HttpSource struct {
  url string
  path string
  client http.Client
  store FileStorage
}

func OpenHttpSourceUrl( url url.URL ) (hsrc *HttpSource, err error ) {
  h := HttpSource{ url: url.String(), path: url.Path }
  return &h, nil
}

func OpenHttpSource( url_root string, path string ) (hsrc *HttpSource, err error ) {
  h := HttpSource{ url: url_root + path, path: path }
  return &h, nil
}

func (fs *HttpSource) SetBackingStore( store FileStorage ) {
	fs.store = store
}

func (fs *HttpSource) ReadAt( p []byte, off int64 ) (n int, err error) {
  if fs.store != nil {
    if _,err := fs.store.HasAt(p,off); err == nil  {
      //fmt.Println("HttpSource: Retrieving",len(p),"from store")
      return fs.store.ReadAt( p, off )
    } else {
      //fmt.Println( "HttpSource: Need to update store, querying HTTP")
      n,_ := fs.ReadHttpAt(p,off)
      fs.store.WriteAt(p[:n], off)

      return n, nil
    }
  } else {
    return fs.ReadHttpAt(p,off)
  }
}

func (fs *HttpSource) ReadHttpAt( p []byte, off int64 ) (n int, err error) {

  request, err := http.NewRequest("GET", fs.url, nil)

  range_str := fmt.Sprintf("bytes=%d-%d", off, off+int64(cap(p)))
  request.Header = map[string][]string{
                    "Range": { range_str },
                  }

  response, err := fs.client.Do( request )

//fmt.Println(response)
//fmt.Println(p)
  //TODO: Check status

  defer response.Body.Close()

  idx := 0
  for {
    n, err = response.Body.Read( p[idx:] )
    idx += n
    if idx >= len(p) || err != nil { break }
  }

  return idx, err
}


func (fs *HttpSource) FileSize() (int64,error) {
  // Don't know if this always works
  request,_ := http.NewRequest("GET", fs.url, nil)
  request.Header = map[string][]string{
                    "Range": { "bytes=0-0" },
                  }
  response, _ := fs.client.Do( request )
  defer response.Body.Close()

  //TODO: Check status

  content_range := response.Header["Content-Range"]
  if content_range == nil {
    panic( fmt.Sprintf("Response header didn't have Content-Range: %v", response.Header ))
  }

  // Extract the Header
  splits := strings.Split( content_range[0], "/")
  if len(splits) != 2 {
    panic( fmt.Sprintf("Couldn't parse the Content-Range header: ", content_range ) )
  }

  //fmt.Println( response.Header )
  l,err := strconv.Atoi(splits[1])
  if err != nil {
    panic( fmt.Sprintf("Couldn't extract content length from \"%s\": %s", splits[1], err.Error()))
  }
  fmt.Println("Got content length", l)
  return int64(l),nil
}

func (fs *HttpSource) Reader() io.Reader {
  request, _ := http.NewRequest("GET", fs.url, nil)
  response, _ := fs.client.Do( request )

  return response.Body
}


func (fs *HttpSource) Path() string {
  return fs.path
}
