package http

import (
	"MenuDigital/controllers/auth"
	"github.com/unrolled/render"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	usecase auth.UseCase
}

func NewAuthMiddleware(usecase auth.UseCase) http.Handler {
	return (&AuthMiddleware{
		usecase: usecase,
	}).Handle()
}

func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := w.Header().Get("Authorization")
		if authHeader == "" {
			render.New().JSON(w, http.StatusInternalServerError, "")
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			render.New().JSON(w, http.StatusUnauthorized, "")
			return
		}

		if headerParts[0] != "Bearer" {
			render.New().JSON(w, http.StatusUnauthorized, "")
			return
		}

		user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
		if err != nil {
			status := http.StatusInternalServerError
			if err == auth.ErrInvalidAccessToken {
				status = http.StatusUnauthorized
			}
			render.New().JSON(w, status, "")
			return
		}

		w.Header().Set(auth.CtxUserKey, user.ID)
	})
}
