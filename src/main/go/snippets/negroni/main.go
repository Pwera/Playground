package main

import (
	"fmt"
	"github.com/Pwera/Playground/src/main/go/snippets/negroni/negronigzip"
	"net/http"

	"github.com/Pwera/Playground/src/main/go/snippets/negroni/custommiddleware"
	"github.com/Pwera/Playground/src/main/go/snippets/negroni/negronilogrus"
	"github.com/Pwera/Playground/src/main/go/snippets/negroni/negronicors"
	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	customMiddleware := custommiddleware.CustomMiddleware{}
	negroniLogrusWrapper := negronilogrus.NewNegroniLogrusWrapper()
	negroniGzipWrapper := negronigzip.NewNegroniGzipWrapper()
	negroniCorsWrapper := negronicors.NewNegroniCorsWrapper()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, h *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})
	negronigzip.NewNegroniGzipWrapper()
	nergoni := negroni.New()
	negroni.New()
	nergoni.UseHandler(mux)

	nergoni.Use(negroni.NewRecovery())
	nergoni.Use(negroni.NewLogger())
	nergoni.Use(customMiddleware)
	nergoni.Use(negroniGzipWrapper)
	nergoni.Use(negroniLogrusWrapper)
	nergoni.Use(negroniCorsWrapper)

	nergoni.Run(":3000")
}
