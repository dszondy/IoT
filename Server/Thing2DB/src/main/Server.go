package main

import (
	"log"
	"time"
	"os"
	"os/signal"
	"syscall"
	"net/http"
	"context"
	"conn"
	"runtime/debug"
	"fmt"
)

func main() {
	runServer()
}

func runServer() bool {
	<-time.NewTimer(10* time.Second).C
	go conn.InitConnectionLayer()
	var signalStop = make(chan os.Signal)
	signal.Notify(signalStop, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	var errorStop = make(chan bool)
	<-time.NewTimer(500*time.Millisecond).C
	router := NewRouter()
	server := &http.Server{
		Addr: ":9848", Handler: router,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go listenAndServe(server, errorStop)

	select {
	case <-errorStop:
		log.Printf("Received error signal, exit.")
		return false

	case signal := <-signalStop:
		log.Printf("Received signal '%v', starting shutdown...", signal)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error: %v", err)
			return false
		}

		return true
	}
}

func listenAndServe(server *http.Server, errorStop chan bool) {
	log.Printf("Listening on http://0.0.0.0%v", server.Addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Error: %v", err)
		errorStop <- true
	}

	log.Printf("Stopped listening")
}

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func errorHandler(handler http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Error catched by route error handler: %s\nError: %v\n%s", name, r, debug.Stack())
				respondWithError(w, http.StatusInternalServerError, fmt.Sprint(r))
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
