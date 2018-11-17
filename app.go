package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func createExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func updateExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func deleteExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/exchages", getAllExchangesEndpoint).Methods("GET")
	r.HandleFunc("/exchages", createExchangesEndpoint).Methods("POST")
	r.HandleFunc("/exchages/{id}", updateExchangesEndpoint).Methods("PUT")
	r.HandleFunc("/exchages/{id}", deleteExchangesEndpoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
