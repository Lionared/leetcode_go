package leetcode_go

/**
78. Subsets | Medium ()
Given a set of distinct integers, nums, return all possible subsets (the power set).
Note: The solution set must not contain duplicate subsets.
Example:
Input: nums = [1,2,3]
Output:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
func Subsets(nums []int) [][]int {
	return [][]int{}
}

/**
191. Number of 1 Bits | Easy (Bit Manipulation)
Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

Example 1:
Input: 11
Output: 3
Explanation: Integer 11 has binary representation 00000000000000000000000000001011

Example 2:
Input: 128
Output: 1
Explanation: Integer 128 has binary representation 00000000000000000000000010000000
*/
func HammingWeight(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
