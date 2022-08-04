package main

func main() {}

//判断是否有重复元素

func containsDuplicate(nums []int) bool {

	m := make(map[int]bool)

	for _,v := range nums {

		if m[v] == true {
			return true
		}else {
			m[v] = true
		}

	}
	return false

}

//求最大子序列和

func maxSubArray(nums []int) int {

	//初始数组
	max := nums[0]
     for i := 1; i < len(nums); i++ {

        if nums[i] + nums[i-1] > nums[i] {
            nums[i] += nums[i-1]
        }
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}

//day2

func twoSum(nums []int, target int) []int {

	var res []int

	for i:=0;i<len(nums);i++ {
		for j:=i+1;j<len(nums);j++ {
			if nums[j] + nums[i] == target {
				return append(res, i,j)
			}
		}
	}

	return res
}

func merge(nums1 []int, m int, nums2 []int, n int)  {

	for i:=0;i<n;i++ {

		
	}
	

}func merge(nums1 []int, m int, nums2 []int, n int)  {

    sorted := make([]int, 0, m+n)
    p1, p2 := 0, 0
    for {
        if p1 == m {
            sorted = append(sorted, nums2[p2:]...)
            break
        }
        if p2 == n {
            sorted = append(sorted, nums1[p1:]...)
            break
        }
        if nums1[p1] < nums2[p2] {
            sorted = append(sorted, nums1[p1])
            p1++
        } else {
            sorted = append(sorted, nums2[p2])
            p2++
        }
    }
    copy(nums1, sorted)

}