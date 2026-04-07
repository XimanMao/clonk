// detector.go
package main

import (
	"time"
)

// bonkDetector tracks bonk events with a cooldown.
type bonkDetector struct {
	cooldown time.Duration
	lastBonk time.Time
}

func newBonkDetector(cooldown time.Duration) *bonkDetector {
	return &bonkDetector{
		cooldown: cooldown,
	}
}

// check returns true if the cooldown has elapsed since the last bonk.
func (d *bonkDetector) check() bool {
	now := time.Now()
	if !d.lastBonk.IsZero() && now.Sub(d.lastBonk) < d.cooldown {
		return false
	}
	d.lastBonk = now
	return true
}
