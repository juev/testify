package main

import "fmt"

func main() {
	fmt.Printf("%v", f(3))
}

func f(num int) float32 {
	return float32(1) / float32(num-5)
}