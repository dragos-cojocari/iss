package api

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	SessionID string   `json:"sessionId"`
	User      UserInfo `json:"user"`
}

// UserInfo represents user information
type UserInfo struct {
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Book represents a book in the library
type Book struct {
	BookID      string   `json:"bookId"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	ISBN        string   `json:"isbn"`
	Category    Category `json:"category"`
	IsAvailable bool     `json:"isAvailable"`
	AddedAt     string   `json:"addedAt"` // Keep as string to avoid parsing issues
}

// Category represents a book category
type Category struct {
	CategoryID  string `json:"categoryId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"` // Keep as string to avoid parsing issues
}

// ErrorResponse represents an API error response
type ErrorResponse struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
}
