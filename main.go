package main

import (
	"fmt"
	"algorithm/algorithm"
)

func main() {
	var numbers []int = []int{ 1, 18, 19, 20, 21, 109, 111, 200, 302 }

	result, _ := algorithm.DichotomyFromLookup(numbers, 1, 0, len(numbers))

	fmt.Printf("result: %v \n", result)
}

