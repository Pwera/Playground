package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pwera/Playground/src/main/go/snippets/di/handlers"
	"github.com/Pwera/Playground/src/main/go/snippets/di/logging"
	"github.com/Pwera/Playground/src/main/go/snippets/di/middlewares"
	"github.com/Pwera/Playground/src/main/go/snippets/di/services"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

func main() {

	defer logging.Logger.Sync()

	builder, err := di.NewBuilder()
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	err = builder.Add(services.Services...)
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	app := builder.Build()
	defer app.Delete()

	r := mux.NewRouter()

	m := func(h http.HandlerFunc) http.HandlerFunc {
		return middlewares.PanicRecoveryMiddleware(
			di.HTTPMiddleware(h, app, func(msg string) {
				logging.Logger.Error(msg)
			}),
			logging.Logger,
		)
	}
	r.HandleFunc("/cars", m(handlers.GetCarListHandler)).Methods("GET")
	r.HandleFunc("/cars", m(handlers.PostCarHandler)).Methods("POST")
	r.HandleFunc("/cars/{carId}", m(handlers.GetCarHandler)).Methods("GET")
	r.HandleFunc("/cars/{carId}", m(handlers.PutCarHandler)).Methods("PUT")
	r.HandleFunc("/cars/{carId}", m(handlers.DeleteCarHandler)).Methods("DELETE")

	port := os.Getenv("SERVER_PORT")
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logging.Logger.Info("Listening on port " + port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Logger.Error(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	logging.Logger.Info("Stopping the http server")
	if err := srv.Shutdown(ctx); err != nil {
		logging.Logger.Error(err.Error())
	}
}
