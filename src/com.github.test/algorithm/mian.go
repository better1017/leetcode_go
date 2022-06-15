package main

import (
	"fmt"
	"sort"
)

func main() {
	/*nums := []int{3, 2, 4}
	taget := 6
	sum := twoSum(nums, taget)
	fmt.Println(sum)*/
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode { //206.反转链表
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
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
func twoSum(nums []int, target int) []int { //1.两数之和
	if len(nums) == 2 {
		return []int{0, 1}
	}

	hashTable := map[int]int{}
	for idx, val := range nums {
		if idx0, ok := hashTable[target-val]; ok {
			return []int{idx0, idx}
		}
		hashTable[val] = idx
	}

	return nil
}
func climbStairs(n int) int { //爬楼梯
	if n <= 3 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 3; i <= n; i++ {
		f3 = f2 + f1
		f1, f2 = f2, f3
	}
	return f3
}
func maxArea(height []int) int { //最大雨水
	max, i, j := -1, 0, len(height)-1
	for i < j {
		println("i=", i, ",j=", j)
		actualHeight := 0
		if height[i] < height[j] {
			actualHeight = height[i]
			i++
		} else {
			actualHeight = height[j]
			j--
		}
		actualArea := actualHeight * (j - i + 1)
		if actualArea > max {
			max = actualArea
		}
		println("max=", max)
	}
	return max
}
func moveZeroes(nums []int) { //移动零
	n := len(nums)
	if n <= 1 || nums == nil {
		return
	}

	for i, j := 0, 0; i < n; i++ {
		if nums[i] != 0 {
			if i == j {
				j++
				continue
			}
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j++
		}
	}
}
