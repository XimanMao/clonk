// detector_test.go
package main

import (
	"testing"
	"time"
)

func TestBonkDetectorCooldown(t *testing.T) {
	d := newBonkDetector(0.05, 2*time.Second)

	// First bonk should fire
	if !d.check(0.1) {
		t.Error("first bonk should trigger")
	}

	// Immediate second bonk should be suppressed by cooldown
	if d.check(0.1) {
		t.Error("second bonk during cooldown should not trigger")
	}
}

func TestBonkDetectorBelowThreshold(t *testing.T) {
	d := newBonkDetector(0.05, 2*time.Second)

	// Below threshold should not fire
	if d.check(0.01) {
		t.Error("amplitude below threshold should not trigger")
	}
}

func TestBonkDetectorAfterCooldown(t *testing.T) {
	d := newBonkDetector(0.05, 50*time.Millisecond)

	if !d.check(0.1) {
		t.Error("first bonk should trigger")
	}

	// Wait for cooldown to expire
	time.Sleep(60 * time.Millisecond)

	if !d.check(0.1) {
		t.Error("bonk after cooldown should trigger")
	}
}
