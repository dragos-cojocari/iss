package ui

import (
	"os"

	"github.com/bork/frontend/internal/api"
	tea "github.com/charmbracelet/bubbletea"
)

// App represents the main application model
type App struct {
	currentView     ViewType
	apiClient       *api.Client
	loginView       *LoginView
	overdueView     *OverdueView
	dashboardView   *DashboardView
	browseBooksView *BrowseBooksView
	currentUser     *api.UserInfo
	width           int
	height          int
}

// ViewType represents different screens in the application
type ViewType int

const (
	LoginViewType ViewType = iota
	OverdueViewType
	DashboardViewType
	BrowseBooksViewType
)

// NewApp creates a new application instance
func NewApp() *App {
	// Get backend URL from environment or use default
	backendURL := os.Getenv("BORK_BACKEND_URL")
	if backendURL == "" {
		backendURL = "http://localhost:8080"
	}

	apiClient := api.NewClient(backendURL)

	return &App{
		currentView: LoginViewType,
		apiClient:   apiClient,
		loginView:   NewLoginView(apiClient),
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
		case "ctrl+c":
			return a, tea.Quit
		case "esc":
			// Esc behavior depends on current view
			switch a.currentView {
			case LoginViewType, DashboardViewType:
				// Quit from login or main dashboard
				return a, tea.Quit
			case OverdueViewType:
				// Go to dashboard (skip overdue)
				a.dashboardView = NewDashboardView(a.apiClient, a.currentUser)
				a.currentView = DashboardViewType
				return a, a.dashboardView.Init()
			case BrowseBooksViewType:
				// Go back to dashboard and cycle quote
				a.currentView = DashboardViewType
				return a, a.dashboardView.Init()
			}
		}

	case LoginSuccessMsg:
		// Transition from login to overdue alert
		a.currentUser = msg.User
		a.overdueView = NewOverdueView()
		a.currentView = OverdueViewType
		return a, a.overdueView.Init()

	case OverdueAcknowledgedMsg:
		// Transition from overdue to dashboard
		a.dashboardView = NewDashboardView(a.apiClient, a.currentUser)
		a.currentView = DashboardViewType
		return a, a.dashboardView.Init()

	case BrowseBooksMsg:
		// Transition from dashboard to browse books
		a.browseBooksView = NewBrowseBooksView(a.apiClient, a.currentUser)
		a.currentView = BrowseBooksViewType
		return a, a.browseBooksView.Init()

	case BackToDashboardMsg:
		// Transition back to dashboard from browse books
		a.currentView = DashboardViewType
		// Cycle to next quote when returning to dashboard
		return a, a.dashboardView.Init()

	case LogoutMsg:
		// Transition back to login
		a.currentUser = nil
		a.loginView = NewLoginView(a.apiClient)
		a.dashboardView = nil
		a.overdueView = nil
		a.browseBooksView = nil
		a.currentView = LoginViewType
		return a, a.loginView.Init()
	}

	// Delegate to current view
	var cmd tea.Cmd
	switch a.currentView {
	case LoginViewType:
		var updatedView tea.Model
		updatedView, cmd = a.loginView.Update(msg)
		a.loginView = updatedView.(*LoginView)

	case OverdueViewType:
		var updatedView tea.Model
		updatedView, cmd = a.overdueView.Update(msg)
		a.overdueView = updatedView.(*OverdueView)

	case DashboardViewType:
		var updatedView tea.Model
		updatedView, cmd = a.dashboardView.Update(msg)
		a.dashboardView = updatedView.(*DashboardView)

	case BrowseBooksViewType:
		var updatedView tea.Model
		updatedView, cmd = a.browseBooksView.Update(msg)
		a.browseBooksView = updatedView.(*BrowseBooksView)
	}

	return a, cmd
}

// View renders the application
func (a *App) View() string {
	switch a.currentView {
	case LoginViewType:
		return a.loginView.View()
	case OverdueViewType:
		return a.overdueView.View()
	case DashboardViewType:
		return a.dashboardView.View()
	case BrowseBooksViewType:
		return a.browseBooksView.View()
	default:
		return "Unknown view"
	}
}
