package main

import (
	"fmt"

	"github.com/prime-factors/service"
)

func main() {
	numbers := []int{231123, 41231, 121231, 11237, 1312333}

	c1 := service.PrimesCheckWithoutGoroutine(numbers)
	c2 := service.PrimesCheckWithGoroutine(numbers)

	fmt.Println(c1)
	fmt.Println(c2)
}
