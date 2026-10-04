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
	}

	result := calculator.CalculateHSR(profile)

	if result.PassesFromJades != 10 {
		t.Errorf("Expected 10 passes from jades, got %d", result.PassesFromJades)
	}

	if result.TotalPullsAvailable != 15 {
		t.Errorf("Expected 15 total pulls, got %d", result.TotalPullsAvailable)
	}

	// Hard pity = 90. Current pity = 65. Pulls to first 5-star = 25.
	if result.PullsToFirst5Star != 25 {
		t.Errorf("Expected 25 pulls to first 5-star, got %d", result.PullsToFirst5Star)
	}

	// Remaining pulls to first 5-star = 25 - 15 = 10
	if result.RemainingPullsToFirst5Star != 10 {
		t.Errorf("Expected 10 remaining pulls to first 5-star, got %d", result.RemainingPullsToFirst5Star)
	}

	// Since guaranteed = true, max pulls to target = 90.
	// Total effective pulls = 65 + 15 = 80. Remaining = 90 - 80 = 10.
	if result.RemainingPullsToTarget != 10 {
		t.Errorf("Expected 10 remaining pulls to target, got %d", result.RemainingPullsToTarget)
	}
}

func TestCalculateR1999(t *testing.T) {
	profile := models.R1999Profile{
		ClearDrop:    1800, // 10 pulls
		Unilog:       10,   // 10 pulls -> total 20 pulls
		CurrentPity:  50,
		IsGuaranteed: false, // Needs 140 max pulls for guarantee
	}

	result := calculator.CalculateR1999(profile)

	if result.UnilogsFromDrops != 10 {
		t.Errorf("Expected 10 unilogs from drops, got %d", result.UnilogsFromDrops)
	}

	if result.TotalPullsAvailable != 20 {
		t.Errorf("Expected 20 total pulls, got %d", result.TotalPullsAvailable)
	}

	// Soft pity = 60. Current pity = 50. Soft pity remaining = 10.
	if result.SoftPityRemaining != 10 {
		t.Errorf("Expected 10 soft pity remaining, got %d", result.SoftPityRemaining)
	}

	// Max pulls to target (50/50) = 140.
	// Total effective = 50 + 20 = 70. Remaining = 140 - 70 = 70.
	if result.RemainingPullsToTarget != 70 {
		t.Errorf("Expected 70 remaining pulls to target, got %d", result.RemainingPullsToTarget)
	}
}
