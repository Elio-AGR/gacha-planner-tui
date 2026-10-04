package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/config"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/ui"
	"github.com/charmbracelet/bubbles/textinput"
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

	// Form inputs for HSR
	hsrFocusIndex int
	hsrInputs     [3]textinput.Model // 0: Jades, 1: Passes, 2: Pity

	// Form inputs for R1999
	r1999FocusIndex int
	r1999Inputs     [3]textinput.Model // 0: Drops, 1: Unilogs, 2: Pity
}

func createNumInput(placeholder string, val int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.CharLimit = 7
	ti.Width = 10
	ti.SetValue(fmt.Sprintf("%d", val))
	ti.Prompt = " "
	return ti
}

func initialModel() model {
	profile, err := config.LoadProfile()
	if err != nil {
		profile = models.NewZeroProfile()
	}

	m := model{
		activeTab:       TabDashboard,
		width:           90,
		height:          30,
		profile:         profile,
		hsrFocusIndex:   0,
		r1999FocusIndex: 0,
	}

	// Initialize HSR Inputs
	m.hsrInputs[0] = createNumInput("19200", profile.HSR.StellarJade)
	m.hsrInputs[1] = createNumInput("15", profile.HSR.SpecialPass)
	m.hsrInputs[2] = createNumInput("65", profile.HSR.CurrentPity)

	// Initialize R1999 Inputs
	m.r1999Inputs[0] = createNumInput("14400", profile.R1999.ClearDrop)
	m.r1999Inputs[1] = createNumInput("68", profile.R1999.Unilog)
	m.r1999Inputs[2] = createNumInput("42", profile.R1999.CurrentPity)

	m.focusCurrentInput()
	return m
}

func (m model) isInputFocused() bool {
	if m.activeTab == TabHSR && m.hsrFocusIndex >= 0 && m.hsrFocusIndex < 3 {
		return true
	}
	if m.activeTab == TabR1999 && m.r1999FocusIndex >= 0 && m.r1999FocusIndex < 3 {
		return true
	}
	return false
}

func (m *model) focusCurrentInput() {
	// Blur all HSR inputs
	for i := range m.hsrInputs {
		m.hsrInputs[i].Blur()
	}
	if m.activeTab == TabHSR && m.hsrFocusIndex >= 0 && m.hsrFocusIndex < 3 {
		m.hsrInputs[m.hsrFocusIndex].Focus()
	}

	// Blur all R1999 inputs
	for i := range m.r1999Inputs {
		m.r1999Inputs[i].Blur()
	}
	if m.activeTab == TabR1999 && m.r1999FocusIndex >= 0 && m.r1999FocusIndex < 3 {
		m.r1999Inputs[m.r1999FocusIndex].Focus()
	}
}

func (m *model) syncHSRProfile() {
	valStr := m.hsrInputs[0].Value()
	if valStr == "" {
		m.profile.HSR.StellarJade = 0
	} else if jades, err := strconv.Atoi(valStr); err == nil && jades >= 0 {
		m.profile.HSR.StellarJade = jades
	}

	valStr = m.hsrInputs[1].Value()
	if valStr == "" {
		m.profile.HSR.SpecialPass = 0
	} else if passes, err := strconv.Atoi(valStr); err == nil && passes >= 0 {
		m.profile.HSR.SpecialPass = passes
	}

	valStr = m.hsrInputs[2].Value()
	if valStr == "" {
		m.profile.HSR.CurrentPity = 0
	} else if pity, err := strconv.Atoi(valStr); err == nil && pity >= 0 {
		if pity > 90 {
			pity = 90
		}
		m.profile.HSR.CurrentPity = pity
	}

	_ = config.SaveProfile(m.profile)
}

func (m *model) syncR1999Profile() {
	valStr := m.r1999Inputs[0].Value()
	if valStr == "" {
		m.profile.R1999.ClearDrop = 0
	} else if drops, err := strconv.Atoi(valStr); err == nil && drops >= 0 {
		m.profile.R1999.ClearDrop = drops
	}

	valStr = m.r1999Inputs[1].Value()
	if valStr == "" {
		m.profile.R1999.Unilog = 0
	} else if unilogs, err := strconv.Atoi(valStr); err == nil && unilogs >= 0 {
		m.profile.R1999.Unilog = unilogs
	}

	valStr = m.r1999Inputs[2].Value()
	if valStr == "" {
		m.profile.R1999.CurrentPity = 0
	} else if pity, err := strconv.Atoi(valStr); err == nil && pity >= 0 {
		if pity > 70 {
			pity = 70
		}
		m.profile.R1999.CurrentPity = pity
	}

	_ = config.SaveProfile(m.profile)
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		keyStr := msg.String()
		isFocused := m.isInputFocused()

		// Emergency exit always available
		if keyStr == "ctrl+c" {
			_ = config.SaveProfile(m.profile)
			return m, tea.Quit
		}

		// Quit shortcut 'q' only active when NOT focused inside a numeric text input
		if keyStr == "q" && !isFocused {
			_ = config.SaveProfile(m.profile)
			return m, tea.Quit
		}

		// Tab switching via Tab / Shift+Tab always available
		if keyStr == "tab" {
			m.activeTab = (m.activeTab + 1) % len(tabNames)
			m.focusCurrentInput()
			return m, nil
		}
		if keyStr == "shift+tab" {
			m.activeTab = (m.activeTab - 1 + len(tabNames)) % len(tabNames)
			m.focusCurrentInput()
			return m, nil
		}

		// Direct numeric tab selection (1, 2, 3) only active when NOT focused in text input
		if !isFocused {
			switch keyStr {
			case "1":
				m.activeTab = TabDashboard
				m.focusCurrentInput()
				return m, nil
			case "2":
				m.activeTab = TabR1999
				m.focusCurrentInput()
				return m, nil
			case "3":
				m.activeTab = TabHSR
				m.focusCurrentInput()
				return m, nil
			}
		}

		// Handle Form Navigation and Text Editing per Tab
		if m.activeTab == TabHSR {
			switch keyStr {
			case "up":
				m.hsrFocusIndex = (m.hsrFocusIndex - 1 + 4) % 4
				m.focusCurrentInput()
				return m, nil
			case "down":
				m.hsrFocusIndex = (m.hsrFocusIndex + 1) % 4
				m.focusCurrentInput()
				return m, nil
			case "g", " ":
				if m.hsrFocusIndex == 3 {
					m.profile.HSR.IsGuaranteed = !m.profile.HSR.IsGuaranteed
					_ = config.SaveProfile(m.profile)
					return m, nil
				}
			case "enter":
				if m.hsrFocusIndex == 3 {
					m.profile.HSR.IsGuaranteed = !m.profile.HSR.IsGuaranteed
					_ = config.SaveProfile(m.profile)
				} else {
					m.hsrFocusIndex = (m.hsrFocusIndex + 1) % 4
					m.focusCurrentInput()
				}
				return m, nil
			}

			// Forward to focused textinput with Numeric-Only filter (0-9)
			if m.hsrFocusIndex < 3 {
				if len(keyStr) == 1 {
					r := rune(keyStr[0])
					if !unicode.IsDigit(r) {
						return m, nil // Ignore non-numeric character
					}
				}

				var cmd tea.Cmd
				m.hsrInputs[m.hsrFocusIndex], cmd = m.hsrInputs[m.hsrFocusIndex].Update(msg)
				cmds = append(cmds, cmd)
				m.syncHSRProfile()
			}
		} else if m.activeTab == TabR1999 {
			switch keyStr {
			case "up":
				m.r1999FocusIndex = (m.r1999FocusIndex - 1 + 4) % 4
				m.focusCurrentInput()
				return m, nil
			case "down":
				m.r1999FocusIndex = (m.r1999FocusIndex + 1) % 4
				m.focusCurrentInput()
				return m, nil
			case "g", " ":
				if m.r1999FocusIndex == 3 {
					m.profile.R1999.IsGuaranteed = !m.profile.R1999.IsGuaranteed
					_ = config.SaveProfile(m.profile)
					return m, nil
				}
			case "enter":
				if m.r1999FocusIndex == 3 {
					m.profile.R1999.IsGuaranteed = !m.profile.R1999.IsGuaranteed
					_ = config.SaveProfile(m.profile)
				} else {
					m.r1999FocusIndex = (m.r1999FocusIndex + 1) % 4
					m.focusCurrentInput()
				}
				return m, nil
			}

			// Forward to focused textinput with Numeric-Only filter (0-9)
			if m.r1999FocusIndex < 3 {
				if len(keyStr) == 1 {
					r := rune(keyStr[0])
					if !unicode.IsDigit(r) {
						return m, nil // Ignore non-numeric character
					}
				}

				var cmd tea.Cmd
				m.r1999Inputs[m.r1999FocusIndex], cmd = m.r1999Inputs[m.r1999FocusIndex].Update(msg)
				cmds = append(cmds, cmd)
				m.syncR1999Profile()
			}
		}
	}

	return m, tea.Batch(cmds...)
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
		content = ui.RenderR1999View(m.profile.R1999, m.r1999Inputs, m.r1999FocusIndex, m.width, contentHeight)
	case TabHSR:
		content = ui.RenderHSRView(m.profile.HSR, m.hsrInputs, m.hsrFocusIndex, m.width, contentHeight)
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
		{"Tab / Shift+Tab", "Switch Tab"},
		{"↑ / ↓ / Enter", "Navigate Fields"},
		{"0-9", "Type Numbers"},
		{"Space / g", "Toggle 50/50"},
		{"Ctrl+C / q*", "Quit"},
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
