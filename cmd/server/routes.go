package main

import (
	"caching-proxy/internal/handler"
	"encoding/json"
	"net/http"
	"sync"

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
	router.Use(middleware.Logger)
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
