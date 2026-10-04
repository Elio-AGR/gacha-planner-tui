package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/config"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Tab Constants
const (
	TabDashboard = 0
	TabR1999     = 1
	TabHSR       = 2
)

var tabNames = []string{
	"1. Dashboard",
	"2. R1999 Planner",
	"3. HSR Planner",
}

// Model Definition
type model struct {
	activeTab int
	width     int
	height    int
	profile   models.UserProfile
}

func initialModel() model {
	profile, err := config.LoadProfile()
	if err != nil {
		profile = models.NewZeroProfile()
	}

	return model{
		activeTab: TabDashboard,
		width:     90,
		height:    28,
		profile:   profile,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			// Auto-save on exit
			_ = config.SaveProfile(m.profile)
			return m, tea.Quit

		case "tab", "right", "l":
			m.activeTab = (m.activeTab + 1) % len(tabNames)

		case "shift+tab", "left", "h":
			m.activeTab = (m.activeTab - 1 + len(tabNames)) % len(tabNames)

		case "1":
			m.activeTab = TabDashboard
		case "2":
			m.activeTab = TabR1999
		case "3":
			m.activeTab = TabHSR
		}
	}

	return m, nil
}

func (m model) View() string {
	var doc strings.Builder

	// Header
	header := renderHeader(m.width)
	doc.WriteString(header + "\n\n")

	// Navigation Tabs
	tabs := renderTabs(m.activeTab, m.width)
	doc.WriteString(tabs + "\n\n")

	// Content Area
	contentHeight := m.height - 10
	if contentHeight < 8 {
		contentHeight = 8
	}

	var content string
	switch m.activeTab {
	case TabDashboard:
		content = ui.RenderDashboard(m.profile, m.width, contentHeight)
	case TabR1999:
		content = ui.RenderR1999View(m.profile.R1999, m.width, contentHeight)
	case TabHSR:
		content = ui.RenderHSRView(m.profile.HSR, m.width, contentHeight)
	}

	doc.WriteString(content + "\n\n")

	// Footer Help Menu
	footer := renderFooter(m.width)
	doc.WriteString(footer)

	return doc.String()
}

func renderHeader(width int) string {
	title := ui.HeaderTitleStyle.Render("✦ GACHA PULLS & BANNER PLANNER ✦")
	subTitle := ui.HeaderSubTitleStyle.Render("🎲 Catppuccin Mocha Cyberpunk Edition")
	badge := ui.BadgeStyle.Render("v1.0.0")

	leftSide := lipgloss.JoinHorizontal(lipgloss.Center, title, subTitle)
	rightSide := badge

	gap := width - lipgloss.Width(leftSide) - lipgloss.Width(rightSide) - 2
	if gap < 0 {
		gap = 0
	}

	headerLine := lipgloss.JoinHorizontal(
		lipgloss.Center,
		leftSide,
		strings.Repeat(" ", gap),
		rightSide,
	)

	return lipgloss.NewStyle().
		Width(width - 2).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(ui.ColorSurface2).
		Padding(0, 1).
		Render(headerLine)
}

func renderTabs(activeTab int, width int) string {
	var renderedTabs []string

	for i, name := range tabNames {
		var style lipgloss.Style
		if i == activeTab {
			style = ui.ActiveTabStyle
		} else {
			style = ui.InactiveTabStyle
		}
		renderedTabs = append(renderedTabs, style.Render(name))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	return lipgloss.NewStyle().
		Width(width - 2).
		Padding(0, 1).
		Render(row)
}

func renderFooter(width int) string {
	keyStyle := lipgloss.NewStyle().
		Foreground(ui.ColorCrust).
		Background(ui.ColorLavender).
		Padding(0, 1).
		Bold(true)

	descStyle := lipgloss.NewStyle().
		Foreground(ui.ColorSubtext1).
		MarginRight(2)

	keys := []struct {
		key  string
		desc string
	}{
		{"1 / 2 / 3", "Select Tab"},
		{"Tab / Shift+Tab", "Switch Tab"},
		{"q / Ctrl+C", "Quit"},
	}

	var parts []string
	for _, item := range keys {
		k := keyStyle.Render(item.key)
		d := descStyle.Render(" " + item.desc)
		parts = append(parts, lipgloss.JoinHorizontal(lipgloss.Center, k, d))
	}

	helpLine := lipgloss.JoinHorizontal(lipgloss.Center, parts...)

	return lipgloss.NewStyle().
		Width(width - 2).
		Border(lipgloss.NormalBorder(), true, false, false, false).
		BorderForeground(ui.ColorSurface1).
		Padding(0, 1).
		Render(helpLine)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}
}
