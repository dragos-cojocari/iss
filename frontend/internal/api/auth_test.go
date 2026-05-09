package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		if r.URL.Path != "/api/auth/login" {
			t.Errorf("Expected path '/api/auth/login', got '%s'", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"sessionId": "test-session-id",
			"user": {
				"userId": "user-123",
				"username": "testuser",
				"email": "test@example.com",
				"firstName": "Test",
				"lastName": "User"
			}
		}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	resp, err := client.Login("testuser", "password")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp == nil {
		t.Fatal("Expected response, got nil")
	}

	if resp.SessionID != "test-session-id" {
		t.Errorf("Expected sessionId 'test-session-id', got '%s'", resp.SessionID)
	}

	if resp.User.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", resp.User.Username)
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"status": 401, "message": "Invalid credentials"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	resp, err := client.Login("wronguser", "wrongpass")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if resp != nil {
		t.Error("Expected nil response on error")
	}
}

func TestLogout_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		if r.URL.Path != "/api/auth/logout" {
			t.Errorf("Expected path '/api/auth/logout', got '%s'", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.URL)

	err := client.Logout()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetCurrentUser_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if r.URL.Path != "/api/auth/me" {
			t.Errorf("Expected path '/api/auth/me', got '%s'", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
			"userId": "user-123",
			"username": "testuser",
			"email": "test@example.com",
			"firstName": "Test",
			"lastName": "User"
		}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	user, err := client.GetCurrentUser()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if user == nil {
		t.Fatal("Expected user, got nil")
	}

	if user.Username != "testuser" {
		t.Errorf("Expected username 'testuser', got '%s'", user.Username)
	}

	if user.Email != "test@example.com" {
		t.Errorf("Expected email 'test@example.com', got '%s'", user.Email)
	}
}

func TestGetCurrentUser_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"status": 401, "message": "Unauthorized"}`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	user, err := client.GetCurrentUser()

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if user != nil {
		t.Error("Expected nil user on error")
	}
}
