package goutils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

// RandomBasedOnProbability returns true or false based on a given probability represented as a string "numerator:denominator".
// It splits the input string, validates, and converts the parts to integers. It then generates a random number between 1 and denominator,
// and checks if this number is within the bounds of the numerator, returning true or false accordingly.
func RandomBasedOnProbability(prob string) (bool, error) {
	// Split the input string at ":"
	parts := strings.Split(prob, ":")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid probability format")
	}
	// Convert the parts to integers
	numerator, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, fmt.Errorf("invalid numerator: %v", err)
	}
	denominator, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("invalid denominator: %v", err)
	}
	// Error if denominator is zero
	if denominator == 0 {
		return false, fmt.Errorf("denominator cannot be zero")
	}
	// Generate a random number and compare it
	rand.Seed(time.Now().UnixNano())
	randomValue := rand.Intn(denominator) + 1
	return randomValue <= numerator, nil
}
