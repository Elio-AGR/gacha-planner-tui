package models

// BannerPreset represents a timeless generic banner timeframe estimation.
type BannerPreset struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	EstimatedDaysLeft int    `json:"estimated_days_left"`
	IsCustom          bool   `json:"is_custom"`
}

// GenericBannerPresets defines timeless generic patch timeframe estimations.
var GenericBannerPresets = []BannerPreset{
	{
		ID:                "current_phase",
		Name:              "Current Banner Phase (14 Days)",
		EstimatedDaysLeft: 14,
		IsCustom:          false,
	},
	{
		ID:                "next_phase",
		Name:              "Next Phase (+21 Days)",
		EstimatedDaysLeft: 21,
		IsCustom:          false,
	},
	{
		ID:                "next_patch",
		Name:              "Next Patch / +1 Patch (+42 Days)",
		EstimatedDaysLeft: 42,
		IsCustom:          false,
	},
	{
		ID:                "in_2_patches",
		Name:              "In 2 Patches (+84 Days)",
		EstimatedDaysLeft: 84,
		IsCustom:          false,
	},
	{
		ID:                "anniversary_event",
		Name:              "Major / Anniversary Event (+126 Days)",
		EstimatedDaysLeft: 126,
		IsCustom:          false,
	},
	{
		ID:                "custom_manual",
		Name:              "Custom Manual Input",
		EstimatedDaysLeft: 0,
		IsCustom:          true,
	},
}

var R1999BannerPresets = GenericBannerPresets
var HSRBannerPresets = GenericBannerPresets
