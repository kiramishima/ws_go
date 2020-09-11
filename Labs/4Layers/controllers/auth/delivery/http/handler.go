package http

import (
	"MenuDigital/controllers/auth"
	"encoding/json"
	"github.com/unrolled/render"
	"net/http"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type signInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	inp := new(signInput)
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(inp); err != nil {
		render.New().JSON(w, http.StatusInternalServerError, "")
		return
	}

	render.New().JSON(w, http.StatusOK, "")
}

type signInResponse struct {
	Token string `json:"token"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	inp := new(signInput)

	if err := json.NewDecoder(r.Body).Decode(inp); err != nil {
		render.New().JSON(w, http.StatusBadRequest, "")
		return
	}

	token, err := h.useCase.SignIn(ct, inp.Email, inp.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			render.New().JSON(w, http.StatusUnauthorized, "")
			return
		}

		render.New().JSON(w, http.StatusInternalServerError, "")
		return
	}
	render.New().JSON(w, http.StatusOK, signInResponse{Token: token})
}
