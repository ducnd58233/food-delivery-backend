package common

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSequence(n int) string {
	b := make([]rune, n)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s) // generate seed

	for i := range b {
		b[i] = letters[r.Intn(99999)%len(letters)]
	}

	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}

	return randSequence(length)
}
