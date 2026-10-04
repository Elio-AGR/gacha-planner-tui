package calculator

import "github.com/Elio-AGR/gacha-planner-tui/pkg/models"

const (
	HSRJadePerPull = 160
	HSRSoftPity    = 74
	HSRHardPity    = 90
)

// HSRResult contains calculation outputs for Honkai: Star Rail.
type HSRResult struct {
	PassesFromJades            int
	TotalPullsAvailable        int
	PullsToFirst5Star          int
	RemainingPullsToFirst5Star int
	MaxPullsToTarget           int
	RemainingPullsToTarget     int
	SoftPityRemaining          int
	CanReachFirst5Star         bool
	CanGuaranteeTarget         bool
}

// CalculateHSR performs gacha pull calculations based on the user's HSR profile.
func CalculateHSR(profile models.HSRProfile) HSRResult {
	passesFromJades := profile.StellarJade / HSRJadePerPull
	totalPulls := profile.SpecialPass + passesFromJades

	// Hard pity for a single 5-star is 90 pulls
	pullsToFirst5Star := HSRHardPity - profile.CurrentPity
	if pullsToFirst5Star < 0 {
		pullsToFirst5Star = 0
	}

	remainingToFirst5Star := pullsToFirst5Star - totalPulls
	if remainingToFirst5Star < 0 {
		remainingToFirst5Star = 0
	}

	// Soft pity check (starts at 74)
	softPityRemaining := HSRSoftPity - profile.CurrentPity
	if softPityRemaining < 0 {
		softPityRemaining = 0
	}

	// Worst case target guarantee: 90 if already guaranteed, 180 if 50/50
	maxPullsToTarget := HSRHardPity
	if !profile.IsGuaranteed {
		maxPullsToTarget = HSRHardPity * 2
	}

	totalEffective := profile.CurrentPity + totalPulls
	remainingToTarget := maxPullsToTarget - totalEffective
	if remainingToTarget < 0 {
		remainingToTarget = 0
	}

	return HSRResult{
		PassesFromJades:            passesFromJades,
		TotalPullsAvailable:        totalPulls,
		PullsToFirst5Star:          pullsToFirst5Star,
		RemainingPullsToFirst5Star: remainingToFirst5Star,
		MaxPullsToTarget:           maxPullsToTarget,
		RemainingPullsToTarget:     remainingToTarget,
		SoftPityRemaining:          softPityRemaining,
		CanReachFirst5Star:         totalPulls >= pullsToFirst5Star,
		CanGuaranteeTarget:         remainingToTarget == 0,
	}
}
