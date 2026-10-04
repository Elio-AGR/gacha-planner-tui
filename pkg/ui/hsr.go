package ui

import (
	"fmt"
	"strings"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/calculator"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
	"github.com/charmbracelet/lipgloss"
)

// RenderHSRView renders the detailed Honkai: Star Rail planner tab.
func RenderHSRView(profile models.HSRProfile, width int, height int) string {
	res := calculator.CalculateHSR(profile)

	title := lipgloss.NewStyle().Bold(true).Foreground(ColorPink).Render("🌌 HONKAI: STAR RAIL DETAILED PLANNER")

	pityStatus := "50/50 Next (On Banner)"
	if profile.IsGuaranteed {
		pityStatus = "GUARANTEED (Lost Previous 50/50)"
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
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Target Character"), ValPinkStyle.Render("Feixiao / Acheron")),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Stellar Jade Savings"), ValPinkStyle.Render(fmt.Sprintf("%d Jades (%d Passes)", profile.StellarJade, res.PassesFromJades))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Special Pass Count"), ValPinkStyle.Render(fmt.Sprintf("%d Passes", profile.SpecialPass))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Total Pulls Available"), ValGreenStyle.Render(fmt.Sprintf("%d Pulls", res.TotalPullsAvailable))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Current Pity State"), ValYellowStyle.Render(fmt.Sprintf("%d / %d", profile.CurrentPity, calculator.HSRHardPity))),
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Guarantee Status"), ValTealStyle.Render(pityStatus)),
		"",
		fmt.Sprintf("%s : %s", LabelStyle.Width(22).Render("Planner Status"), statusBadge),
		"",
		ValSubtextStyle.Render(fmt.Sprintf("• Soft Pity Threshold (74) : %d pulls remaining", res.SoftPityRemaining)),
		ValSubtextStyle.Render(fmt.Sprintf("• 5★ Hard Pity Threshold    : %d pulls remaining", res.PullsToFirst5Star)),
		ValSubtextStyle.Render(fmt.Sprintf("• Max Pulls for Target      : %d total pulls needed from 0 pity", res.MaxPullsToTarget)),
	}

	infoBox := lipgloss.NewStyle().
		MarginTop(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorSurface2).
		Padding(0, 1).
		Render(ValSubtextStyle.Render("💡 Tip: Soft pity in Honkai: Star Rail begins at 74 pulls with exponential rate increase up to 90."))

	return HSRCardStyle.
		Width(width - 6).
		Render(lipgloss.JoinVertical(lipgloss.Left, strings.Join(rows, "\n"), infoBox))
}
