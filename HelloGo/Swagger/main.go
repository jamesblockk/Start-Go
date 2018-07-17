package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jamesblockk/Start-Go/HelloGo/Swagger/models"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", helloWorld).Methods("GET")
	router.HandleFunc("/usr/{id:[0-9]+}", GetUsr).Methods("GET")

	router.Path("/usr/{id:[0-9]+}").
		HandlerFunc(CreateUsr).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {

}

func GetUsr(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usr := models.GetUsrReq{ID: params["id"]}

	respondJSON(w, http.StatusCreated, usr)

}

func CreateUsr(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	act := r.FormValue("act")
	pwd := r.FormValue("pwd")
	fmt.Println(id, act, pwd)
	usr := models.CreateUsrReq{Act: act, Pwd: pwd}
	// respondError(w, http.StatusBadRequest, "Invalid request payload")

	respondJSON(w, http.StatusCreated, usr)
}

// send a payload of JSON content
func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// send a JSON error message
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
