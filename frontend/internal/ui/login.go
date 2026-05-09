package ui

import (
	"strings"

	"github.com/bork/frontend/internal/api"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// LoginSuccessMsg is sent when login succeeds
type LoginSuccessMsg struct {
	User *api.UserInfo
}

// LoginErrorMsg is sent when login fails
type LoginErrorMsg struct {
	Error string
}

// LoginView represents the login screen
type LoginView struct {
	username     string
	password     string
	focusedField int // 0 = username, 1 = password, 2 = button
	cursorBlink  bool
	apiClient    *api.Client
	statusMsg    string
	isLoading    bool
}

// NewLoginView creates a new login view
func NewLoginView(apiClient *api.Client) *LoginView {
	return &LoginView{
		username:     "",
		password:     "",
		focusedField: 0,
		cursorBlink:  true,
		apiClient:    apiClient,
		statusMsg:    "",
		isLoading:    false,
	}
}

// Init initializes the login view
func (l *LoginView) Init() tea.Cmd {
	return nil
}

// Update handles messages for the login view
func (l *LoginView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case LoginSuccessMsg:
		l.isLoading = false
		l.statusMsg = "Login successful!"
		return l, nil

	case LoginErrorMsg:
		l.isLoading = false
		l.statusMsg = "Error: " + msg.Error
		return l, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "down":
			l.focusedField = (l.focusedField + 1) % 3
			return l, nil

		case "shift+tab", "up":
			l.focusedField = (l.focusedField - 1 + 3) % 3
			return l, nil

		case "enter":
			if l.focusedField == 2 {
				// Login button pressed - authenticate
				if l.username == "" || l.password == "" {
					l.statusMsg = "Please enter username and password"
					return l, nil
				}
				l.isLoading = true
				l.statusMsg = "Logging in..."
				return l, l.performLogin()
			}
			// Move to next field
			l.focusedField = (l.focusedField + 1) % 3
			return l, nil

		case "backspace":
			if l.focusedField == 0 && len(l.username) > 0 {
				l.username = l.username[:len(l.username)-1]
			} else if l.focusedField == 1 && len(l.password) > 0 {
				l.password = l.password[:len(l.password)-1]
			}
			return l, nil

		default:
			// Handle text input
			if len(msg.String()) == 1 {
				if l.focusedField == 0 {
					l.username += msg.String()
				} else if l.focusedField == 1 {
					l.password += msg.String()
				}
			}
			return l, nil
		}
	}

	return l, nil
}

// View renders the login screen
func (l *LoginView) View() string {
	// Styles
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00BFFF")).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(2)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00BFFF")).
		Padding(1, 2).
		Width(60)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true)

	inputStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#1a1a1a")).
		Padding(0, 1).
		Width(40)

	focusedInputStyle := inputStyle.Copy().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#00BFFF"))

	buttonStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#00BFFF")).
		Padding(0, 3).
		Bold(true)

	focusedButtonStyle := buttonStyle.Copy().
		Background(lipgloss.Color("#0080FF"))

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		Align(lipgloss.Center).
		MarginTop(2)

	// ASCII Art Logo
	logo := `
    ██████╗  ██████╗ ██████╗ ██╗  ██╗
    ██╔══██╗██╔═══██╗██╔══██╗██║ ██╔╝
    ██████╔╝██║   ██║██████╔╝█████╔╝
    ██╔══██╗██║   ██║██╔══██╗██╔═██╗
    ██████╔╝╚██████╔╝██║  ██║██║  ██╗
    ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝
    `

	subtitle := "Book Organization & Rental Kiosk v0.1.0"

	// Build form
	var formBuilder strings.Builder

	// Username field
	formBuilder.WriteString(labelStyle.Render("Username:"))
	formBuilder.WriteString("\n")
	usernameDisplay := l.username
	if l.focusedField == 0 && l.cursorBlink {
		usernameDisplay += "█"
	}
	if len(usernameDisplay) == 0 {
		usernameDisplay = " "
	}
	if l.focusedField == 0 {
		formBuilder.WriteString(focusedInputStyle.Render(usernameDisplay))
	} else {
		formBuilder.WriteString(inputStyle.Render(usernameDisplay))
	}
	formBuilder.WriteString("\n\n")

	// Password field
	formBuilder.WriteString(labelStyle.Render("Password:"))
	formBuilder.WriteString("\n")
	passwordDisplay := strings.Repeat("*", len(l.password))
	if l.focusedField == 1 && l.cursorBlink {
		passwordDisplay += "█"
	}
	if len(passwordDisplay) == 0 {
		passwordDisplay = " "
	}
	if l.focusedField == 1 {
		formBuilder.WriteString(focusedInputStyle.Render(passwordDisplay))
	} else {
		formBuilder.WriteString(inputStyle.Render(passwordDisplay))
	}
	formBuilder.WriteString("\n\n")

	// Login button
	formBuilder.WriteString(strings.Repeat(" ", 15))
	if l.focusedField == 2 {
		formBuilder.WriteString(focusedButtonStyle.Render("[ Login ]"))
	} else {
		formBuilder.WriteString(buttonStyle.Render("[ Login ]"))
	}

	// Assemble the view
	var view strings.Builder
	view.WriteString(titleStyle.Render(logo))
	view.WriteString("\n")
	view.WriteString(titleStyle.Render(subtitle))
	view.WriteString("\n\n")
	view.WriteString(lipgloss.NewStyle().Align(lipgloss.Center).Render(
		boxStyle.Render(formBuilder.String()),
	))
	view.WriteString("\n")

	// Status message
	if l.statusMsg != "" {
		statusStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Align(lipgloss.Center).
			MarginTop(1)
		if strings.Contains(l.statusMsg, "successful") {
			statusStyle = statusStyle.Foreground(lipgloss.Color("#4CAF50"))
		}
		view.WriteString(statusStyle.Render(l.statusMsg))
		view.WriteString("\n")
	}

	view.WriteString(helpStyle.Render("Tab: Next Field | Enter: Submit | Esc: Exit"))

	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(80).
		Render(view.String())
}

// performLogin performs the login API call
func (l *LoginView) performLogin() tea.Cmd {
	return func() tea.Msg {
		resp, err := l.apiClient.Login(l.username, l.password)
		if err != nil {
			return LoginErrorMsg{Error: err.Error()}
		}
		return LoginSuccessMsg{User: &resp.User}
	}
}
