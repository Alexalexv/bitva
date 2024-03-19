package models

import "math/rand"

func RandGenerator(min, max int) int {
	max++
	return rand.Intn(max-min) + min
}
