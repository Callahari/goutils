package goutils

import (
	"math/rand"
	"time"
)

func RandomIntBetween(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomFloat32Between(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomFloat64Between(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomDateBetween(startDate, endDate time.Time) time.Time {
	//min := time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC).Unix()
	//max := time.Date(2023, 12, 31, 23, 0, 0, 0, time.UTC).Unix()
	delta := endDate.Unix() - startDate.Unix()

	sec := rand.Int63n(delta) + startDate.Unix()
	return time.Unix(sec, 0)
}
