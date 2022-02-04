package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var stopSigs = []os.Signal{syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "hello~")
	})

	srv := http.Server{Addr: ":8080"}

	connClosed := make(chan struct{})
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, stopSigs...)

		sig := <-sigChan
		log.Println("received: ", sig)

		_ = srv.Shutdown(context.Background())

		close(connClosed)
	}()

	fmt.Println("Listening... ", 8080)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-connClosed
}
