// detector_test.go
package main

import (
	"testing"
	"time"
)

func TestBonkDetectorCooldown(t *testing.T) {
	d := newBonkDetector(2 * time.Second)

	// First bonk should fire
	if !d.check() {
		t.Error("first bonk should trigger")
	}

	// Immediate second bonk should be suppressed by cooldown
	if d.check() {
		t.Error("second bonk during cooldown should not trigger")
	}
}

func TestBonkDetectorAfterCooldown(t *testing.T) {
	d := newBonkDetector(50 * time.Millisecond)

	if !d.check() {
		t.Error("first bonk should trigger")
	}

	// Wait for cooldown to expire
	time.Sleep(60 * time.Millisecond)

	if !d.check() {
		t.Error("bonk after cooldown should trigger")
	}
}
