package lazyfs

import "testing"


func TestHttpSource(t *testing.T) {
  const TestUrl = "https://raw.githubusercontent.com/amarburg/lazyfs/master/test_files/foo.fs"


  fs,err := OpenHttpSource(TestUrl)

  if err != nil {
    t.Error("Couldn't create fs:", err)
  }

  for _,test := range test_pairs {

    buf := make([]byte, BufSize)
    n,err := fs.ReadAt( buf, test.offset )

    //TODO:  Error handling

    buf = buf[:n]

    if n != test.length {
      t.Error("Expected",test.length,"bytes, got",n)
    }

    if !CheckTestFile( buf, test.offset ) {
      t.Error("Reading",cap(buf),"bytes from HTTP source at offset",
              test.offset,"is incorrect(",n," bytes returned): ",
              err, buf )
    }

  }
}
