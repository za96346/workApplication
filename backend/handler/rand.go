package handler

import (
	"math/rand"
	"time"
)

func Rand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return randomInt(min, max)
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
