// detector.go
package main

import (
	"time"
)

// bonkDetector tracks bonk events with threshold and cooldown.
type bonkDetector struct {
	threshold float64
	cooldown  time.Duration
	lastBonk  time.Time
}

func newBonkDetector(threshold float64, cooldown time.Duration) *bonkDetector {
	return &bonkDetector{
		threshold: threshold,
		cooldown:  cooldown,
	}
}

// check returns true if the amplitude exceeds threshold and cooldown has elapsed.
func (d *bonkDetector) check(amplitude float64) bool {
	if amplitude < d.threshold {
		return false
	}
	now := time.Now()
	if !d.lastBonk.IsZero() && now.Sub(d.lastBonk) < d.cooldown {
		return false
	}
	d.lastBonk = now
	return true
}
