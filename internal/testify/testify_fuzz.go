//go:build gofuzz

package testify

import "strconv"

func Fuzz(data []byte) int {
	num, err := strconv.Atoi(string(data))
	if err != nil {
		return 0
	}
	if _, err := F(num); err != nil {
		return 0
	}
	return 1
}
