package randomizer

import "math/rand"

// GetRandom value from slice
func GetRandom(values []string) string {
	return values[rand.Intn(len(values))]
}