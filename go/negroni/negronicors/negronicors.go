package negronicors

import (
	"github.com/rs/cors"
)

type NegroniCorsWrapper struct {
	*cors.Cors
}

func NewNegroniCorsWrapper() NegroniCorsWrapper {
	return NegroniCorsWrapper{cors.New(cors.Options{
		AllowedOrigins: []string{"http://foo.com"},
	})}
}
