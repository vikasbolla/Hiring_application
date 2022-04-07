package routes

import (
	"log"
	"net/http"

	"github.com/vikasbolla/Hiring_application/Hiring_Application/service"

	"github.com/gorilla/mux"
)

func Handle() {
	myrouter := mux.NewRouter()
	//myrouter.Handle("/getall", GetallDetails).Methods("GET")
	myrouter.HandleFunc("/external-all", service.ExternalAllCandidates).Methods("GET")
	myrouter.HandleFunc("/external-post-requirement", service.ExternalPostRequirement).Methods("POST")
	myrouter.HandleFunc("/external-match-requirement", service.ExternalMatchRequirement).Methods("POST")

	myrouter.HandleFunc("/internal-all", service.InternalAllCandidates).Methods("GET")
	myrouter.HandleFunc("/internal-post-requirement", service.InternalPostRequirement).Methods("POST")
	myrouter.HandleFunc("/internal-match-requirement-optional", service.InternalMatchRequirement).Methods("POST")
	log.Fatal(http.ListenAndServe(":9001", myrouter))
}
