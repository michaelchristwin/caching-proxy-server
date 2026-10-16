package handler

import (
	"caching-proxy/internal/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type SimaHandler struct{}

func NewSimaHandler() *SimaHandler {
	return &SimaHandler{}
}

func (h *SimaHandler) RegisterRoutes() *chi.Mux {
	simaRouter := chi.NewRouter()
	simaRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {

		model.Ok(w, r, map[string]string{"message": "Cao-Cao says hi"})
	})
	return simaRouter
}
