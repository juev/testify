//go:build gofuzz

package main

import fuzz "github.com/google/gofuzz"

func Fuzz(data []byte) int {
	var num int
	fuzz.NewFromGoFuzz(data).Fuzz(&num)
	f(num)
	return 0
}
