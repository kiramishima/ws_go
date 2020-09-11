package http

import (
	"MenuDigital/controllers/auth"
	"github.com/go-chi/chi"
)

func RegisterHTTPEndpoints(router *chi.Mux, uc auth.UseCase) {
	h := NewHandler(uc)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.SignUp)
		r.Post("/sign-in", h.SignIn)
	})
}
