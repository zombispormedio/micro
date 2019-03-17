package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

func waitForShutdown(onShutdown func()) chan struct{} {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-sigint
		onShutdown()
		close(idleConnsClosed)
	}()
	return idleConnsClosed
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	directory := os.Getenv("STATIC_DIR")
	if port == "" {
		port = "public"
	}

	fs := &fasthttp.FS{
		Root:               directory,
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: false,
		Compress:           true,
	}

	handler := fs.NewRequestHandler()

	server := &fasthttp.Server{
		Handler:              handler,
		Name:                 "Micro site server",
		ReadTimeout:          5 * time.Second,
		WriteTimeout:         10 * time.Second,
		MaxKeepaliveDuration: 5 * time.Second,
	}

	waitingForShutdown := waitForShutdown(func() {
		if err := server.Shutdown(); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
			os.Exit(0)
		}
	})

	log.Printf("Starting Server at %s", port)
	if err := server.ListenAndServe(":" + port); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}

	<-waitingForShutdown
}
