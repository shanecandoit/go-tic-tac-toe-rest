package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {

	InitBoard()

	req, err := http.NewRequest("GET", "/boards", nil)
	if err != nil {
		t.Fatal(err)
	}
	tr := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(tr, req)
	if status := tr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[[0,0,0,0,0,0,0,0,0],[0,0,0,0,0,0,-1,0,0],[0,0,0,0,1,0,-1,0,0]]`
	if tr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			tr.Body.String(), expected)
	}
}
