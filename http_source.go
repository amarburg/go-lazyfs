package lazyfs

import "github.com/valyala/fasthttp"
import "net/http"
import "fmt"
import "io"
import "strings"
import "strconv"
import "net/url"
import "bytes"
import "time"

import prom "github.com/prometheus/client_golang/prometheus"

//==== prometheus instrumentation ==

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	promHttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_requests_total",
			Help:        "Total number of HTTP requests.",
			ConstLabels: prom.Labels{"handler": "lazyhttp"},
		},
		[]string{"code", "root", "method"},
	)

	promHttpResponseSize = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:        "lazyhttp_content_size_bytes",
			Help:        "Total size of HTTP content requested.",
			ConstLabels: prom.Labels{"handler": "lazyhttp"},
		},
		[]string{"root"},
	)

	promHttpDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:        "http_request_duration_microseconds",
			Help:        "Duration of HTTP request.",
			ConstLabels: prom.Labels{"handler": "lazyhttp"},
		},
		[]string{"root"},
	)
)

func init() {
	prometheus.MustRegister(promHttpRequests)
	prometheus.MustRegister(promHttpResponseSize)
	prometheus.MustRegister(promHttpDuration)
	// prometheus.MustRegister(promCacheSize)
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

	client fasthttp.Client
}

func OpenHttpSource(url url.URL) (hsrc *HttpSource, err error) {
	h := HttpSource{url: url,
		client: fasthttp.Client{},
	}

	return &h, nil
}

func (fs *HttpSource) ReadAt(p []byte, off int64) (n int, err error) {

	startTime := time.Now()
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(fs.url.String())

	// Byte ranges are inclusive...
	//fmt.Printf("Reading %d-%d from %s\n", off, off+int64(cap(p))-1, fs.url.String())
	request.Header.SetByteRange(int(off), int(off)+cap(p)-1)

	response := fasthttp.AcquireResponse()
	err = fs.client.Do(request, response)

	if err != nil {
		return 0, fmt.Errorf("Error response from HTTP client")
	}

	promHttpRequests.With(prom.Labels{"code": fmt.Sprintf("%d", response.StatusCode),
		"root":   fs.url.String(),
		"method": "GET"}).Inc()

	if err != nil {
		// fs.Stats.Errors++
		return 0, fmt.Errorf("error from HTTP client: %s\n", err.Error())
	}

	// This is probably not idiomatic
	buffer := bytes.NewBuffer(nil)
	response.BodyWriteTo(buffer)
	responseLength := response.Header.ContentLength()

	copy(p, buffer.Bytes()[:responseLength])

	dt := time.Since(startTime)
	promHttpResponseSize.With(prom.Labels{
		"root": fs.url.String()}).Observe(float64(len(p)))
	promHttpDuration.With(prom.Labels{
		"root": fs.url.String()}).Observe(float64(dt.Nanoseconds()) / 1000.0)

	fasthttp.ReleaseResponse(response)
	fasthttp.ReleaseRequest(request)

	return responseLength, err
}

func (fs *HttpSource) FileSize() (int64, error) {

	//startTime := time.Now()
	request := fasthttp.AcquireRequest()
	request.SetRequestURI(fs.url.String())
	request.Header.SetByteRange(0, 0)

	response := fasthttp.AcquireResponse()
	err := fs.client.Do(request, response)

	if err != nil {
		return int64(-1), err
	} else if response == nil {
		return int64(-1), fmt.Errorf("Received no response")
	}

	//TODO: Check status

	content_range := response.Header.Peek("Content-Range")

	if content_range == nil {
		return int64(-1), fmt.Errorf("Response header didn't have Content-Range: %v", response.Header)

		// As a fallback, look for content-length
		// content_length := response.Header["Content-Length"]
		// if content_length == nil {
		// 	panic( fmt.Sprintf("Response header didn't have Content-Range or Content-Length: %v", response.Header ))
		// }
		//
		// l,err := strconv.Atoi( content_length[0] )
		// if err != nil {
		// 	panic( fmt.Sprintf("Couldn't extract content length from \"%s\": %s", content_length[0], err.Error()))
		// }
		// return int64( l ),nil
	}

	// Extract the Header
	splits := strings.Split(string(content_range), "/")
	if len(splits) != 2 {
		return int64(-1), fmt.Errorf("Couldn't parse the Content-Range header: ", content_range)
	}

	//fmt.Println( response.Header )
	l, err := strconv.Atoi(splits[1])
	if err != nil {
		return int64(-1), fmt.Errorf("Couldn't extract content length from \"%s\": %s", splits[1], err.Error())
	}

	fasthttp.ReleaseResponse(response)
	fasthttp.ReleaseRequest(request)

	return int64(l), nil
}

func (fs *HttpSource) Reader() io.Reader {
	request, _ := http.NewRequest("GET", fs.url.String(), nil)
	client := http.Client{}
	response, _ := client.Do(request)

	return response.Body
}

func (fs *HttpSource) Path() string {
	return fs.url.Host + fs.url.Path
}
