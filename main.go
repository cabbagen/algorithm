package main

import (
	"fmt"
	"algorithm/algorithm"
)

func main() {
	index, isExist := algorithm.DichotomyFromLookup([]int {1, 2, 6, 9, 12, 14, 91}, 1, 0, 7)
	fmt.Printf("index: %d  result: %v \n", index, isExist)
}
