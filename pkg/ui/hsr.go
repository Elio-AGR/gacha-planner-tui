package ui

import (
	"fmt"
	"strings"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/calculator"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// RenderHSRView renders the 1-column interactive Honkai: Star Rail planner tab.
func RenderHSRView(profile models.HSRProfile, inputs [5]textinput.Model, focusIndex int, width int, height int) string {
	res := calculator.CalculateHSR(profile)

	presetIndex := profile.TargetBannerIndex
	if presetIndex < 0 || presetIndex >= len(models.HSRBannerPresets) {
		presetIndex = 0
	}
	currentPreset := models.HSRBannerPresets[presetIndex]

	title := lipgloss.NewStyle().Bold(true).Foreground(ColorPink).Render("🌌 HONKAI: STAR RAIL PLANNER & BANNER TIMEFRAME SELECTOR")

	pityStatus := "50/50 Next"
	if profile.IsGuaranteed {
		pityStatus = "GUARANTEED"
	}

	var statusBadge string
	if res.CanGuaranteeTarget {
		statusBadge = StatusReadyStyle.Render("✅ GUARANTEED SECURED NOW")
	} else {
		statusBadge = StatusNeedPullsStyle.Render(fmt.Sprintf("⚡ NEED %d MORE PULLS NOW", res.RemainingPullsToTarget))
	}

	var projBadge string
	if res.ProjectedCanGuarantee {
		projBadge = StatusReadyStyle.Render("✅ TARGET ACHIEVABLE BY END OF BANNER")
	} else {
		projBadge = StatusNeedPullsStyle.Render(fmt.Sprintf("⚡ SHORT BY %d PULLS AT END OF BANNER", res.ProjectedShortfall))
	}

	pointer := func(idx int) string {
		if idx == focusIndex {
			return lipgloss.NewStyle().Foreground(ColorPink).Bold(true).Render("▶ ")
		}
		return "  "
	}

	// Toggle Row Text (Focus Index 3)
	toggleText := "[ ⚡ 50/50 NEXT ] (Space/g/Enter)"
	if profile.IsGuaranteed {
		toggleText = "[ ✅ GUARANTEED ] (Space/g/Enter)"
	}
	if focusIndex == 3 {
		toggleText = lipgloss.NewStyle().Foreground(ColorCrust).Background(ColorPink).Bold(true).Render(toggleText)
	} else {
		toggleText = lipgloss.NewStyle().Foreground(ColorPink).Bold(true).Render(toggleText)
	}

	// Preset Selector Text (Focus Index 6)
	bannerSelectorText := fmt.Sprintf("[◄ %s ►] (←/→/Space)", currentPreset.Name)
	if focusIndex == 6 {
		bannerSelectorText = lipgloss.NewStyle().Foreground(ColorCrust).Background(ColorPink).Bold(true).Render(bannerSelectorText)
	} else {
		bannerSelectorText = lipgloss.NewStyle().Foreground(ColorPink).Bold(true).Render(bannerSelectorText)
	}

	lbl := func(text string) string {
		return LabelStyle.Render(fmt.Sprintf("%-24s", text))
	}

	// 1-Column Form Rows in exact requested order
	row1 := fmt.Sprintf("%s%s: %s", pointer(0), lbl("Stellar Jades"), inputs[0].View())
	row2 := fmt.Sprintf("%s%s: %s", pointer(1), lbl("Special Passes"), inputs[1].View())
	row3 := fmt.Sprintf("%s%s: %s", pointer(2), lbl("Current Pity (0-90)"), inputs[2].View())
	row4 := fmt.Sprintf("%s%s: %s", pointer(3), lbl("Rate-Up Guaranteed"), toggleText)
	row5 := fmt.Sprintf("%s%s: %s", pointer(4), lbl("Days Remaining (Manual)"), inputs[3].View())
	row6 := fmt.Sprintf("%s%s: %s", pointer(5), lbl("Daily Income (Jades)"), inputs[4].View())
	row7 := fmt.Sprintf("%s%s: %s", pointer(6), lbl("Quick Timeframe Preset"), bannerSelectorText)
	row7Hint := lipgloss.NewStyle().Foreground(ColorSubtext0).Italic(true).Render("   ↳ Quick Fill Helper (Auto-sets Days Remaining; manual edit supported)")

	formBox := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(ColorSurface2).
		Padding(0, 1).
		Render(strings.Join([]string{row1, row2, row3, row4, row5, row6, row7, row7Hint}, "\n"))

	winRateStr := fmt.Sprintf("%.1f%%", res.WinRate)

	calcRows := []string{
		fmt.Sprintf("%s %s | %s %s | %s %s",
			LabelStyle.Render("Timeframe Preset:"), ValPinkStyle.Render(currentPreset.Name),
			LabelStyle.Render("Current Pulls:"), ValGreenStyle.Render(fmt.Sprintf("%d", res.TotalPullsAvailable)),
			LabelStyle.Render("Win Rate:"), ValGreenStyle.Render(winRateStr),
		),
		fmt.Sprintf("%s %s | %s %s | %s %s",
			LabelStyle.Render("Current Pity    :"), ValYellowStyle.Render(fmt.Sprintf("%d / %d (%s)", profile.CurrentPity, calculator.HSRHardPity, pityStatus)),
			LabelStyle.Render("Soft Pity in:"), ValSubtextStyle.Render(fmt.Sprintf("%dp", res.SoftPityRemaining)),
			LabelStyle.Render("Hard Pity in:"), ValSubtextStyle.Render(fmt.Sprintf("%dp", res.PullsToFirst5Star)),
		),
		fmt.Sprintf("%s %s", LabelStyle.Render("Current Status       :"), statusBadge),
		fmt.Sprintf("%s +%d Pulls (+%d Jades in %dd) ➔ %s %d Total Pulls",
			LabelStyle.Render("Projected Income     :"), res.IncomePulls, profile.BannerDays*profile.DailyIncome, profile.BannerDays,
			LabelStyle.Render("End-of-Banner:"), res.ProjectedPulls,
		),
		fmt.Sprintf("%s %s", LabelStyle.Render("Projection Status    :"), projBadge),
	}

	calcBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorPink).
		Padding(0, 1).
		Render(strings.Join(calcRows, "\n"))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		formBox,
		calcBox,
	)
}
