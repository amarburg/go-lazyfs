package lazyfs

import "fmt"
import "net/http"
import "net/url"
import "errors"
import "golang.org/x/net/html"

type HttpFSSource struct {
  url_root string
  store FSStorage
}

func OpenHttpFSSource( url_root string ) (*HttpFSSource, error) {
  fs := HttpFSSource{ url_root: url_root }
  return &fs, nil
}

func (fs *HttpFSSource) SetBackingStore( store FSStorage ) {
  fs.store = store
}

func (fs *HttpFSSource ) Open( path string ) (*HttpSource, error) {
  url,_ := url.Parse(fs.url_root + path)
  src,err := OpenHttpSource( *url )

  // if fs.store != nil {
  //     st,err := fs.store.Store( src )
  //
  //     if err != nil {
  //       panic(fmt.Sprintf("Couldn't create store for file %s",path))
  //     }
  //
  //     //src.SetBackingStore( st )
  // }

  return src,err
}

type DirListing struct {
  Path string
  Children []string
}

func (fs *HttpFSSource ) ReadHttpDir( path string ) (DirListing,error){
  client := http.Client{}

fmt.Printf( "Querying: %s\n", fs.url_root + path )

  response, err := client.Get( fs.url_root + path )

  listing := DirListing{ Path: path }

  //fmt.Println( response, err )

  if( response.StatusCode != 200 ) {
    return listing, errors.New(fmt.Sprintf("Got HTTP response %d", response.StatusCode))
  }

  defer response.Body.Close()
  d := html.NewTokenizer( response.Body )

  //fmt.Println(d)

  for {
  	tt := d.Next()

    if tt == html.ErrorToken { break; }

    // fmt.Println(tt)

    // Big ugly brutal
    switch tt {
  	case html.StartTagToken:
      token := d.Token()

      for _,attr := range token.Attr {
        if attr.Key == "href" {
          //fmt.Println(attr.Key)
          val := attr.Val

          tt = d.Next()
          if( tt == html.TextToken ) {
            next := d.Token()
            text := next.Data
            //fmt.Println(text)

            if val == text  {
              listing.Children = append(listing.Children, text)
            }
          }
        }
      }

    }
  }

  return listing, err
}
