package leetcode_go

/**
No.1. Two Sum | Easy
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
func twoSum(nums []int, target int) []int {
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
	No.4. Median of Two Sorted Arrays | Hard
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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
	} else {
		return a
	}
}

/*
	No.27. Remove Element | Easy
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
