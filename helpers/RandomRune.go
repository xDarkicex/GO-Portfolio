package helpers

import (
	"math/rand"
	"time"
)

// RandStringRunes generates a random string of given length using specified runes.
func RandStringRunes(n int) string {
	const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rand.Seed(time.Now().UnixNano())

	runes := []rune(letterRunes)
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	return string(b)
}

// RandRune generates a random rune from the specified runes.
func RandRune() rune {
	const letterRunes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	rand.Seed(time.Now().UnixNano())

	runes := []rune(letterRunes)
	return runes[rand.Intn(len(runes))]
}
