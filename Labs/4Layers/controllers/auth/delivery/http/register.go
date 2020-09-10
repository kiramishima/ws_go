package http

import (
	"github.com/gin-gonic/gin"
	"MenuDigital/controllers/auth"
	"github.com/go-chi/chi"
)

func RegisterHTTPEndpoints(router *chi.Router, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
	}
}