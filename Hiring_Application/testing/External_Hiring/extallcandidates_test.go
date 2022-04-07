package testing

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/vikasbolla/Hiring_application/Hiring_Application/service"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/external-all", service.ExternalAllCandidates).Methods("GET")
	return router
}

func TestExtAllCandidates(t *testing.T) {
	request, _ := http.NewRequest("GET", "/external-all", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "*OK* response is expected")
}
