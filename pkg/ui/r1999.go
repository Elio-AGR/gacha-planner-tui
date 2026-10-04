package ui

import (
	"fmt"
	"strings"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/calculator"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
	"github.com/charmbracelet/lipgloss"
)

// RenderR1999View renders the detailed Reverse: 1999 planner tab.
func RenderR1999View(profile models.R1999Profile, width int, height int) string {
	res := calculator.CalculateR1999(profile)

	title := lipgloss.NewStyle().Bold(true).Foreground(ColorTeal).Render("📜 REVERSE: 1999 DETAILED PLANNER")

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

	rows := []string{
		title,
		"",
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Target Character"), ValTealStyle.Render("Jiu Niangzi / Lucy")),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Clear Drops Savings"), ValTealStyle.Render(fmt.Sprintf("%d Drops (%d Pulls)", profile.ClearDrop, res.UnilogsFromDrops))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Unilog Inventory"), ValTealStyle.Render(fmt.Sprintf("%d Unilogs", profile.Unilog))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Total Pulls Available"), ValGreenStyle.Render(fmt.Sprintf("%d Pulls", res.TotalPullsAvailable))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Current Pity State"), ValYellowStyle.Render(fmt.Sprintf("%d / %d", profile.CurrentPity, calculator.R1999HardPity))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Guarantee Status"), ValPinkStyle.Render(pityStatus)),
		"",
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
		Render(ValSubtextStyle.Render("💡 Tip: Hard pity in Reverse: 1999 is at 70 pulls with soft pity rate increases starting at 50 pulls."))

	return R1999CardStyle.
		Width(width - 6).
		Render(lipgloss.JoinVertical(lipgloss.Left, strings.Join(rows, "\n"), infoBox))
}
