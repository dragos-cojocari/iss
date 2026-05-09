package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("http://localhost:8080")

	if client == nil {
		t.Fatal("Expected client to be created, got nil")
	}

	if client.baseURL != "http://localhost:8080" {
		t.Errorf("Expected baseURL to be 'http://localhost:8080', got '%s'", client.baseURL)
	}

	if client.httpClient == nil {
		t.Error("Expected httpClient to be initialized")
	}

	if client.httpClient.Jar == nil {
		t.Error("Expected cookie jar to be initialized")
	}
}

func TestDoRequest_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if r.URL.Path != "/api/test" {
			t.Errorf("Expected path '/api/test', got '%s'", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	var result map[string]string
	err := client.doRequest(http.MethodGet, "/api/test", nil, &result)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result["message"] != "success" {
		t.Errorf("Expected message 'success', got '%s'", result["message"])
	}
}

func TestDoRequest_WithBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("Expected Content-Type header to be 'application/json'")
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	body := map[string]string{"key": "value"}
	err := client.doRequest(http.MethodPost, "/api/test", body, nil)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDoRequest_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"status": 400, "message": "Bad request"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	err := client.doRequest(http.MethodGet, "/api/test", nil, nil)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	expectedMsg := "Bad request (status 400)"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

func TestDoRequest_404Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := NewClient(server.URL)

	err := client.doRequest(http.MethodGet, "/api/notfound", nil, nil)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
