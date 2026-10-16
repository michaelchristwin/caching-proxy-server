package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"

	"caching-proxy/internal/handler"
	internal_middleware "caching-proxy/internal/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type cachedResponse struct {
	status int
	header http.Header
	body   []byte
}

var cache = struct {
	sync.RWMutex
	data map[string]cachedResponse
}{
	data: make(map[string]cachedResponse),
}

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				a.Value = slog.StringValue(a.Value.Time().UTC().Format(time.RFC3339))
			}
			return a
		},
	}))
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(internal_middleware.Logger(logger))
	router.Use(middleware.Recoverer)
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})
	simaHandler := handler.NewSimaHandler()
	simaRouter := simaHandler.RegisterRoutes()
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/sima", simaRouter)
	})
	return router
}
