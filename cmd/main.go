package main

import (
	"fmt"

	"github.com/juev/testify/internal/testify"
)

func main() {
	result, _ := testify.F(3)
	fmt.Printf("%v", result)
}
