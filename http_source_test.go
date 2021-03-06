package lazyfs

import "net/url"
import "testing"

import "github.com/amarburg/go-lazyfs-testfiles"
import "github.com/amarburg/go-lazyfs-testfiles/http_server"

func TestHttpSource(t *testing.T) {
	srv := lazyfs_testfiles_http_server.HttpServer()

	AlphabetURL, _ := url.Parse(srv.Url + lazyfs_testfiles.AlphabetFile)
	fs, err := OpenHttpSource(*AlphabetURL)

	if err != nil {
		t.Error("Couldn't create fs:", err)
	}

	for _, test := range test_pairs {

		buf := make([]byte, BufSize)
		n, err := fs.ReadAt(buf, test.offset)

		//TODO:  Error handling

		buf = buf[:n]

		if n != test.length {
			t.Error("Expected", test.length, "bytes, got", n)
		}

		if !CheckTestFile(buf, test.offset) {
			t.Error("Reading", cap(buf), "bytes from HTTP source at offset",
				test.offset, "is incorrect(", n, " bytes returned): ",
				err, buf)
		}

	}

	srv.Stop()
}
