// accelerometer.go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/taigrr/apple-silicon-accelerometer/detector"
	"github.com/taigrr/apple-silicon-accelerometer/sensor"
	"github.com/taigrr/apple-silicon-accelerometer/shm"
)

const (
	pollInterval   = 10 * time.Millisecond
	maxSampleBatch = 200
	sensorStartup  = 100 * time.Millisecond
)

// monitorAccelerometer starts the sensor and sends bonk events on the channel.
// Blocks until ctx is cancelled.
func monitorAccelerometer(ctx context.Context, bonkCh chan<- struct{}) error {
	accelRing, err := shm.CreateRing(shm.NameAccel)
	if err != nil {
		return fmt.Errorf("creating accel shm: %w", err)
	}
	defer accelRing.Close()
	defer accelRing.Unlink()

	sensorReady := make(chan struct{})
	sensorErr := make(chan error, 1)

	go func() {
		close(sensorReady)
		if err := sensor.Run(sensor.Config{
			AccelRing: accelRing,
			Restarts:  0,
		}); err != nil {
			sensorErr <- err
		}
	}()

	select {
	case <-sensorReady:
	case err := <-sensorErr:
		return fmt.Errorf("sensor worker failed: %w", err)
	case <-ctx.Done():
		return nil
	}

	time.Sleep(sensorStartup)

	det := detector.New()
	bd := newBonkDetector(750 * time.Millisecond)
	var lastAccelTotal uint64

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sensorErr:
			return fmt.Errorf("sensor worker failed: %w", err)
		case <-ticker.C:
		}

		now := time.Now()
		tNow := float64(now.UnixNano()) / 1e9

		samples, newTotal := accelRing.ReadNew(lastAccelTotal, shm.AccelScale)
		lastAccelTotal = newTotal
		if len(samples) > maxSampleBatch {
			samples = samples[len(samples)-maxSampleBatch:]
		}

		nSamples := len(samples)
		for idx, sample := range samples {
			tSample := tNow - float64(nSamples-idx-1)/float64(det.FS)
			det.Process(sample.X, sample.Y, sample.Z, tSample)
		}

		if len(det.Events) == 0 {
			continue
		}

		ev := det.Events[len(det.Events)-1]
		det.Events = det.Events[:0]
		if ev.Label == "MAJOR" && bd.check() {
			select {
			case bonkCh <- struct{}{}:
			default:
				// bonk already pending, skip
			}
		}
	}
}
