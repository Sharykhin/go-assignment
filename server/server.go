package server

import (
	"fmt"
	"github.com/Sharykhin/go-assignment/logger"
	"log"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	serverWriteTimeout = 30 * time.Second
	serverReadTimeout = 30 * time.Second
	serverIdleTimeout = 15 * time.Second
	serverShutDownTimeout = 10* time.Second

)

// ListenAndServe starts a new http server with graceful shutdown
func ListenAndServe(port string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	s := &http.Server{
		Handler:      router(),
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: serverWriteTimeout,
		ReadTimeout:  serverReadTimeout,
		IdleTimeout:  serverIdleTimeout,
	}

	go func() {
		err := s.ListenAndServe()
		if err != http.ErrServerClosed {
			logger.Log.Errorf("[server][ListenAndServe] failed to start http server: %v", err)
		}
	}()
	logger.Log.Infof("[server][ListenAndServe] Server is up and running on %s port", port)
	sig := <-interrupt
	log.Printf("[server][ListenAndServe] got interrupt signal %s, going to gracefully shutdown the server", sig)
	ctx, cancel := context.WithTimeout(context.Background(), serverShutDownTimeout)
	defer cancel()
	err := s.Shutdown(ctx)
	if err != nil {
		logger.Log.Errorf("[server][ListenAndServe] failed to gracefully shutdown the server; %v", err)
	}

	logger.Log.Infof("[server][ListenAndServe] the server successfully stopped")
}