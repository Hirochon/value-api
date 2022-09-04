package rest_api

import (
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	httpRouter := initCustomRouter()
	testServer := httptest.NewServer(httpRouter)
	defer testServer.Close()

	testRequest := httptest.NewRequest("GET", testServer.URL+"/health", nil)
	testResponse := httptest.NewRecorder()

	httpRouter.ServeHTTP(testResponse, testRequest)

	if testResponse.Code != 200 {
		t.Errorf("status code is not 200. got %v", testResponse.Code)
	}
}
