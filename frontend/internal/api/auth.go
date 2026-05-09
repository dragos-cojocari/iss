package api

import "net/http"

// Login authenticates the user and stores the session cookie
func (c *Client) Login(username, password string) (*LoginResponse, error) {
	req := LoginRequest{
		Username: username,
		Password: password,
	}

	var resp LoginResponse
	if err := c.doRequest(http.MethodPost, "/api/auth/login", req, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Logout invalidates the current session
func (c *Client) Logout() error {
	return c.doRequest(http.MethodPost, "/api/auth/logout", nil, nil)
}

// GetCurrentUser retrieves the currently authenticated user
func (c *Client) GetCurrentUser() (*UserInfo, error) {
	var user UserInfo
	if err := c.doRequest(http.MethodGet, "/api/auth/me", nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
