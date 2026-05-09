package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// OverdueAcknowledgedMsg is sent when user acknowledges overdue alert
type OverdueAcknowledgedMsg struct{}

// OverdueView represents the overdue alert screen
type OverdueView struct {
	hasOverdue bool
}

// NewOverdueView creates a new overdue view
func NewOverdueView() *OverdueView {
	return &OverdueView{
		hasOverdue: false, // Dummy: no overdue books yet
	}
}

// Init initializes the overdue view
func (o *OverdueView) Init() tea.Cmd {
	return nil
}

// Update handles messages for the overdue view
func (o *OverdueView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			// Acknowledge and continue to dashboard
			return o, func() tea.Msg {
				return OverdueAcknowledgedMsg{}
			}
		}
	}

	return o, nil
}

// View renders the overdue alert screen
func (o *OverdueView) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(MatrixGreen).
		Align(lipgloss.Center).
		MarginTop(2).
		MarginBottom(2)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(MatrixGreen).
		Background(MatrixBlack).
		Padding(2, 4).
		Width(70)

	messageStyle := lipgloss.NewStyle().
		Foreground(MatrixDarkGreen).
		Align(lipgloss.Center).
		MarginBottom(2)

	buttonStyle := lipgloss.NewStyle().
		Foreground(MatrixBlack).
		Background(MatrixGreen).
		Padding(0, 3).
		Bold(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(MatrixDarkGreen).
		Align(lipgloss.Center).
		MarginTop(2)

	var content strings.Builder

	if o.hasOverdue {
		// Future: Display actual overdue books
		content.WriteString(titleStyle.Render("⚠️  OVERDUE NOTICE  ⚠️"))
		content.WriteString("\n\n")
		content.WriteString(MatrixError.Copy().Align(lipgloss.Center).Render("You have overdue books! Please return them as soon as possible."))
		content.WriteString("\n\n")
		// Placeholder for overdue book list
		content.WriteString(messageStyle.Render("(Overdue books will be displayed here)"))
	} else {
		// No overdue books
		content.WriteString(titleStyle.Render("✓ All Clear"))
		content.WriteString("\n\n")
		content.WriteString(MatrixSuccess.Copy().Align(lipgloss.Center).Render("You have no overdue books."))
	}

	content.WriteString("\n\n")
	content.WriteString(strings.Repeat(" ", 25))
	content.WriteString(buttonStyle.Render("[ Continue ]"))

	var view strings.Builder
	view.WriteString(lipgloss.NewStyle().Align(lipgloss.Center).Render(
		boxStyle.Render(content.String()),
	))
	view.WriteString("\n")
	view.WriteString(helpStyle.Render("Enter: Continue to Dashboard"))

	return lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(80).
		Render(view.String())
}
