package one

import "sort"

// 15. 三数之和

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
// 注意：答案中不可以包含重复的三元组。

// 示例 1：
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]

// 示例 2：
//输入：nums = []
//输出：[]

// 示例 3：
//输入：nums = [0]
//输出：[]

// 提示：
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵

func threeSum(nums []int) [][]int { //15.三数之和
	// 排序+夹逼
	n := len(nums)
	var res [][]int
	if n < 3 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
			} else if sum > 0 {
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}

	return res
}
