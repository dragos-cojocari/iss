package api

import "net/http"

// GetAvailableBooks retrieves all available books
func (c *Client) GetAvailableBooks() ([]Book, error) {
	var books []Book
	if err := c.doRequest(http.MethodGet, "/api/books/available", nil, &books); err != nil {
		return nil, err
	}

	return books, nil
}

// GetAllBooks retrieves all books
func (c *Client) GetAllBooks() ([]Book, error) {
	var books []Book
	if err := c.doRequest(http.MethodGet, "/api/books", nil, &books); err != nil {
		return nil, err
	}

	return books, nil
}

// SearchBooks searches for books by title or author
func (c *Client) SearchBooks(query string) ([]Book, error) {
	var books []Book
	path := "/api/books/search?q=" + query
	if err := c.doRequest(http.MethodGet, path, nil, &books); err != nil {
		return nil, err
	}

	return books, nil
}
