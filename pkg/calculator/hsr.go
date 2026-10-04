package calculator

import (
	"github.com/Elio-AGR/gacha-planner-tui/pkg/models"
)

const (
	HSRJadePerPull = 160
	HSRSoftPity    = 74
	HSRHardPity    = 90
	HSRBaseRate    = 0.006
	HSRSoftRamp    = 0.06
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

	// Win Rate & Banner Projection
	WinRate               float64
	IncomePulls           int
	ProjectedPulls        int
	ProjectedShortfall    int
	ProjectedCanGuarantee bool
}

// CalculateHSR performs gacha pull calculations based on the user's HSR profile.
func CalculateHSR(profile models.HSRProfile) HSRResult {
	passesFromJades := profile.StellarJade / HSRJadePerPull
	totalPulls := profile.SpecialPass + passesFromJades

	pullsToFirst5Star := HSRHardPity - profile.CurrentPity
	if pullsToFirst5Star < 0 {
		pullsToFirst5Star = 0
	}

	remainingToFirst5Star := pullsToFirst5Star - totalPulls
	if remainingToFirst5Star < 0 {
		remainingToFirst5Star = 0
	}

	softPityRemaining := HSRSoftPity - profile.CurrentPity
	if softPityRemaining < 0 {
		softPityRemaining = 0
	}

	maxPullsToTarget := HSRHardPity
	if !profile.IsGuaranteed {
		maxPullsToTarget = HSRHardPity * 2
	}

	totalEffective := profile.CurrentPity + totalPulls
	remainingToTarget := maxPullsToTarget - totalEffective
	if remainingToTarget < 0 {
		remainingToTarget = 0
	}

	// Win Rate Calculation
	winRate := calculateWinRate(profile.CurrentPity, totalPulls, profile.IsGuaranteed, HSRSoftPity, HSRHardPity, HSRBaseRate, HSRSoftRamp)

	// Daily Income & Banner Countdown Projection
	days := profile.BannerDays
	if days < 0 {
		days = 0
	}
	dailyInc := profile.DailyIncome
	if dailyInc < 0 {
		dailyInc = 0
	}

	incomeJades := days * dailyInc
	incomePulls := incomeJades / HSRJadePerPull
	projectedPulls := totalPulls + incomePulls
	projectedEffective := profile.CurrentPity + projectedPulls

	projectedShortfall := maxPullsToTarget - projectedEffective
	projectedCanGuarantee := projectedShortfall <= 0
	if projectedShortfall < 0 {
		projectedShortfall = 0
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

		WinRate:               winRate,
		IncomePulls:           incomePulls,
		ProjectedPulls:        projectedPulls,
		ProjectedShortfall:    projectedShortfall,
		ProjectedCanGuarantee: projectedCanGuarantee,
	}
}
