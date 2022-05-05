package src

import (
	"fmt"
	"testing"
)

func TestLastAndFirstAppearance(t *testing.T) {

	nums := []int{1, 2, 2, 4, 5, 6, 7}
	fmt.Println(LastAppearance(nums, 2))
	fmt.Println(FirstAppearance(nums, 2))

}

//33. 搜索旋转排序数组

func TestSearch(t *testing.T) {

	nums := []int{4, 5, 6, 7, 8, 1, 2, 3}

	fmt.Println(Search(nums, 2))

}

//154.寻找旋转排序数组中的最小值I
//155.寻找旋转排序数组中的最小值II

func TestFindMin(t *testing.T) {
	nums := []int{3, 3, 1, 3, 3, 3, 3, 3}
	fmt.Println(FindMinII(nums))
	//fmt.Println(FindMinI(nums))
}

//74. 搜索二维矩阵

func TestSearchMatrix(t *testing.T) {
	//nums := []int{3, 4, 5, 1, 2}
	//fmt.Println(findMin(nums))
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	//matrix := [][]int{{0}, {1}, {2}}
	target := 0

	fmt.Println(searchMatrix(matrix, target))
}

//34. 在排序数组中查找元素的第一个和最后一个位置

func TestBinarySearch(t *testing.T) {
	nums := []int{0}
	fmt.Println(searchRange(nums, 0))

}

//81. 搜索旋转排序数组 II

func TestSearchII(t *testing.T) {
	nums := []int{1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	fmt.Println(SearchII(nums, 0))
}

//162. 寻找峰值

func TestFindPeakElement(t *testing.T) {

	nums := []int{1, 2, 0, 3, 1, 5, 6, 3}

	fmt.Println(FindPeakElement(nums))

}

func TestTwoSum(t *testing.T) {

	nums := []int{-1, -1, 1}

	fmt.Println(TwoSum(nums, -2))

}
