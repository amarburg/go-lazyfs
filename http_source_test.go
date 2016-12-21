package lazyfs

import "testing"
import "fmt"



func TestHttpSource(t *testing.T) {
  url := "https://raw.githubusercontent.com/amarburg/lazyfs/master/test_files/foo.fs"

  fs,err := OpenHttpSource(url)

  if err != nil {
    t.Error("Couldn't create fs:", err)
  }

  buf := make([]byte,10)
  n,err := fs.ReadAt( buf, 0 )

fmt.Println(buf,n,err)

  if !CheckTestFile( buf, 0 ) {
    t.Error("Data from HTTP source is incorrect(",n,"): ", buf)
  }
}
