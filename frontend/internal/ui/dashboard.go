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

// BrowseBooksMsg is sent when user selects Browse Books
type BrowseBooksMsg struct{}

// DashboardView represents the main dashboard
type DashboardView struct {
	apiClient    *api.Client
	user         *api.UserInfo
	selectedMenu int
}

// NewDashboardView creates a new dashboard view
func NewDashboardView(apiClient *api.Client, user *api.UserInfo) *DashboardView {
	return &DashboardView{
		apiClient:    apiClient,
		user:         user,
		selectedMenu: 0,
	}
}

// Init initializes the dashboard view
func (d *DashboardView) Init() tea.Cmd {
	// Cycle to next quote
	GetQuoteManager().NextQuote()
	return nil
}

// Update handles messages for the dashboard view
func (d *DashboardView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if d.selectedMenu > 0 {
				d.selectedMenu--
			}
			return d, nil

		case "down", "j":
			if d.selectedMenu < 3 {
				d.selectedMenu++
			}
			return d, nil

		case "1":
			d.selectedMenu = 0
			return d, func() tea.Msg {
				return BrowseBooksMsg{}
			}

		case "2":
			d.selectedMenu = 1
			return d, nil

		case "3":
			d.selectedMenu = 2
			return d, nil

		case "4":
			d.selectedMenu = 3
			return d, func() tea.Msg {
				d.apiClient.Logout()
				return LogoutMsg{}
			}

		case "enter":
			// Execute selected menu item
			if d.selectedMenu == 0 {
				return d, func() tea.Msg {
					return BrowseBooksMsg{}
				}
			} else if d.selectedMenu == 3 {
				return d, func() tea.Msg {
					d.apiClient.Logout()
					return LogoutMsg{}
				}
			}
			return d, nil
		}
	}

	return d, nil
}

// View renders the dashboard
func (d *DashboardView) View() string {
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

	menuItemStyle := lipgloss.NewStyle().
		Foreground(MatrixDarkGreen).
		Background(MatrixBlack).
		Padding(0, 2).
		Width(70)

	selectedMenuStyle := menuItemStyle.Copy().
		Foreground(MatrixBlack).
		Background(MatrixGreen).
		Bold(true).
		Width(70)

	helpStyle := lipgloss.NewStyle().
		Foreground(MatrixDarkGreen).
		Align(lipgloss.Center).
		MarginTop(1)

	var view strings.Builder

	// Header
	header := fmt.Sprintf("BORK - Main Menu                    User: %s", d.user.Username)
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

	// Current Rentals (placeholder)
	rentalsBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(MatrixDarkGreen).
		Background(MatrixBlack).
		Padding(1, 2).
		Width(74)

	var rentalsContent strings.Builder
	rentalsContent.WriteString(lipgloss.NewStyle().Bold(true).Foreground(MatrixGreen).Render("MY CURRENT RENTALS (0/3)"))
	rentalsContent.WriteString("\n\n")
	rentalsContent.WriteString(lipgloss.NewStyle().Foreground(MatrixDarkGreen).Render("No active rentals"))

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
	view.WriteString("\n")

	// Help
	view.WriteString(helpStyle.Render("↑/↓: Navigate | Enter/1-4: Select | Esc: Exit"))

	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(view.String())
}
