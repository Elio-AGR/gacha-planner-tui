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
	"github.com/charmbracelet/bubbles/viewport"
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

	viewport viewport.Model
	ready    bool

	// Form inputs for HSR (0: Jades, 1: Passes, 2: Pity, 3: Banner Days, 4: Daily Income)
	// Focus indices (0..6): 0:Jades, 1:Passes, 2:Pity, 3:Toggle Guaranteed, 4:Banner Days, 5:Daily Income, 6:Preset Helper
	hsrFocusIndex int
	hsrInputs     [5]textinput.Model

	// Form inputs for R1999 (0: Drops, 1: Unilogs, 2: Pity, 3: Banner Days, 4: Daily Income)
	// Focus indices (0..6): 0:Drops, 1:Unilogs, 2:Pity, 3:Toggle Guaranteed, 4:Banner Days, 5:Daily Income, 6:Preset Helper
	r1999FocusIndex int
	r1999Inputs     [5]textinput.Model
}

func createNumInput(placeholder string, val int) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.CharLimit = 7
	ti.Width = 8
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
		height:          24,
		profile:         profile,
		hsrFocusIndex:   0,
		r1999FocusIndex: 0,
	}

	// Initialize HSR Inputs
	m.hsrInputs[0] = createNumInput("19200", profile.HSR.StellarJade)
	m.hsrInputs[1] = createNumInput("15", profile.HSR.SpecialPass)
	m.hsrInputs[2] = createNumInput("65", profile.HSR.CurrentPity)
	m.hsrInputs[3] = createNumInput("14", profile.HSR.BannerDays)
	m.hsrInputs[4] = createNumInput("90", profile.HSR.DailyIncome)

	// Initialize R1999 Inputs
	m.r1999Inputs[0] = createNumInput("14400", profile.R1999.ClearDrop)
	m.r1999Inputs[1] = createNumInput("68", profile.R1999.Unilog)
	m.r1999Inputs[2] = createNumInput("42", profile.R1999.CurrentPity)
	m.r1999Inputs[3] = createNumInput("14", profile.R1999.BannerDays)
	m.r1999Inputs[4] = createNumInput("80", profile.R1999.DailyIncome)

	m.focusCurrentInput()
	return m
}

func (m model) isInputFocused() bool {
	if m.activeTab == TabHSR {
		switch m.hsrFocusIndex {
		case 0, 1, 2, 4, 5:
			return true
		}
	}
	if m.activeTab == TabR1999 {
		switch m.r1999FocusIndex {
		case 0, 1, 2, 4, 5:
			return true
		}
	}
	return false
}

func (m *model) focusCurrentInput() {
	// Blur all HSR inputs
	for i := range m.hsrInputs {
		m.hsrInputs[i].Blur()
	}
	if m.activeTab == TabHSR {
		switch m.hsrFocusIndex {
		case 0:
			m.hsrInputs[0].Focus()
		case 1:
			m.hsrInputs[1].Focus()
		case 2:
			m.hsrInputs[2].Focus()
		case 4:
			m.hsrInputs[3].Focus()
		case 5:
			m.hsrInputs[4].Focus()
		}
	}

	// Blur all R1999 inputs
	for i := range m.r1999Inputs {
		m.r1999Inputs[i].Blur()
	}
	if m.activeTab == TabR1999 {
		switch m.r1999FocusIndex {
		case 0:
			m.r1999Inputs[0].Focus()
		case 1:
			m.r1999Inputs[1].Focus()
		case 2:
			m.r1999Inputs[2].Focus()
		case 4:
			m.r1999Inputs[3].Focus()
		case 5:
			m.r1999Inputs[4].Focus()
		}
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

	valStr = m.hsrInputs[3].Value()
	if valStr == "" {
		m.profile.HSR.BannerDays = 0
	} else if days, err := strconv.Atoi(valStr); err == nil && days >= 0 {
		m.profile.HSR.BannerDays = days
	}

	valStr = m.hsrInputs[4].Value()
	if valStr == "" {
		m.profile.HSR.DailyIncome = 0
	} else if inc, err := strconv.Atoi(valStr); err == nil && inc >= 0 {
		m.profile.HSR.DailyIncome = inc
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

	valStr = m.r1999Inputs[3].Value()
	if valStr == "" {
		m.profile.R1999.BannerDays = 0
	} else if days, err := strconv.Atoi(valStr); err == nil && days >= 0 {
		m.profile.R1999.BannerDays = days
	}

	valStr = m.r1999Inputs[4].Value()
	if valStr == "" {
		m.profile.R1999.DailyIncome = 0
	} else if inc, err := strconv.Atoi(valStr); err == nil && inc >= 0 {
		m.profile.R1999.DailyIncome = inc
	}

	_ = config.SaveProfile(m.profile)
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) updateViewportContent() string {
	var rawContent string
	vpHeight := m.viewport.Height
	if vpHeight < 5 {
		vpHeight = 5
	}

	switch m.activeTab {
	case TabDashboard:
		rawContent = ui.RenderDashboard(m.profile, m.width, vpHeight)
	case TabR1999:
		rawContent = ui.RenderR1999View(m.profile.R1999, m.r1999Inputs, m.r1999FocusIndex, m.width, vpHeight)
	case TabHSR:
		rawContent = ui.RenderHSRView(m.profile.HSR, m.hsrInputs, m.hsrFocusIndex, m.width, vpHeight)
	}

	return rawContent
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := lipgloss.Height(renderHeader(msg.Width)) + lipgloss.Height(renderTabs(m.activeTab, msg.Width)) + 2
		footerHeight := lipgloss.Height(renderFooter(msg.Width)) + 1
		vpHeight := msg.Height - headerHeight - footerHeight
		if vpHeight < 5 {
			vpHeight = 5
		}

		if !m.ready {
			m.viewport = viewport.New(msg.Width, vpHeight)
			m.viewport.SetContent(m.updateViewportContent())
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = vpHeight
			m.viewport.SetContent(m.updateViewportContent())
		}

	case tea.KeyMsg:
		keyStr := msg.String()
		isFocused := m.isInputFocused()

		// 1. Emergency Exit (Ctrl+C, Esc)
		if keyStr == "ctrl+c" || keyStr == "esc" {
			_ = config.SaveProfile(m.profile)
			return m, tea.Quit
		}

		// 2. Quit shortcut 'q' active when NOT focused in text input
		if keyStr == "q" && !isFocused {
			_ = config.SaveProfile(m.profile)
			return m, tea.Quit
		}

		// 3. Viewport Scrolling via PgUp / PgDn / Shift+Up / Shift+Down
		switch keyStr {
		case "pgup", "shift+up":
			m.viewport.HalfViewUp()
			return m, nil
		case "pgdown", "shift+down":
			m.viewport.HalfViewDown()
			return m, nil
		}

		// 4. Tab switching via Tab / Shift+Tab (Always handled top-level)
		if keyStr == "tab" {
			m.activeTab = (m.activeTab + 1) % len(tabNames)
			m.focusCurrentInput()
			if m.ready {
				m.viewport.SetContent(m.updateViewportContent())
			}
			return m, nil
		}
		if keyStr == "shift+tab" {
			m.activeTab = (m.activeTab - 1 + len(tabNames)) % len(tabNames)
			m.focusCurrentInput()
			if m.ready {
				m.viewport.SetContent(m.updateViewportContent())
			}
			return m, nil
		}

		// 5. Direct numeric tab selection (1, 2, 3) active when NOT focused in text input
		if !isFocused {
			switch keyStr {
			case "1":
				m.activeTab = TabDashboard
				m.focusCurrentInput()
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil
			case "2":
				m.activeTab = TabR1999
				m.focusCurrentInput()
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil
			case "3":
				m.activeTab = TabHSR
				m.focusCurrentInput()
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil
			}
		}

		// 6. Up & Down Form Field Navigation, Preset Cycling & Toggle Handling
		if m.activeTab == TabHSR {
			switch keyStr {
			case "up":
				m.hsrFocusIndex = (m.hsrFocusIndex - 1 + 7) % 7
				m.focusCurrentInput()
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil
			case "down":
				m.hsrFocusIndex = (m.hsrFocusIndex + 1) % 7
				m.focusCurrentInput()
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil

			case "left", "h":
				if m.hsrFocusIndex == 6 {
					numPresets := len(models.HSRBannerPresets)
					m.profile.HSR.TargetBannerIndex = (m.profile.HSR.TargetBannerIndex - 1 + numPresets) % numPresets
					preset := models.HSRBannerPresets[m.profile.HSR.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.HSR.BannerDays = preset.EstimatedDaysLeft
						m.hsrInputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				}
			case "right", "l":
				if m.hsrFocusIndex == 6 {
					numPresets := len(models.HSRBannerPresets)
					m.profile.HSR.TargetBannerIndex = (m.profile.HSR.TargetBannerIndex + 1) % numPresets
					preset := models.HSRBannerPresets[m.profile.HSR.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.HSR.BannerDays = preset.EstimatedDaysLeft
						m.hsrInputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				}

			case "g", " ":
				if m.hsrFocusIndex == 3 {
					m.profile.HSR.IsGuaranteed = !m.profile.HSR.IsGuaranteed
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				} else if m.hsrFocusIndex == 6 {
					numPresets := len(models.HSRBannerPresets)
					m.profile.HSR.TargetBannerIndex = (m.profile.HSR.TargetBannerIndex + 1) % numPresets
					preset := models.HSRBannerPresets[m.profile.HSR.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.HSR.BannerDays = preset.EstimatedDaysLeft
						m.hsrInputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				}

			case "enter":
				if m.hsrFocusIndex == 3 {
					m.profile.HSR.IsGuaranteed = !m.profile.HSR.IsGuaranteed
					_ = config.SaveProfile(m.profile)
				} else if m.hsrFocusIndex == 6 {
					numPresets := len(models.HSRBannerPresets)
					m.profile.HSR.TargetBannerIndex = (m.profile.HSR.TargetBannerIndex + 1) % numPresets
					preset := models.HSRBannerPresets[m.profile.HSR.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.HSR.BannerDays = preset.EstimatedDaysLeft
						m.hsrInputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
				} else {
					m.hsrFocusIndex = (m.hsrFocusIndex + 1) % 7
					m.focusCurrentInput()
				}
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil
			}

			// Forward character inputs to focused textinput with Numeric-Only filter
			if m.hsrFocusIndex == 0 || m.hsrFocusIndex == 1 || m.hsrFocusIndex == 2 || m.hsrFocusIndex == 4 || m.hsrFocusIndex == 5 {
				if len(keyStr) == 1 {
					r := rune(keyStr[0])
					if !unicode.IsDigit(r) {
						return m, nil
					}
				}

				inputIdx := m.hsrFocusIndex
				if m.hsrFocusIndex == 4 {
					inputIdx = 3
				} else if m.hsrFocusIndex == 5 {
					inputIdx = 4
				}

				var cmd tea.Cmd
				m.hsrInputs[inputIdx], cmd = m.hsrInputs[inputIdx].Update(msg)
				cmds = append(cmds, cmd)
				m.syncHSRProfile()
			}
		} else if m.activeTab == TabR1999 {
			switch keyStr {
			case "up":
				m.r1999FocusIndex = (m.r1999FocusIndex - 1 + 7) % 7
				m.focusCurrentInput()
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil
			case "down":
				m.r1999FocusIndex = (m.r1999FocusIndex + 1) % 7
				m.focusCurrentInput()
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil

			case "left", "h":
				if m.r1999FocusIndex == 6 {
					numPresets := len(models.R1999BannerPresets)
					m.profile.R1999.TargetBannerIndex = (m.profile.R1999.TargetBannerIndex - 1 + numPresets) % numPresets
					preset := models.R1999BannerPresets[m.profile.R1999.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.R1999.BannerDays = preset.EstimatedDaysLeft
						m.r1999Inputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				}
			case "right", "l":
				if m.r1999FocusIndex == 6 {
					numPresets := len(models.R1999BannerPresets)
					m.profile.R1999.TargetBannerIndex = (m.profile.R1999.TargetBannerIndex + 1) % numPresets
					preset := models.R1999BannerPresets[m.profile.R1999.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.R1999.BannerDays = preset.EstimatedDaysLeft
						m.r1999Inputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				}

			case "g", " ":
				if m.r1999FocusIndex == 3 {
					m.profile.R1999.IsGuaranteed = !m.profile.R1999.IsGuaranteed
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				} else if m.r1999FocusIndex == 6 {
					numPresets := len(models.R1999BannerPresets)
					m.profile.R1999.TargetBannerIndex = (m.profile.R1999.TargetBannerIndex + 1) % numPresets
					preset := models.R1999BannerPresets[m.profile.R1999.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.R1999.BannerDays = preset.EstimatedDaysLeft
						m.r1999Inputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
					if m.ready {
						m.viewport.SetContent(m.updateViewportContent())
					}
					return m, nil
				}

			case "enter":
				if m.r1999FocusIndex == 3 {
					m.profile.R1999.IsGuaranteed = !m.profile.R1999.IsGuaranteed
					_ = config.SaveProfile(m.profile)
				} else if m.r1999FocusIndex == 6 {
					numPresets := len(models.R1999BannerPresets)
					m.profile.R1999.TargetBannerIndex = (m.profile.R1999.TargetBannerIndex + 1) % numPresets
					preset := models.R1999BannerPresets[m.profile.R1999.TargetBannerIndex]
					if !preset.IsCustom {
						m.profile.R1999.BannerDays = preset.EstimatedDaysLeft
						m.r1999Inputs[3].SetValue(fmt.Sprintf("%d", preset.EstimatedDaysLeft))
					}
					_ = config.SaveProfile(m.profile)
				} else {
					m.r1999FocusIndex = (m.r1999FocusIndex + 1) % 7
					m.focusCurrentInput()
				}
				if m.ready {
					m.viewport.SetContent(m.updateViewportContent())
				}
				return m, nil
			}

			// Forward character inputs to focused textinput with Numeric-Only filter
			if m.r1999FocusIndex == 0 || m.r1999FocusIndex == 1 || m.r1999FocusIndex == 2 || m.r1999FocusIndex == 4 || m.r1999FocusIndex == 5 {
				if len(keyStr) == 1 {
					r := rune(keyStr[0])
					if !unicode.IsDigit(r) {
						return m, nil
					}
				}

				inputIdx := m.r1999FocusIndex
				if m.r1999FocusIndex == 4 {
					inputIdx = 3
				} else if m.r1999FocusIndex == 5 {
					inputIdx = 4
				}

				var cmd tea.Cmd
				m.r1999Inputs[inputIdx], cmd = m.r1999Inputs[inputIdx].Update(msg)
				cmds = append(cmds, cmd)
				m.syncR1999Profile()
			}
		}

		// Also update viewport for scrolling
		var vpCmd tea.Cmd
		m.viewport, vpCmd = m.viewport.Update(msg)
		cmds = append(cmds, vpCmd)
	}

	if m.ready {
		m.viewport.SetContent(m.updateViewportContent())
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var doc strings.Builder

	// Sticky Header
	header := renderHeader(m.width)
	doc.WriteString(header + "\n")

	// Sticky Navigation Tabs
	tabs := renderTabs(m.activeTab, m.width)
	doc.WriteString(tabs + "\n")

	// Scrollable Middle Content Viewport
	if m.ready {
		doc.WriteString(m.viewport.View() + "\n")
	} else {
		doc.WriteString("Loading view...\n")
	}

	// Sticky Footer
	footer := renderFooter(m.width)
	doc.WriteString(footer)

	return doc.String()
}

func renderHeader(width int) string {
	title := ui.HeaderTitleStyle.Render("✦ GACHA PULLS & BANNER PLANNER ✦")
	badge := ui.BadgeStyle.Render("v1.0.0")

	leftSide := title
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
		{"← / → / Space", "Change Preset"},
		{"0-9", "Type Numbers"},
		{"Space / g", "Toggle 50/50"},
		{"Esc / Ctrl+C / q*", "Quit"},
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
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}
}
