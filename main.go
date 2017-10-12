package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Foo.")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handleHome).Methods("GET")

	s := &http.Server{
		Handler:        r,
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening on :8080")
	log.Fatal(s.ListenAndServe())
}
