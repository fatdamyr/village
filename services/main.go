package main

import (
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/gorilla/mux"
)

var version = "1.0.0"

func versionEndpoint(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(fmt.Sprintf("{version: '%s'}", version)))
}

func healthzEndpoint(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte("{status: 'healthy'}"))
}

func main() {
	log.Println("Starting app...")

	/*
	   Configure the application here. This could include reading ENV variables, loading
	   a configuration files, etc.
	*/
	httpAddr := os.Getenv("HTTP_ADDR")

	router := mux.NewRouter()
	router.HandleFunc("/version", versionEndpoint).Methods("GET")
	router.HandleFunc("/healthz", healthzEndpoint).Methods("GET")

	/* Default Handler */
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, html, version)
	}).Methods("GET")

	/*
		Start the HTTP Server
	*/
	log.Printf("HTTP Service listening on %s", httpAddr)
	httpErr := http.ListenAndServe(httpAddr, router)
	if httpErr != nil {
		log.Fatal(httpErr)
	}

}
