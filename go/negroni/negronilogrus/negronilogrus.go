package negronilogrus

import (
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
)

type NegroniLogrusWrapper struct {
	negroni.Handler
}


func NewNegroniLogrusWrapper() NegroniLogrusWrapper{
	return NegroniLogrusWrapper{negronilogrus.NewMiddleware()}
}