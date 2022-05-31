package storage

import (
	"math/rand"
	"time"
)

func genId() string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz"
	var result string
	for i := 0; i < 20; i++ {
		result += string(chars[rand.Intn(len(chars))])
	}
	return result
}
