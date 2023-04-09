package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
}

func NewServer(mux http.Handler, addr string) *Server {
	return &Server{
		srv: &http.Server{Handler: mux, Addr: addr},
	}
}


func (s *Server) Run(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.srv.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %v", err)
			return err
		}
		return nil
	})

	<-ctx.Done()
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("faild to shutdown: %v", err)
	}

	return eg.Wait()
}