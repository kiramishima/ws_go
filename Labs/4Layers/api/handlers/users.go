package handlers

import (
	"MenuDigital/pkg/entities"
	"MenuDigital/pkg/users"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func usersIndex(service users.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		ctx := r.Context()
		var data []entities.User
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.FindAll(ctx)
		default:
			data, err = service.Search(ctx, name)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entities.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
		}

		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	}
}

// MakeUsersHandlers Make url Handlers
func MakeUsersHandlers(r *chi.Mux, service users.UseCase) {

	r.Get("/v1/users", usersIndex(service))
}
