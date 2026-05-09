package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// App represents the main application model
type App struct {
	currentView View
	loginView   *LoginView
	width       int
	height      int
}

// View represents different screens in the application
type View int

const (
	LoginViewType View = iota
	// Future views: DashboardView, BrowseBooksView, etc.
)

// NewApp creates a new application instance
func NewApp() *App {
	return &App{
		currentView: LoginViewType,
		loginView:   NewLoginView(),
	}
}

// Init initializes the application
func (a *App) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the application state
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		return a, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return a, tea.Quit
		}
	}

	// Delegate to current view
	switch a.currentView {
	case LoginViewType:
		updatedView, cmd := a.loginView.Update(msg)
		a.loginView = updatedView.(*LoginView)
		return a, cmd
	}

	return a, nil
}

// View renders the application
func (a *App) View() string {
	switch a.currentView {
	case LoginViewType:
		return a.loginView.View()
	default:
		return "Unknown view"
	}
}
