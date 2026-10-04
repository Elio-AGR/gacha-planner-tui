package calculator_test

import (
	"testing"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/calculator"
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
)

func TestCalculateHSR(t *testing.T) {
	profile := models.HSRProfile{
		StellarJade:  1600, // 10 pulls
		SpecialPass:  5,    // 5 pulls -> total 15 pulls
		CurrentPity:  65,
		IsGuaranteed: true,
		BannerDays:   14,
		DailyIncome:  90,
	}

	result := calculator.CalculateHSR(profile)

	if result.PassesFromJades != 10 {
		t.Errorf("Expected 10 passes from jades, got %d", result.PassesFromJades)
	}

	if result.TotalPullsAvailable != 15 {
		t.Errorf("Expected 15 total pulls, got %d", result.TotalPullsAvailable)
	}

	if result.RemainingPullsToTarget != 10 {
		t.Errorf("Expected 10 remaining pulls to target, got %d", result.RemainingPullsToTarget)
	}

	if result.WinRate <= 0.0 || result.WinRate > 100.0 {
		t.Errorf("Win rate should be between 0.0 and 100.0, got %f", result.WinRate)
	}

	// 14 days * 90 jades = 1260 jades = 7 income pulls
	if result.IncomePulls != 7 {
		t.Errorf("Expected 7 income pulls, got %d", result.IncomePulls)
	}

	if result.ProjectedPulls != 22 {
		t.Errorf("Expected 22 projected pulls, got %d", result.ProjectedPulls)
	}
}

func TestCalculateR1999(t *testing.T) {
	profile := models.R1999Profile{
		ClearDrop:    1800, // 10 pulls
		Unilog:       10,   // 10 pulls -> total 20 pulls
		CurrentPity:  50,
		IsGuaranteed: false,
		BannerDays:   14,
		DailyIncome:  80,
	}

	result := calculator.CalculateR1999(profile)

	if result.UnilogsFromDrops != 10 {
		t.Errorf("Expected 10 unilogs from drops, got %d", result.UnilogsFromDrops)
	}

	if result.TotalPullsAvailable != 20 {
		t.Errorf("Expected 20 total pulls, got %d", result.TotalPullsAvailable)
	}

	if result.WinRate <= 0.0 || result.WinRate > 100.0 {
		t.Errorf("Win rate should be between 0.0 and 100.0, got %f", result.WinRate)
	}

	// 14 days * 80 drops = 1120 drops = 6 income pulls
	if result.IncomePulls != 6 {
		t.Errorf("Expected 6 income pulls, got %d", result.IncomePulls)
	}
}
