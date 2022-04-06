package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handle() {
	myrouter := mux.NewRouter()
	//myrouter.Handle("/getall", GetallDetails).Methods("GET")
	myrouter.HandleFunc("/external-all", ExternalAllCandidates).Methods("GET")
	myrouter.HandleFunc("/external-post-requirement", ExternalPostRequirement).Methods("POST")
	myrouter.HandleFunc("/external-match-requirement", ExternalMatchRequirement).Methods("POST")

	myrouter.HandleFunc("/internal-all", InternalAllCandidates).Methods("GET")
	myrouter.HandleFunc("/internal-post-requirement", InternalPostRequirement).Methods("POST")
	myrouter.HandleFunc("/internal-match-requirement-optional", InternalMatchRequirement).Methods("POST")
	log.Fatal(http.ListenAndServe(":9001", myrouter))
}
