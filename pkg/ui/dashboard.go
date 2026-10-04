package ui

import (
	"fmt"
	"strings"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/calculator"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
	"github.com/charmbracelet/lipgloss"
)

// RenderDashboard renders the side-by-side Dashboard view using Catppuccin Mocha styles.
func RenderDashboard(profile models.UserProfile, width int, height int) string {
	hsrRes := calculator.CalculateHSR(profile.HSR)
	r1999Res := calculator.CalculateR1999(profile.R1999)

	availableWidth := width - 6
	if availableWidth < 20 {
		availableWidth = 20
	}

	// Determine side-by-side vs stacked layout based on window width
	isSideBySide := availableWidth >= 70
	var cardWidth int

	if isSideBySide {
		cardWidth = (availableWidth - 3) / 2
	} else {
		cardWidth = availableWidth
	}

	hsrCard := renderHSRDashboardCard(profile.HSR, hsrRes, cardWidth)
	r1999Card := renderR1999DashboardCard(profile.R1999, r1999Res, cardWidth)

	var cardsContainer string
	if isSideBySide {
		cardsContainer = lipgloss.JoinHorizontal(lipgloss.Top, hsrCard, "  ", r1999Card)
	} else {
		cardsContainer = lipgloss.JoinVertical(lipgloss.Left, hsrCard, r1999Card)
	}

	totalCombinedPulls := hsrRes.TotalPullsAvailable + r1999Res.TotalPullsAvailable

	summaryBox := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(ColorMauve).
		Padding(0, 1).
		Width(availableWidth).
		Render(fmt.Sprintf(
			"%s %s",
			lipgloss.NewStyle().Foreground(ColorMauve).Bold(true).Render("🎲 Combined Savings Overview:"),
			lipgloss.NewStyle().Foreground(ColorGreen).Bold(true).Render(fmt.Sprintf("%d Total Pulls Available Across All Games", totalCombinedPulls)),
		))

	return lipgloss.JoinVertical(lipgloss.Left, cardsContainer, "", summaryBox)
}

func renderHSRDashboardCard(p models.HSRProfile, res calculator.HSRResult, cardWidth int) string {
	title := lipgloss.NewStyle().Bold(true).Foreground(ColorPink).Render("🌌 Honkai: Star Rail")

	pityStatus := "50/50 Next"
	if p.IsGuaranteed {
		pityStatus = "GUARANTEED"
	}

	// Status Indicator Badge
	var statusBadge string
	if res.CanGuaranteeTarget {
		statusBadge = StatusReadyStyle.Render("✅ Ready for Guarantee")
	} else {
		statusBadge = StatusNeedPullsStyle.Render(fmt.Sprintf("⚡ Need %d More Pulls", res.RemainingPullsToTarget))
	}

	lines := []string{
		title,
		"",
		fmt.Sprintf("%s %s", LabelStyle.Render("Total Savings:"), ValPinkStyle.Render(fmt.Sprintf("%d Jades + %d Passes", p.StellarJade, p.SpecialPass))),
		fmt.Sprintf("%s %s", LabelStyle.Render("Total Pulls  :"), ValGreenStyle.Render(fmt.Sprintf("%d Pulls", res.TotalPullsAvailable))),
		fmt.Sprintf("%s %s (%s)", LabelStyle.Render("Current Pity :"), ValYellowStyle.Render(fmt.Sprintf("%d / %d", p.CurrentPity, calculator.HSRHardPity)), pityStatus),
		"",
		fmt.Sprintf("%s %s", LabelStyle.Render("Status       :"), statusBadge),
		"",
		ValSubtextStyle.Render(fmt.Sprintf("• Soft Pity in : %d pulls", res.SoftPityRemaining)),
		ValSubtextStyle.Render(fmt.Sprintf("• 5★ Guarantee : %d pulls", res.PullsToFirst5Star)),
	}

	return HSRCardStyle.
		Width(cardWidth).
		Render(strings.Join(lines, "\n"))
}

func renderR1999DashboardCard(p models.R1999Profile, res calculator.R1999Result, cardWidth int) string {
	title := lipgloss.NewStyle().Bold(true).Foreground(ColorTeal).Render("📜 Reverse: 1999")

	pityStatus := "50/50 Next"
	if p.IsGuaranteed {
		pityStatus = "GUARANTEED"
	}

	// Status Indicator Badge
	var statusBadge string
	if res.CanGuaranteeTarget {
		statusBadge = StatusReadyStyle.Render("✅ Ready for Guarantee")
	} else {
		statusBadge = StatusNeedPullsStyle.Render(fmt.Sprintf("⚡ Need %d More Pulls", res.RemainingPullsToTarget))
	}

	lines := []string{
		title,
		"",
		fmt.Sprintf("%s %s", LabelStyle.Render("Total Savings:"), ValTealStyle.Render(fmt.Sprintf("%d Drops + %d Unilogs", p.ClearDrop, p.Unilog))),
		fmt.Sprintf("%s %s", LabelStyle.Render("Total Pulls  :"), ValGreenStyle.Render(fmt.Sprintf("%d Pulls", res.TotalPullsAvailable))),
		fmt.Sprintf("%s %s (%s)", LabelStyle.Render("Current Pity :"), ValYellowStyle.Render(fmt.Sprintf("%d / %d", p.CurrentPity, calculator.R1999HardPity)), pityStatus),
		"",
		fmt.Sprintf("%s %s", LabelStyle.Render("Status       :"), statusBadge),
		"",
		ValSubtextStyle.Render(fmt.Sprintf("• Soft Pity in : %d pulls", res.SoftPityRemaining)),
		ValSubtextStyle.Render(fmt.Sprintf("• 6★ Guarantee : %d pulls", res.PullsToFirst6Star)),
	}

	return R1999CardStyle.
		Width(cardWidth).
		Render(strings.Join(lines, "\n"))
}
