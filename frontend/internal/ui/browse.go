package ui

import (
	"fmt"
	"strings"

	"github.com/bork/frontend/internal/api"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BackToDashboardMsg is sent when user goes back to dashboard
type BackToDashboardMsg struct{}

// BooksLoadedMsg is sent when books are loaded
type BooksLoadedMsg struct {
	Books []api.Book
}

// BooksErrorMsg is sent when books fail to load
type BooksErrorMsg struct {
	Error string
}

// BrowseBooksView represents the browse books screen
type BrowseBooksView struct {
	apiClient    *api.Client
	user         *api.UserInfo
	books        []api.Book
	currentPage  int
	booksPerPage int
	selectedBook int
	isLoading    bool
	errorMsg     string
}

// NewBrowseBooksView creates a new browse books view
func NewBrowseBooksView(apiClient *api.Client, user *api.UserInfo) *BrowseBooksView {
	return &BrowseBooksView{
		apiClient:    apiClient,
		user:         user,
		books:        []api.Book{},
		currentPage:  0,
		booksPerPage: 4,
		selectedBook: 0,
		isLoading:    true,
	}
}

// Init initializes the browse books view
func (b *BrowseBooksView) Init() tea.Cmd {
	// Cycle to next quote
	GetQuoteManager().NextQuote()
	return b.loadBooks()
}

// Update handles messages for the browse books view
func (b *BrowseBooksView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case BooksLoadedMsg:
		b.isLoading = false
		b.books = msg.Books
		return b, nil

	case BooksErrorMsg:
		b.isLoading = false
		b.errorMsg = msg.Error
		return b, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if b.selectedBook > 0 {
				b.selectedBook--
			}
			return b, nil

		case "down", "j":
			maxIndex := b.getBooksOnCurrentPage() - 1
			if b.selectedBook < maxIndex {
				b.selectedBook++
			}
			return b, nil

		case "left", "h", "pgup":
			// Previous page
			if b.currentPage > 0 {
				b.currentPage--
				b.selectedBook = 0
			}
			return b, nil

		case "right", "l", "pgdown":
			// Next page
			totalPages := b.getTotalPages()
			if b.currentPage < totalPages-1 {
				b.currentPage++
				b.selectedBook = 0
			}
			return b, nil

		case "r":
			// Refresh books
			b.isLoading = true
			b.errorMsg = ""
			return b, b.loadBooks()

		case "b", "backspace":
			// Go back to dashboard
			return b, func() tea.Msg {
				return BackToDashboardMsg{}
			}
		}
	}

	return b, nil
}

// View renders the browse books screen
func (b *BrowseBooksView) View() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(MatrixGreen).
		Background(MatrixBlack).
		Padding(0, 2)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(MatrixGreen).
		Background(MatrixBlack).
		Padding(1, 2).
		Width(74)

	bookStyle := lipgloss.NewStyle().
		Foreground(MatrixDarkGreen).
		Background(MatrixBlack).
		Padding(1, 2).
		Width(68).
		MarginBottom(1)

	selectedBookStyle := bookStyle.Copy().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(MatrixHighlight).
		Background(MatrixGray).
		Width(68)

	availableStyle := lipgloss.NewStyle().
		Foreground(MatrixHighlight).
		Bold(true)

	rentedStyle := lipgloss.NewStyle().
		Foreground(MatrixRed)

	helpStyle := lipgloss.NewStyle().
		Foreground(MatrixDarkGreen).
		Align(lipgloss.Center).
		MarginTop(1)

	var view strings.Builder

	// Header
	header := fmt.Sprintf("BORK - Browse Books                 User: %s", b.user.Username)
	view.WriteString(headerStyle.Render(header))
	view.WriteString("\n")

	// Matrix quote
	quote := GetQuoteManager().GetQuote()
	quoteStyle := lipgloss.NewStyle().
		Foreground(MatrixDarkGreen).
		Italic(true).
		Align(lipgloss.Center).
		Width(74)
	view.WriteString(quoteStyle.Render(fmt.Sprintf("\" %s \"", quote)))
	view.WriteString("\n\n")

	// Books list
	var booksContent strings.Builder

	if b.isLoading {
		booksContent.WriteString(MatrixHighlightText.Render("Loading books..."))
	} else if b.errorMsg != "" {
		booksContent.WriteString(MatrixError.Render("Error: " + b.errorMsg))
	} else if len(b.books) == 0 {
		booksContent.WriteString(MatrixText.Render("No books available"))
	} else {
		// Pagination info
		totalPages := b.getTotalPages()
		pageInfo := fmt.Sprintf("Page %d of %d | Total Books: %d", b.currentPage+1, totalPages, len(b.books))
		booksContent.WriteString(lipgloss.NewStyle().Bold(true).Render(pageInfo))
		booksContent.WriteString("\n\n")

		// Display books for current page
		startIdx := b.currentPage * b.booksPerPage
		endIdx := startIdx + b.booksPerPage
		if endIdx > len(b.books) {
			endIdx = len(b.books)
		}

		for i := startIdx; i < endIdx; i++ {
			book := b.books[i]
			localIdx := i - startIdx

			status := "[RENTED]"
			statusStyle := rentedStyle
			if book.IsAvailable {
				status = "[AVAILABLE]"
				statusStyle = availableStyle
			}

			var bookContent strings.Builder
			bookContent.WriteString(fmt.Sprintf("📖 %s ", book.Title))
			bookContent.WriteString(statusStyle.Render(status))
			bookContent.WriteString("\n")
			bookContent.WriteString(fmt.Sprintf("   Author: %s\n", book.Author))
			bookContent.WriteString(fmt.Sprintf("   Category: %s\n", book.Category.Name))
			bookContent.WriteString(fmt.Sprintf("   ISBN: %s", book.ISBN))

			if localIdx == b.selectedBook {
				booksContent.WriteString(selectedBookStyle.Render(bookContent.String()))
			} else {
				booksContent.WriteString(bookStyle.Render(bookContent.String()))
			}
			booksContent.WriteString("\n")
		}
	}

	view.WriteString(boxStyle.Render(booksContent.String()))
	view.WriteString("\n")

	// Help
	helpText := "↑/↓: Navigate | ←/→: Page | R: Refresh | B: Back | Esc: Exit"
	view.WriteString(helpStyle.Render(helpText))

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(view.String())
}

// loadBooks loads available books from the API
func (b *BrowseBooksView) loadBooks() tea.Cmd {
	return func() tea.Msg {
		books, err := b.apiClient.GetAvailableBooks()
		if err != nil {
			return BooksErrorMsg{Error: err.Error()}
		}
		return BooksLoadedMsg{Books: books}
	}
}

// getTotalPages calculates total number of pages
func (b *BrowseBooksView) getTotalPages() int {
	if len(b.books) == 0 {
		return 1
	}
	return (len(b.books) + b.booksPerPage - 1) / b.booksPerPage
}

// getBooksOnCurrentPage returns the number of books on current page
func (b *BrowseBooksView) getBooksOnCurrentPage() int {
	startIdx := b.currentPage * b.booksPerPage
	endIdx := startIdx + b.booksPerPage
	if endIdx > len(b.books) {
		endIdx = len(b.books)
	}
	return endIdx - startIdx
}
