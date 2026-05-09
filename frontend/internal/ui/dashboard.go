package ui

import (
	"fmt"
	"strings"

	"github.com/bork/frontend/internal/api"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// LogoutMsg is sent when user logs out
type LogoutMsg struct{}

// BooksLoadedMsg is sent when books are loaded
type BooksLoadedMsg struct {
	Books []api.Book
}

// BooksErrorMsg is sent when books fail to load
type BooksErrorMsg struct {
	Error string
}

// DashboardView represents the main dashboard
type DashboardView struct {
	apiClient    *api.Client
	user         *api.UserInfo
	books        []api.Book
	selectedMenu int
	isLoading    bool
	errorMsg     string
}

// NewDashboardView creates a new dashboard view
func NewDashboardView(apiClient *api.Client, user *api.UserInfo) *DashboardView {
	return &DashboardView{
		apiClient:    apiClient,
		user:         user,
		books:        []api.Book{},
		selectedMenu: 0,
		isLoading:    true,
	}
}

// Init initializes the dashboard view
func (d *DashboardView) Init() tea.Cmd {
	return d.loadBooks()
}

// Update handles messages for the dashboard view
func (d *DashboardView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case BooksLoadedMsg:
		d.isLoading = false
		d.books = msg.Books
		return d, nil

	case BooksErrorMsg:
		d.isLoading = false
		d.errorMsg = msg.Error
		return d, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			d.selectedMenu = 0
			return d, nil
		case "2":
			d.selectedMenu = 1
			return d, nil
		case "3":
			d.selectedMenu = 2
			return d, nil
		case "4":
			// Logout
			return d, func() tea.Msg {
				d.apiClient.Logout()
				return LogoutMsg{}
			}
		case "r":
			// Refresh books
			d.isLoading = true
			d.errorMsg = ""
			return d, d.loadBooks()
		}
	}

	return d, nil
}

// View renders the dashboard
func (d *DashboardView) View() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00BFFF")).
		Background(lipgloss.Color("#1a1a1a")).
		Padding(0, 2)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00BFFF")).
		Padding(1, 2).
		Width(74)

	menuItemStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Padding(0, 2)

	selectedMenuStyle := menuItemStyle.Copy().
		Foreground(lipgloss.Color("#00BFFF")).
		Bold(true)

	bookStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		MarginBottom(1)

	availableStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4CAF50")).
		Bold(true)

	rentedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF6B6B"))

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Align(lipgloss.Center).
		MarginTop(1)

	var view strings.Builder

	// Header
	header := fmt.Sprintf("BORK - Main Menu                    User: %s", d.user.Username)
	view.WriteString(headerStyle.Render(header))
	view.WriteString("\n\n")

	// Current Rentals (placeholder)
	rentalsBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#666666")).
		Padding(1, 2).
		Width(74)

	var rentalsContent strings.Builder
	rentalsContent.WriteString(lipgloss.NewStyle().Bold(true).Render("MY CURRENT RENTALS (0/3)"))
	rentalsContent.WriteString("\n\n")
	rentalsContent.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Render("No active rentals"))

	view.WriteString(rentalsBox.Render(rentalsContent.String()))
	view.WriteString("\n\n")

	// Menu
	var menuContent strings.Builder
	menuContent.WriteString(lipgloss.NewStyle().Bold(true).Render("MENU"))
	menuContent.WriteString("\n\n")

	menuItems := []string{
		"[1] 📚 Browse Books",
		"[2] 🛒 View Cart (0 items) - Coming Soon",
		"[3] 📋 My Rentals - Coming Soon",
		"[4] 🚪 Logout",
	}

	for i, item := range menuItems {
		if i == d.selectedMenu {
			menuContent.WriteString(selectedMenuStyle.Render(item))
		} else {
			menuContent.WriteString(menuItemStyle.Render(item))
		}
		menuContent.WriteString("\n")
	}

	view.WriteString(boxStyle.Render(menuContent.String()))
	view.WriteString("\n\n")

	// Available Books
	var booksContent strings.Builder
	booksContent.WriteString(lipgloss.NewStyle().Bold(true).Render("AVAILABLE BOOKS"))
	booksContent.WriteString("\n\n")

	if d.isLoading {
		booksContent.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#FFB74D")).Render("Loading books..."))
	} else if d.errorMsg != "" {
		booksContent.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF6B6B")).Render("Error: " + d.errorMsg))
	} else if len(d.books) == 0 {
		booksContent.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Render("No books available"))
	} else {
		// Display first 5 books
		displayCount := len(d.books)
		if displayCount > 5 {
			displayCount = 5
		}

		for i := 0; i < displayCount; i++ {
			book := d.books[i]
			status := "[RENTED]"
			statusStyle := rentedStyle
			if book.IsAvailable {
				status = "[AVAILABLE]"
				statusStyle = availableStyle
			}

			bookLine := fmt.Sprintf("📖 %s", book.Title)
			booksContent.WriteString(bookStyle.Render(bookLine))
			booksContent.WriteString(" ")
			booksContent.WriteString(statusStyle.Render(status))
			booksContent.WriteString("\n")

			authorLine := fmt.Sprintf("   Author: %s", book.Author)
			booksContent.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#999999")).Render(authorLine))
			booksContent.WriteString("\n")

			categoryLine := fmt.Sprintf("   Category: %s", book.Category.Name)
			booksContent.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#999999")).Render(categoryLine))
			booksContent.WriteString("\n\n")
		}

		if len(d.books) > 5 {
			moreText := fmt.Sprintf("... and %d more books", len(d.books)-5)
			booksContent.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#666666")).Italic(true).Render(moreText))
		}
	}

	view.WriteString(boxStyle.Render(booksContent.String()))
	view.WriteString("\n")

	// Help
	view.WriteString(helpStyle.Render("1-4: Select Option | R: Refresh | Esc: Exit"))

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(view.String())
}

// loadBooks loads available books from the API
func (d *DashboardView) loadBooks() tea.Cmd {
	return func() tea.Msg {
		books, err := d.apiClient.GetAvailableBooks()
		if err != nil {
			return BooksErrorMsg{Error: err.Error()}
		}
		return BooksLoadedMsg{Books: books}
	}
}
