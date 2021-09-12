package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jgoralcz/cdbapi/src/lib/helpers"
	"github.com/jgoralcz/cdbapi/src/routes"
)

// Run initializes the http server with the echo framework.
func Run() {
	echo := routes.Routes()

	port := helpers.GetEnvOrDefault("PORT", "8443")
	env := helpers.GetEnvOrDefault("ENV", "LOCAL")
	addr := ":" + port

	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	go func() {
		log.Info("Starting ", env, " server on port ", port)
		echo.Logger.Fatal(echo.StartServer(server))
	}()

	waitForShutdown(server)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Fatal("Shutting down...")
}
