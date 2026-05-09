package ui

import "github.com/charmbracelet/lipgloss"

// Matrix color palette inspired by The Matrix
var (
	MatrixGreen     = lipgloss.Color("#00FF41") // Bright matrix green
	MatrixDarkGreen = lipgloss.Color("#008F11") // Darker green
	MatrixBlack     = lipgloss.Color("#0D0208") // Almost black background
	MatrixGray      = lipgloss.Color("#003B00") // Dark green-gray
	MatrixHighlight = lipgloss.Color("#39FF14") // Neon green highlight
	MatrixRed       = lipgloss.Color("#FF0000") // Error red
)

// Matrix-themed styles
var (
	// Title style for headers
	MatrixTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(MatrixGreen).
			Background(MatrixBlack).
			Padding(0, 1)

	// Regular text
	MatrixText = lipgloss.NewStyle().
			Foreground(MatrixDarkGreen)

	// Highlighted/important text
	MatrixHighlightText = lipgloss.NewStyle().
				Foreground(MatrixHighlight).
				Bold(true)

	// Box/container style
	MatrixBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(MatrixGreen).
			Padding(1, 2).
			Background(MatrixBlack)

	// Selected item style (inverted)
	MatrixSelected = lipgloss.NewStyle().
			Foreground(MatrixBlack).
			Background(MatrixGreen).
			Bold(true).
			Padding(0, 1)

	// Input field style
	MatrixInput = lipgloss.NewStyle().
			Foreground(MatrixGreen).
			Background(MatrixGray).
			Padding(0, 1)

	// Error message style
	MatrixError = lipgloss.NewStyle().
			Foreground(MatrixRed).
			Bold(true)

	// Success message style
	MatrixSuccess = lipgloss.NewStyle().
			Foreground(MatrixHighlight).
			Bold(true)

	// Help text style
	MatrixHelp = lipgloss.NewStyle().
			Foreground(MatrixDarkGreen).
			Italic(true)

	// Header bar style
	MatrixHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(MatrixGreen).
			Background(MatrixBlack).
			Padding(0, 2)
)
