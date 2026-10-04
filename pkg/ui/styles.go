package ui

import "github.com/charmbracelet/lipgloss"

// Catppuccin Mocha Color Palette
var (
	ColorRosewater = lipgloss.Color("#f5e0dc")
	ColorFlamingo  = lipgloss.Color("#f2cdcd")
	ColorPink      = lipgloss.Color("#f5c2e7")
	ColorMauve     = lipgloss.Color("#cba6f7")
	ColorRed       = lipgloss.Color("#f38ba8")
	ColorPeach     = lipgloss.Color("#fab387")
	ColorYellow    = lipgloss.Color("#f9e2af")
	ColorGreen     = lipgloss.Color("#a6e3a1")
	ColorTeal      = lipgloss.Color("#94e2d5")
	ColorCyan      = lipgloss.Color("#89dceb")
	ColorSapphire  = lipgloss.Color("#74c7ec")
	ColorBlue      = lipgloss.Color("#89b4fa")
	ColorLavender  = lipgloss.Color("#b4befe")
	ColorText      = lipgloss.Color("#cdd6f4")
	ColorSubtext1  = lipgloss.Color("#bac2de")
	ColorSubtext0  = lipgloss.Color("#a6adc8")
	ColorOverlay2  = lipgloss.Color("#9399b2")
	ColorSurface2  = lipgloss.Color("#585b70")
	ColorSurface1  = lipgloss.Color("#45475a")
	ColorSurface0  = lipgloss.Color("#313244")
	ColorBase      = lipgloss.Color("#1e1e2e")
	ColorCrust     = lipgloss.Color("#11111b")
)

// Global Reusable UI Styles
var (
	// Header
	HeaderTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorCrust).
				Background(ColorMauve).
				Padding(0, 2)

	HeaderSubTitleStyle = lipgloss.NewStyle().
				Foreground(ColorPink).
				Bold(true).
				Padding(0, 1)

	BadgeStyle = lipgloss.NewStyle().
			Foreground(ColorBase).
			Background(ColorTeal).
			Padding(0, 1).
			Bold(true)

	// Tabs
	ActiveTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorCrust).
			Background(ColorPeach).
			Padding(0, 2).
			MarginRight(1)

	InactiveTabStyle = lipgloss.NewStyle().
			Foreground(ColorSubtext0).
			Background(ColorSurface0).
			Padding(0, 2).
			MarginRight(1)

	// Cards
	HSRCardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPink).
			Padding(1, 2)

	R1999CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorTeal).
			Padding(1, 2)

	DashboardCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorMauve).
				Padding(1, 2)

	// Status Badges
	StatusReadyStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorCrust).
				Background(ColorGreen).
				Padding(0, 1)

	StatusNeedPullsStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorCrust).
				Background(ColorYellow).
				Padding(0, 1)

	// Content Labels & Values
	LabelStyle = lipgloss.NewStyle().
			Foreground(ColorSubtext0)

	ValGreenStyle = lipgloss.NewStyle().
			Foreground(ColorGreen).
			Bold(true)

	ValPinkStyle = lipgloss.NewStyle().
			Foreground(ColorPink).
			Bold(true)

	ValTealStyle = lipgloss.NewStyle().
			Foreground(ColorTeal).
			Bold(true)

	ValYellowStyle = lipgloss.NewStyle().
			Foreground(ColorYellow).
			Bold(true)

	ValSubtextStyle = lipgloss.NewStyle().
			Foreground(ColorSubtext1)
)
