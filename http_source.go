package lazyfs

import "net/http"
import "fmt"
import "io"
import "strings"
import "strconv"
import "net/url"

import prom "github.com/prometheus/client_golang/prometheus"

//==== Prometheus instrumentation ==

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PromHttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
      ConstLabels: prom.Labels{ "handler": "lazyhttp" },
		},
    []string{"code","root","method"},
	)

  PromHttpResponseSize = prometheus.NewSummaryVec(
    prometheus.SummaryOpts{
      Name: "lazyhttp_content_size_bytes",
			Help: "Total size of HTTP content requested.",
      ConstLabels: prom.Labels{ "handler": "lazyhttp" },
    },
    []string{"root",},
  )

// 	PromCacheMisses = prometheus.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Name: "cache_misses_total",
// 			Help: "Number of cache misses.",
// 		},
//     []string{"store"},
// 	)
//
// 	PromCacheSize = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 		Name: "cache_size",
// 		Help: "Number of elements in cache",
// 	},
//   []string{"store"},
// )
)

func init() {
	prometheus.MustRegister(PromHttpRequests)
  prometheus.MustRegister(PromHttpResponseSize)
	// prometheus.MustRegister(PromCacheSize)
}

//====


// type HttpStatistics struct {
//   Transactions int
//   Errors int
//   ContentBytesRead int
//   TotalBytesWritten, TotalBytesRead int
// }

type HttpSource struct {
  url url.URL
}

func OpenHttpSource( url url.URL ) (hsrc *HttpSource, err error ) {
  h := HttpSource{ url: url }
  return &h, nil
}

// func (fs *HttpSource) SetBackingStore( store FileStorage ) {
// 	fs.store = store
// }

func (fs *HttpSource) ReadAt( p []byte, off int64 ) (n int, err error) {

  request, err := http.NewRequest("GET", fs.url.String(), nil)

  range_str := fmt.Sprintf("bytes=%d-%d", off, off+int64(cap(p)))
  request.Header = map[string][]string{
                    "Range": { range_str },
                  }

  client := http.Client{}
  response, err := client.Do( request )

PromHttpRequests.With( prom.Labels{"code": fmt.Sprintf("%d",response.StatusCode),
              "root": fs.url.String(),
              "method": "GET" } ).Inc()

  // How to get size of tx/rx without serializing twice?
  //fs.Stats.Transactions++

  if err != nil {
    // fs.Stats.Errors++
    panic( fmt.Sprintf("error from HTTP client: %s\n", err.Error() ))
  } else if response == nil {
    panic( "Nil response from HTTP client")
  }

// fmt.Println(response.Header)
//   cl := response.Header["Content-Length"]
//   if cl != nil {
//     b,_ := strconv.Atoi(response.Header["Content-Length"][0])
//     //fs.Stats.ContentBytesRead += b
//   }

  defer response.Body.Close()

  idx := 0
  for {
    n, err = response.Body.Read( p[idx:] )
    idx += n
    if idx >= len(p) || err != nil { break }
  }

  PromHttpResponseSize.With( prom.Labels{
                "root": fs.url.String() } ).Observe( float64(len(p)) )

  // fs.Stats.ContentBytesRead += len(p)

  return idx, err
}


func (fs *HttpSource) FileSize() (int64,error) {
  // Don't know if this always works
  request,_ := http.NewRequest("GET", fs.url.String(), nil)
  request.Header = map[string][]string{
                    "Range": { "bytes=0-0" },
                  }

  client := http.Client{}
  response, _ := client.Do( request )
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
  //fmt.Println("Got content length", l)
  return int64(l),nil
}

func (fs *HttpSource) Reader() io.Reader {
  request, _ := http.NewRequest("GET", fs.url.String(), nil)
  client := http.Client{}
  response, _ := client.Do( request )

  return response.Body
}


func (fs *HttpSource) Path() string {
  return fs.url.Host + fs.url.Path
}
