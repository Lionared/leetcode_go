package leetcode_go

import (
	"sort"
)

func JosephProblem(n, m int) int {
	winner := 0
	for i := 2; i <= n; i++ {
		winner = (winner + m) % i
	}
	winner += 1
	return winner
}

/**
@param int n: total numbers
@param int m: outer number
@param int rate: rate
*/
func JosephOuter(n, m, rate int) int {
	if rate <= 0 {
		return 0
	}
	if rate == 1 {
		return (n + m - 1) % n
	} else {
		//old_num := JosephOuter(n-1, m, rate-1)
		return (JosephOuter(n-1, m, rate-1) + m) % n
	}
}

/**
No.1. Two Sum | Easy (Array, Hash Table)
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

两数之和:
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
*/
func TwoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	tmp := make(map[int]int)
	for index, _ := range nums {
		x := nums[index]
		f := target - x
		if _, ok := tmp[f]; ok {
			ret[0] = tmp[f]
			ret[1] = index
			return ret
		}
		tmp[x] = index
	}
	return ret
}

/**
No.4. Median of Two Sorted Arrays | Hard (Array, Binary Search, Depart)
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5

两个排序数组的中位数:
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。
请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total&0x1 == 1 {
		return findM(nums1, nums2, total/2+1)
	} else {
		return (findM(nums1, nums2, total/2) + findM(nums1, nums2, total/2+1)) / 2.0
	}
}

func findM(A, B []int, k int) float64 {
	if len(A) > len(B) {
		return findM(B, A, k)
	}
	if len(A) == 0 {
		return float64(B[k-1])
	}
	if k == 1 {
		return float64(min(A[0], B[0]))
	}
	pa := min(k/2, len(A))
	pb := k - pa
	if A[pa-1] < B[pb-1] {
		return findM(A[pa:], B, k-pa)
	} else if A[pa-1] > B[pb-1] {
		return findM(A, B[pb:], k-pb)
	} else {
		return float64(A[pa-1])
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
No.11. Container With Most Water | Medium (Array, Two Pointers)
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are
drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a
container, such that the container contains the most water.
Note: You may not slant the container and n is at least 2.

给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。画 n 条垂直线，使得垂直线 i 的两个端点分别为 (i, ai) 和
(i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
注意：你不能倾斜容器，n 至少是2。
*/
func MaxArea(height []int) int {
	left, right, area := 0, len(height)-1, 0
	for left < right {
		area = max(area, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

/**
15. 3Sum | Medium (Array, Two Pointers)
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets
in the array which gives the sum of zero.

Note:
The solution set must not contain duplicate triplets.

Example:
Given array nums = [-1, 0, 1, 2, -1, -4],
A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	start := len(nums) / 2
	fixed := nums[start]
	target := 0 - fixed
	left, right := start, start
	if nums[left] > 0 {
		left--
	}
	if nums[right] < 0 {
		right++
	}
	for left > 0 && right < len(nums) {
		target++
	}
	return [][]int{}
}

/*
No.27. Remove Element | Easy (Array, Two Pointers)
Given an array nums and a value val, remove all instances of that value in-place and return the new length.
Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:
Given nums = [3,2,2,3], val = 3,
Your function should return length = 2, with the first two elements of nums being 2.
It doesn't matter what you leave beyond the returned length.

移除元素:
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/
func RemoveElement(nums []int, val int) int {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}
