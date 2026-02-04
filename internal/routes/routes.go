package routes

import (
	"io"
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

func SetupRoutes(origin string) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path
		cache.RLock()
		cached, ok := cache.data[key]
		cache.RUnlock()
		if ok {
			// cache hit
			for k, vv := range cached.header {
				for _, v := range vv {
					w.Header().Add(k, v)
				}
			}
			w.Header().Set("X-Cache", "HIT")
			w.WriteHeader(cached.status)
			w.Write(cached.body)
			return
		}

		resp, err := http.Get(origin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.Header().Set("X-Cache", "MISS")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
		cache.Lock()
		cache.data[key] = cachedResponse{
			status: resp.StatusCode,
			header: resp.Header.Clone(),
			body:   body,
		}
		cache.Unlock()

	})
	return r
}
