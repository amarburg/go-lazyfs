package lazyfs


//import ( "io" )


type Path string

type Storage interface {
  Open( path string ) (*FileStorage, error)
}

type Source interface {
  Open( path string ) (*FileSourceage, error)
}

type LazyFS struct {
  storage Storage
  source Source
}
