package main

import (
	"fmt"
	"github.com/pwera/negroni/negronigzip"
	"net/http"
	"os"
	"runtime/pprof"

	"github.com/pwera/negroni/custommiddleware"
	"github.com/pwera/negroni/negronicors"
	"github.com/pwera/negroni/negronilogrus"
	"github.com/urfave/negroni"
)

func main() {
	f, err := os.Create("./cpu.prof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	mux := http.NewServeMux()
	customMiddleware := custommiddleware.CustomMiddleware{}
	negroniLogrusWrapper := negronilogrus.NewNegroniLogrusWrapper()
	negroniGzipWrapper := negronigzip.NewNegroniGzipWrapper()
	negroniCorsWrapper := negronicors.NewNegroniCorsWrapper()
	middleware2 := custommiddleware.NewSecondMiddleware()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, h *http.Request) {
		fmt.Fprintf(w, "Welcome")
	})
	//negronigzip.NewNegroniGzipWrapper()
	nergoni := negroni.New()
	//negroni.New()
	nergoni.UseHandler(mux)

	nergoni.Use(negroni.NewRecovery())
	nergoni.Use(negroni.NewLogger())
	nergoni.Use(customMiddleware)
	nergoni.Use(negroniGzipWrapper)
	nergoni.Use(negroniLogrusWrapper)
	nergoni.Use(negroniCorsWrapper)
	nergoni.Use(middleware2)

	nergoni.Run(":3000")
}
