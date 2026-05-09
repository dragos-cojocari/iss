package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllBooks_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		if r.URL.Path != "/api/books" {
			t.Errorf("Expected path '/api/books', got '%s'", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[
			{
				"bookId": "book-1",
				"title": "The Matrix",
				"author": "Wachowski Sisters",
				"isbn": "978-0-123456-78-9",
				"isAvailable": true,
				"category": {
					"categoryId": "cat-1",
					"name": "Science Fiction"
				}
			},
			{
				"bookId": "book-2",
				"title": "Neuromancer",
				"author": "William Gibson",
				"isbn": "978-0-987654-32-1",
				"isAvailable": false,
				"category": {
					"categoryId": "cat-1",
					"name": "Science Fiction"
				}
			}
		]`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	books, err := client.GetAllBooks()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(books) != 2 {
		t.Fatalf("Expected 2 books, got %d", len(books))
	}

	if books[0].Title != "The Matrix" {
		t.Errorf("Expected first book title 'The Matrix', got '%s'", books[0].Title)
	}

	if books[0].IsAvailable != true {
		t.Error("Expected first book to be available")
	}

	if books[1].Title != "Neuromancer" {
		t.Errorf("Expected second book title 'Neuromancer', got '%s'", books[1].Title)
	}
}

func TestGetAllBooks_EmptyList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	books, err := client.GetAllBooks()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(books))
	}
}

func TestSearchBooks_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		expectedPath := "/api/books/search"
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path '%s', got '%s'", expectedPath, r.URL.Path)
		}

		query := r.URL.Query().Get("q")
		if query != "Matrix" {
			t.Errorf("Expected query 'Matrix', got '%s'", query)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[
			{
				"bookId": "book-1",
				"title": "The Matrix",
				"author": "Wachowski Sisters",
				"isbn": "978-0-123456-78-9",
				"isAvailable": true,
				"category": {
					"categoryId": "cat-1",
					"name": "Science Fiction"
				}
			}
		]`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	books, err := client.SearchBooks("Matrix")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(books) != 1 {
		t.Fatalf("Expected 1 book, got %d", len(books))
	}

	if books[0].Title != "The Matrix" {
		t.Errorf("Expected book title 'The Matrix', got '%s'", books[0].Title)
	}
}

func TestSearchBooks_NoResults(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()

	client := NewClient(server.URL)

	books, err := client.SearchBooks("NonExistent")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(books) != 0 {
		t.Errorf("Expected 0 books, got %d", len(books))
	}
}

func TestSearchBooks_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(server.URL)

	books, err := client.SearchBooks("test")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if books != nil {
		t.Error("Expected nil books on error")
	}
}
