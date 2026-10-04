package models

// HSRProfile stores the gacha savings and pity state for Honkai: Star Rail.
type HSRProfile struct {
	StellarJade       int  `json:"stellar_jade"`
	SpecialPass       int  `json:"special_pass"`
	CurrentPity       int  `json:"current_pity"`
	IsGuaranteed      bool `json:"is_guaranteed"`
	BannerDays        int  `json:"banner_days"`
	DailyIncome       int  `json:"daily_income"`
	TargetBannerIndex int  `json:"target_banner_index"`
}

// R1999Profile stores the gacha savings and pity state for Reverse: 1999.
type R1999Profile struct {
	ClearDrop         int  `json:"clear_drop"`
	Unilog            int  `json:"unilog"`
	CurrentPity       int  `json:"current_pity"`
	IsGuaranteed      bool `json:"is_guaranteed"`
	BannerDays        int  `json:"banner_days"`
	DailyIncome       int  `json:"daily_income"`
	TargetBannerIndex int  `json:"target_banner_index"`
}

// UserProfile aggregates all game profiles for the user.
type UserProfile struct {
	HSR   HSRProfile   `json:"hsr"`
	R1999 R1999Profile `json:"r1999"`
}

// NewDefaultProfile creates a profile with initial sample data.
func NewDefaultProfile() UserProfile {
	return UserProfile{
		HSR: HSRProfile{
			StellarJade:       19200,
			SpecialPass:       15,
			CurrentPity:       65,
			IsGuaranteed:      true,
			BannerDays:        14,
			DailyIncome:       90,
			TargetBannerIndex: 0,
		},
		R1999: R1999Profile{
			ClearDrop:         14400,
			Unilog:            68,
			CurrentPity:       42,
			IsGuaranteed:      false,
			BannerDays:        14,
			DailyIncome:       80,
			TargetBannerIndex: 0,
		},
	}
}

// NewZeroProfile creates a profile with zeroed currency values and standard defaults.
func NewZeroProfile() UserProfile {
	return UserProfile{
		HSR: HSRProfile{
			StellarJade:       0,
			SpecialPass:       0,
			CurrentPity:       0,
			IsGuaranteed:      false,
			BannerDays:        14,
			DailyIncome:       90,
			TargetBannerIndex: 0,
		},
		R1999: R1999Profile{
			ClearDrop:         0,
			Unilog:            0,
			CurrentPity:       0,
			IsGuaranteed:      false,
			BannerDays:        14,
			DailyIncome:       80,
			TargetBannerIndex: 0,
		},
	}
}
