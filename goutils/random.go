package goutils

import (
	"math/rand"
	"time"
)

// RandomIntBetween returns a random integer between the specified min and max values (inclusive).
func RandomIntBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomFloat32Between returns a random float32 value in the range [min, max).
func RandomFloat32Between(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

// RandomFloat64Between returns a random float64 number between the specified min and max values.
func RandomFloat64Between(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomDateBetween generates a random date between the given startDate and endDate.
func RandomDateBetween(startDate, endDate time.Time) time.Time {
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}
	duration := endDate.Sub(startDate)
	if duration == 0 {
		return startDate
	}
	randomDuration := time.Duration(rand.Int63n(int64(duration)))
	return startDate.Add(randomDuration)
}
