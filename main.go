package testify

import "fmt"

func main() {
	fmt.Printf("%v", f(3))
}

func f(num int) float32 {
	if num == 1 {
		panic("")
	}
	return float32(1) / float32(num-5)
}
