package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/gorilla/mux"
)

type app struct {
	Router *mux.Router
	DB     *mgo.Database
}

func (a *app) getAllExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	exchanges := exchange{}
	result, err := exchanges.getAll(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, result)
}

func (a *app) createExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	exchange := exchange{}
	if err := json.NewDecoder(r.Body).Decode(&exchange); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	exchange.ID = bson.NewObjectId()
	if err := exchange.insert(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, exchange)
}

func (a *app) updateExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	exchange := exchange{}
	if err := json.NewDecoder(r.Body).Decode(&exchange); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := exchange.update(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *app) deleteExchangesEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	exchange := exchange{}
	if err := json.NewDecoder(r.Body).Decode(&exchange); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := exchange.delete(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *app) initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/exchange", a.getAllExchangesEndpoint).Methods("GET")
	a.Router.HandleFunc("/exchange", a.createExchangesEndpoint).Methods("POST")
	a.Router.HandleFunc("/exchange", a.updateExchangesEndpoint).Methods("PUT")
	a.Router.HandleFunc("/exchange", a.deleteExchangesEndpoint).Methods("DELETE")
	connection := connector{
		Server:   "mongodb://192.168.33.10:27017",
		Database: "currency",
	}
	a.DB = connection.connect()
	if err := http.ListenAndServe(":3000", a.Router); err != nil {
		log.Fatal(err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
