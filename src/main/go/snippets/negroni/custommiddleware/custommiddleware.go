package custommiddleware

import (
	"fmt"
	"net/http"
)

type CustomMiddleware struct{}

func (c CustomMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("CustomMiddleware")
	next(rw, r)
}
