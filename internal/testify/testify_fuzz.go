//go:build gofuzz

package testify

import fuzz "github.com/google/gofuzz"

func Fuzz(data []byte) int {
	var num int
	fuzz.NewFromGoFuzz(data).Fuzz(&num)
	if _, err := F(num); err != nil {
		return 0
	}
	return 1
}
