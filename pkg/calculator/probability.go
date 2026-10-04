package calculator

// calculateWinRate calculates the cumulative probability (%) of acquiring the target character
// considering current pity, total available pulls, guaranteed status (50/50), soft pity, and hard pity.
func calculateWinRate(pity int, pulls int, isGuaranteed bool, softPity int, hardPity int, baseRate float64, softRamp float64) float64 {
	maxNeeded := hardPity
	if !isGuaranteed {
		maxNeeded = hardPity * 2
	}

	totalPity := pity + pulls
	if totalPity >= maxNeeded {
		return 100.0
	}
	if pulls <= 0 {
		return 0.0
	}

	// Calculate probability of getting 5-star/6-star in 1st cycle
	probNotPulling1st := 1.0
	currPity := pity
	for i := 0; i < pulls; i++ {
		currPity++
		rate := baseRate
		if currPity >= softPity {
			rate = baseRate + softRamp*float64(currPity-softPity+1)
			if rate > 1.0 {
				rate = 1.0
			}
		}
		probNotPulling1st *= (1.0 - rate)
		if currPity >= hardPity {
			probNotPulling1st = 0.0
			break
		}
	}
	probGet1st := 1.0 - probNotPulling1st

	if isGuaranteed {
		winRate := probGet1st * 100.0
		if winRate > 100.0 {
			return 100.0
		}
		return winRate
	}

	// 50/50 Case:
	// 50% chance win on 1st pity drop
	winRate := probGet1st * 50.0

	// If enough pulls remain for a 2nd pity cycle after 1st 5-star drop
	if pulls > (hardPity - pity) {
		remainingPulls := pulls - (hardPity - pity)
		probNotPulling2nd := 1.0
		currPity2 := 0
		for i := 0; i < remainingPulls; i++ {
			currPity2++
			rate := baseRate
			if currPity2 >= softPity {
				rate = baseRate + softRamp*float64(currPity2-softPity+1)
				if rate > 1.0 {
					rate = 1.0
				}
			}
			probNotPulling2nd *= (1.0 - rate)
			if currPity2 >= hardPity {
				probNotPulling2nd = 0.0
				break
			}
		}
		probGet2nd := 1.0 - probNotPulling2nd
		winRate += probGet1st * 0.5 * probGet2nd * 100.0
	}

	if winRate > 100.0 {
		return 100.0
	}
	if winRate < 0.0 {
		return 0.0
	}
	return winRate
}
