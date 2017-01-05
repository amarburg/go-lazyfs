package lazyfs


//import ( "io" )


type Path string


type FSSource interface {
  Open( path string ) (FileSource, error)
}

type FSStorage interface {
  Store( source FileSource ) (FileStorage, error)
}


// type LazyFS struct {
//   storage FSStorage
//   source FSSource
// }


// func Open( storage FSStorage, source FSSource, path string ) {
//
// }
