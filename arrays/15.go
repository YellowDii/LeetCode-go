package arrays

import "sort"

// 15.三数之和
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			// 确保跟之前不重复
			continue
		}
		third := n - 1
		target := -1 * nums[first] // second+third = -first
		// 找second
		for second := first + 1; second < n; second++ {
			// 循环的过程就是second变大的过程
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				// 位置重合了 直接退出
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}
	}

	return ans
}
