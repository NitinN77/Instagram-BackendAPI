package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/posts/61609550ecb29a3bbf21a034", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"_id":"61609550ecb29a3bbf21a034","caption":"NOSTEDI2","image":"NOSDTE22","timestamp":"2021-10-08T19:00:32.264Z","author":"NOSTEDIER"}`

	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreatePost(t *testing.T) {
	var jsonStr = []byte(`{"Caption":"testpost","Image":"test.png","Author":"nitin"}`)
	req, err := http.NewRequest("POST", "/api/posts", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreatePost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"InsertedID":"61617b6a91e6558a630edc60"}`
	if len(strings.TrimSpace(rr.Body.String())) != len(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"Name":"testuser","Email":"testuser.com","Password":"testpass"}`)
	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"InsertedID":"61617eaf8d9c98800c05de5f"}`
	if len(strings.TrimSpace(rr.Body.String())) != len(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/users/61611de1db4f6a84611d8a08", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"_id":"61611de1db4f6a84611d8a08","name":"nitin","email":"nostedi.com","password":"$2a$10$pF184F28Zh/3h95HAxEYsOYLORYmAu7i7KNil8K2r.Q3eEd7yu/Mu"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetPostsByUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/posts/users/616092d22ddd1b01a79a4653?limit=2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetPostsByUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"Posts":[{"_id":"61609535ecb29a3bbf21a032","caption":"STEDI","image":"SDTE","timestamp":"2021-10-08T19:00:05.148Z","author":"STEDIER"},{"_id":"61611c84bab9653aeaae0cc6","caption":"seres","image":"sefsf","timestamp":"2021-10-09T04:37:24.687Z","author":"STEDIER"}],"lowerId":"61611c84bab9653aeaae0cc6"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
