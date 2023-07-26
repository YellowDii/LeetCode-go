package arrays

import "sort"

// 18.四数之和
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			fourth := n - 1
			sum := target - nums[first] - nums[second]
			for third := second + 1; third < n; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				for third < fourth && nums[third]+nums[fourth] > sum {
					fourth--
				}
				if third == fourth {
					break
				}
				if nums[third]+nums[fourth] == sum {
					ans = append(ans, []int{nums[first], nums[second], nums[third], nums[fourth]})
				}
			}
		}

	}
	return ans
}
