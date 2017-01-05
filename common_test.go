package lazyfs

import "net/url"
import "reflect"

var BufSize = 10

var test_pairs = []struct {
  offset int64
  length int
}{
  {0,  BufSize},
  {10, BufSize},
  {20, 6},
}

var AlphabetSize int64 = 26

var LocalFilesRoot = "test_files/a/"
var TestUrlRoot = "https://raw.githubusercontent.com/amarburg/lazyfs/master/test_files/a/"
var AlphabetPath = "x/alphabet.fs"
var AlphabetUrl,_ = url.Parse( TestUrlRoot + AlphabetPath )

var LocalStoreRoot = "test_files/localfilestore/"
var SparseStoreRoot = "test_files/localsparsestore/"

var SparseHttpStoreRoot = "test_files/httpsparse/"

var BadPath = "test_files/a/y/foo.fs"


var OOIRawDataRootURL = "https://rawdata.oceanobservatories.org/files/"


// var LocalAlphabetPath = LocalFilesRoot + AlphabetPath
// var LocalBadPath      = LocalFilesRoot + BadPath
// var AlphabetUrl = TestUrlRoot + AlphabetPath


func CheckTestFile( buf []byte, off int64 ) bool {
  // I'm sure there's a beautiful idiomatic Go way to do this
  test_string := [26]byte{ 65, 66, 67, 68, 69, 70, 71, 72, 73, 74,
    75, 76, 77, 78, 79, 80, 81, 82, 83, 84,
    85, 86, 87, 88, 89, 90 }

  l := int(off)+len(buf)
  return reflect.DeepEqual(buf, test_string[int(off):l] )
}
