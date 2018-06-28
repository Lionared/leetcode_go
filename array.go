package leetcode_go

/*
	No. 27 Remove Element | easy
	给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
	不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
	元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
	Given an array nums and a value val, remove all instances of that value in-place and return the new length.
	Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
	The order of elements can be changed. It doesn't matter what you leave beyond the new length.

	Example 1:
	Given nums = [3,2,2,3], val = 3,
	Your function should return length = 2, with the first two elements of nums being 2.
	It doesn't matter what you leave beyond the returned length.

	Example 2:
	Given nums = [0,1,2,2,3,0,4,2], val = 2,
	Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.
	Note that the order of those five elements can be arbitrary.
	It doesn't matter what values are set beyond the returned length.
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
