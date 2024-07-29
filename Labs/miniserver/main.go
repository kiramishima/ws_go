package miniserver

import (
	"fmt"
	"log"
	"net/http"
	"flag"
)

var (
	MEDIA_DIR = flag.String("MDir", "", "MDir = /media_files")
	PORT = flag.Int("Port", 3100, "Port = 8080")
)

func init() {
	fmt.Println("Bonjour vers l'init fonction")
	flag.Parse()
}

func main() {
	http.Handle("/", addHeaders(http.FileServer(http.Dir(MEDIA_DIR))))
	fmt.Printf("Starting server on %v\n", PORT)
	log.Printf("Serving %s on HTTP port: %v\n", MEDIA_DIR, PORT)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}