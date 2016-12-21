package lazyfs


//import ( "io" )


type Path string

type Storage interface {
  Open( path string ) FileStorage
}

type Source interface {
  Open( path string ) FileStorage
}


type LazyFS struct {
  storage Storage
  source Source
}
