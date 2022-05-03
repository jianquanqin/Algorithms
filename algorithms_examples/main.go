package main

import (
	"algorithms_examples/algorithms_examples/src"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10}
	fmt.Println(src.BinarySearch(nums, 8))
}
