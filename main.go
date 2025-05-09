package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/gen/api/v1/apiv1connect"
	"project/service/api"
	"syscall"
	"time"

	"connectrpc.com/vanguard"
)

//go:generate buf generate
func main() {
	apiService := api.NewService()

	mux := http.NewServeMux()

	vanguardServices := []*vanguard.Service{
		vanguard.NewService(apiv1connect.NewApiServiceHandler(apiService)),
	}
	transcoder, err := vanguard.NewTranscoder(
		vanguardServices,
		vanguard.WithUnknownHandler(http.NotFoundHandler()),
	)
	if err != nil {
		log.Fatalf("Failed to create Vanguard transcoder: %v", err)
	}
	mux.Handle("/", transcoder)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// rest is server starting and grace stop code

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("HTTP server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server gracefully stopped")
}
