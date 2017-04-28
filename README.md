# LazyFS


[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d81c2280fe8441d58c566d5c99cc30c8)](https://www.codacy.com/app/amarburg/go-lazyfs?utm_source=github.com&utm_medium=referral&utm_content=amarburg/go-lazyfs&utm_campaign=badger)
[![GoDoc](https://godoc.org/github.com/amarburg/go-lazyfs?status.svg)](https://godoc.org/github.com/amarburg/go-lazyfs)
[![wercker status](https://app.wercker.com/status/385d92fd589e317c6cf9f538013efc54/s/master "wercker status")](https://app.wercker.com/project/byKey/385d92fd589e317c6cf9f538013efc54)

Provides transparent byte-level caching of files pulled from a remote source (currently either another filesystem or an HTTP server that supports `Content-Range`).   

Tried to make a modular system where different Sources (e.g., HTTP, local file) and Stores (e.g., a local flat file, a database) can be paired up.   The resulting Source implements the Go [`ReaderAt`](https://golang.org/pkg/io/#ReaderAt) interface.

Currently has the these sources:

* __Local file__
* __HTTP__

And these stores:

* __Local files__ (makes a copy of the remote file to a local path)
* __Sparse file__ (creates an empty file of the same length as the source, fills in bytes as they are pulled from the source).   I think this will be very space efficient on compressed filesystems.   Not sure about the performance consequences.   Right now this store isn't persistent (it doesn't store the map of which bytes have been cached between executions, that's on the TODO list).

In the long-run I think one or more database-backed Stores will be the preferred solution for active deployment.

My first Golang library, so excuse any non-idiomatic patterns or other tomfoolery.

The repo [go-lazyfs-benchmarking](https://github.com/amarburg/go-lazyfs-benchmarking) contains a number of benchmarking tests which are computationally (or network bandwidth) more expensive (not appropriate as frequently-run testcases).

## TODO

- [ ] Write [GoDoc](https://blog.golang.org/godoc-documenting-go-code)
- [ ] Verify it's published on [GoDoc.org](https://godoc.org/)
- [ ] Complete test cases

## License

This library released is under the [MIT License](http://opensource.org/licenses/MIT)
