package helpers

import (
	"math/rand"
	"time"
)

func NomorAcak(n int) int {
	rand.Seed(time.Now().UnixNano())
	
	value := rand.Intn(n)
	return value
}