package user

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var userRegexp = regexp.MustCompile(`^\/users\/(\d+)$`)

	if matches := userRegexp.FindStringSubmatch(r.URL.Path); len(matches) == 0 {
		handleUsers(w, r)
	} else {
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handleUser(w, r, id)
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Add("content-type", "application/json")
		msg, err := json.Marshal(getAll())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write(msg)
	case http.MethodPost:
		var user User
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user = add(user)

		msg, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(msg)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleUser(w http.ResponseWriter, r *http.Request, id int) {
	switch r.Method {
	case http.MethodGet:
		u, err := getOne(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		data, err := json.Marshal(u)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		w.Write(data)
		return
	case http.MethodPut:
		var user User
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		u, err := update(user, id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		data, err := json.Marshal(u)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("content-type", "application/json")
		w.Write(data)

	case http.MethodDelete:
		if !delete(id) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
