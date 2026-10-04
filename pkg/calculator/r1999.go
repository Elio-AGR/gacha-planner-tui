package calculator

import "github.com/Elio-AGR/gacha-planner-tui/pkg/models"

const (
	R1999DropPerPull = 180
	R1999SoftPity    = 60
	R1999HardPity    = 70
)

// R1999Result contains calculation outputs for Reverse: 1999.
type R1999Result struct {
	UnilogsFromDrops           int
	TotalPullsAvailable        int
	PullsToFirst6Star          int
	RemainingPullsToFirst6Star int
	MaxPullsToTarget           int
	RemainingPullsToTarget     int
	SoftPityRemaining          int
	CanReachFirst6Star         bool
	CanGuaranteeTarget         bool
}

// CalculateR1999 performs gacha pull calculations based on the user's Reverse: 1999 profile.
func CalculateR1999(profile models.R1999Profile) R1999Result {
	unilogsFromDrops := profile.ClearDrop / R1999DropPerPull
	totalPulls := profile.Unilog + unilogsFromDrops

	// Hard pity for a single 6-star is 70 pulls
	pullsToFirst6Star := R1999HardPity - profile.CurrentPity
	if pullsToFirst6Star < 0 {
		pullsToFirst6Star = 0
	}

	remainingToFirst6Star := pullsToFirst6Star - totalPulls
	if remainingToFirst6Star < 0 {
		remainingToFirst6Star = 0
	}

	// Soft pity check (starts at 60)
	softPityRemaining := R1999SoftPity - profile.CurrentPity
	if softPityRemaining < 0 {
		softPityRemaining = 0
	}

	// Worst case target guarantee: 70 if already guaranteed, 140 if 50/50
	maxPullsToTarget := R1999HardPity
	if !profile.IsGuaranteed {
		maxPullsToTarget = R1999HardPity * 2
	}

	totalEffective := profile.CurrentPity + totalPulls
	remainingToTarget := maxPullsToTarget - totalEffective
	if remainingToTarget < 0 {
		remainingToTarget = 0
	}

	return R1999Result{
		UnilogsFromDrops:           unilogsFromDrops,
		TotalPullsAvailable:        totalPulls,
		PullsToFirst6Star:          pullsToFirst6Star,
		RemainingPullsToFirst6Star: remainingToFirst6Star,
		MaxPullsToTarget:           maxPullsToTarget,
		RemainingPullsToTarget:     remainingToTarget,
		SoftPityRemaining:          softPityRemaining,
		CanReachFirst6Star:         totalPulls >= pullsToFirst6Star,
		CanGuaranteeTarget:         remainingToTarget == 0,
	}
}
