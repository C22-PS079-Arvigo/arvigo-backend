package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	randomChars = "abcdefghijklmnopqrstuvwxyz0123456789"
)

func GenerateRandomStringWithTimestamp(randomLen int) string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random string
	randomString := make([]byte, randomLen)
	for i := 0; i < randomLen; i++ {
		randomString[i] = randomChars[rand.Intn(len(randomChars))]
	}

	// Get the current timestamp
	timestamp := time.Now().Unix()

	// Combine the random string and timestamp
	result := fmt.Sprintf("%s_%d", randomString, timestamp)

	return result
}
