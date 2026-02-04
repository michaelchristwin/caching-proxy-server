package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"caching-proxy/internal/routes"
)

func run(ctx context.Context, w io.Writer) error {
	serverPort := flag.Int("port", 0, "port number")
	origin := flag.String("origin", "", "origin")
	flag.Parse()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	r := routes.SetupRoutes(*origin)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(*serverPort),
		Handler: r,
	}

	errCh := make(chan error, 1)

	go func() {
		fmt.Fprintf(w, "server started | addr=%s | pid=%d\n", srv.Addr, os.Getpid())
		errCh <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		// shutdown signal received
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		return nil

	case err := <-errCh:

		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
}

func main() {

	ctx := context.Background()
	if err := run(ctx, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}
