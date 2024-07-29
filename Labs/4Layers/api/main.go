package main

import (
	"MenuDigital/api/handlers"
	"MenuDigital/config"
	"MenuDigital/pkg/users"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	// "MenuDigital/config"
)

var (
	DB_USER   = os.Getenv("DB_USER")
	DB_PWD    = os.Getenv("DB_PASSWORD")
	DB_SERVER = os.Getenv("DB_SERVER")
	DB        = os.Getenv("DB")
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("SecretKey")), nil)
}

func main() {
	// mysql
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_DATABASE))
	db.SetMaxOpenConns(config.MYSQL_CONNECTION_POOL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	// router
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	// Service
	usersRepo := users.NewMySQLRepository(db)
	usersService := users.NewService(usersRepo)

	// Handlers
	handlers.MakeUsersHandlers(r, usersService)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		println("Pong")
		w.WriteHeader(http.StatusOK)
	})

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      http.DefaultServeMux,
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
