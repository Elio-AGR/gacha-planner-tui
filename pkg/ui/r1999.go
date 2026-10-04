package ui

import (
	"fmt"
	"strings"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/calculator"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// RenderR1999View renders the interactive Reverse: 1999 planner tab.
func RenderR1999View(profile models.R1999Profile, inputs [3]textinput.Model, focusIndex int, width int, height int) string {
	res := calculator.CalculateR1999(profile)

	title := lipgloss.NewStyle().Bold(true).Foreground(ColorTeal).Render("📜 REVERSE: 1999 PLANNER & INPUT FORM")

	pityStatus := "50/50 Next (On Banner)"
	if profile.IsGuaranteed {
		pityStatus = "GUARANTEED (100% Rate Up)"
	}

	var statusBadge string
	if res.CanGuaranteeTarget {
		statusBadge = StatusReadyStyle.Render("✅ GUARANTEED TARGET SECURED")
	} else {
		statusBadge = StatusNeedPullsStyle.Render(fmt.Sprintf("⚡ NEED %d MORE PULLS TO GUARANTEE TARGET", res.RemainingPullsToTarget))
	}

	// Focus indicators for rows
	pointer := func(idx int) string {
		if idx == focusIndex {
			return lipgloss.NewStyle().Foreground(ColorTeal).Bold(true).Render("▶ ")
		}
		return "  "
	}

	// Toggle row view
	toggleText := "[ ⚡ 50/50 NEXT ] (Press Space/g/Enter to toggle)"
	if profile.IsGuaranteed {
		toggleText = "[ ✅ GUARANTEED ] (Press Space/g/Enter to toggle)"
	}
	if focusIndex == 3 {
		toggleText = lipgloss.NewStyle().Foreground(ColorCrust).Background(ColorTeal).Bold(true).Render(toggleText)
	} else {
		toggleText = lipgloss.NewStyle().Foreground(ColorTeal).Bold(true).Render(toggleText)
	}

	formRows := []string{
		fmt.Sprintf("%s%s : %s", pointer(0), LabelStyle.Width(20).Render("Clear Drops"), inputs[0].View()),
		fmt.Sprintf("%s%s : %s", pointer(1), LabelStyle.Width(20).Render("Unilogs"), inputs[1].View()),
		fmt.Sprintf("%s%s : %s", pointer(2), LabelStyle.Width(20).Render("Current Pity (0-70)"), inputs[2].View()),
		fmt.Sprintf("%s%s : %s", pointer(3), LabelStyle.Width(20).Render("Guaranteed Status"), toggleText),
	}

	formBox := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(ColorSurface2).
		Padding(0, 1).
		Render(strings.Join(formRows, "\n"))

	calcRows := []string{
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Unilogs from Drops"), ValTealStyle.Render(fmt.Sprintf("%d Unilogs", res.UnilogsFromDrops))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Total Pulls Available"), ValGreenStyle.Render(fmt.Sprintf("%d Pulls", res.TotalPullsAvailable))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Current Pity State"), ValYellowStyle.Render(fmt.Sprintf("%d / %d (%s)", profile.CurrentPity, calculator.R1999HardPity, pityStatus))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Planner Status"), statusBadge),
		"",
		ValSubtextStyle.Render(fmt.Sprintf("• Soft Pity Threshold (60) : %d pulls remaining", res.SoftPityRemaining)),
		ValSubtextStyle.Render(fmt.Sprintf("• 6★ Hard Pity Threshold    : %d pulls remaining", res.PullsToFirst6Star)),
		ValSubtextStyle.Render(fmt.Sprintf("• Max Pulls for Target      : %d total pulls needed from 0 pity", res.MaxPullsToTarget)),
	}

	infoBox := lipgloss.NewStyle().
		MarginTop(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorSurface2).
		Padding(0, 1).
		Render(ValSubtextStyle.Render("💡 Tip: Use ↑/↓ to navigate fields. Type to edit values. Changes auto-save instantly."))

	return R1999CardStyle.
		Width(width - 6).
		Render(lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			formBox,
			"",
			strings.Join(calcRows, "\n"),
			infoBox,
		))
}
