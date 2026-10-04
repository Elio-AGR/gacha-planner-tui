package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Elio-AGR/gacha-planner-tui/pkg/config"
)

func TestSaveAndLoadProfile(t *testing.T) {
	// Override HOME env variable for temporary test directory
	tempHome := t.TempDir()
	t.Setenv("HOME", tempHome)

	expectedPath := filepath.Join(tempHome, ".config", "gacha-planner", "data.json")

	// 1. Initial Load when file doesn't exist -> should create default 0 profile file
	profile, err := config.LoadProfile()
	if err != nil {
		t.Fatalf("Unexpected error loading profile: %v", err)
	}

	if profile.HSR.StellarJade != 0 || profile.R1999.ClearDrop != 0 {
		t.Errorf("Expected zero initial values, got HSR Jade=%d, R1999 Drop=%d", profile.HSR.StellarJade, profile.R1999.ClearDrop)
	}

	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Fatalf("Expected file %s to be created, but it was missing", expectedPath)
	}

	// 2. Modify & Save Profile
	profile.HSR.StellarJade = 16000
	profile.HSR.SpecialPass = 20
	profile.HSR.CurrentPity = 50
	profile.HSR.IsGuaranteed = true

	profile.R1999.ClearDrop = 18000
	profile.R1999.Unilog = 15
	profile.R1999.CurrentPity = 30
	profile.R1999.IsGuaranteed = false

	if err := config.SaveProfile(profile); err != nil {
		t.Fatalf("Failed to save profile: %v", err)
	}

	// 3. Reload Profile and verify
	reloaded, err := config.LoadProfile()
	if err != nil {
		t.Fatalf("Failed to reload profile: %v", err)
	}

	if reloaded.HSR.StellarJade != 16000 || reloaded.R1999.ClearDrop != 18000 {
		t.Errorf("Reloaded profile mismatch! HSR Jade=%d, R1999 Drop=%d", reloaded.HSR.StellarJade, reloaded.R1999.ClearDrop)
	}
}
