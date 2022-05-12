package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {

	nums := []int{0, 1, 2, 4, 1, 4, 5, 6, 2}

	fmt.Println(containsDuplicate(nums))
}
