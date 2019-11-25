package negronigzip

import (
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/urfave/negroni"
)

type NegroniGzipWrapper struct{
	negroni.Handler
}

func NewNegroniGzipWrapper() NegroniGzipWrapper {
	return NegroniGzipWrapper{gzip.Gzip(gzip.DefaultCompression)}
}



