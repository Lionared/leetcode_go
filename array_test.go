package leetcode_go

import (
	"testing"
)

func TestJosephProblem(t *testing.T) {
	n, m := 9, 4
	winner := JosephProblem(n, m)
	t.Log("JosephProblem Winner is", winner)
}

func TestJosephOuter(t *testing.T) {
	n, m := 9, 4
	rate := 1
	for ; rate < n; rate++ {
		outer := JosephOuter(n, m, rate)
		t.Log("Round", rate, "outer is", outer)
	}
}

/**
1
*/
func TestTwoSum(t *testing.T) {
	//nums := []int{2, 7, 11, 15}
	nums := []int{5, 7, 2, 4, 3, 6, 1}
	target := 5
	if found := TwoSum(nums, target); found[0] == 2 && found[1] == 4 {
		t.Log("TwoSum passed!")
	} else {
		t.Error("TwoSum failed!")
	}
}

/**
4
*/
func TestFindMedianSortedArrays(t *testing.T) {
	//nums1 := []int{1,3}
	//nums2 := []int{2}
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{2, 3, 4, 5}
	if median := FindMedianSortedArrays(nums1, nums2); median == 3 {
		t.Log("FindMedianSortedArrays passed!")
	} else {
		t.Error("FindMedianSortedArrays failed!")
	}
}

/**
11.
*/
func TestMaxArea(t *testing.T) {
	height := []int{5, 7, 2, 4, 3, 6, 1}
	area := MaxArea(height)
	t.Log("MaxArea is", area)
}

/**
15.
*/
func TestThreeSum(t *testing.T) {
	nums := []int{-2, 0, 0, 2, 2}
	//nums := []int{0,0,0,0}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSum(nums)
	t.Log("ThreeSum is:", result)
}

/**
27.
*/
func TestRemoveElement(t *testing.T) {
	nums := []int{-1, -1, 0, 0, 1, 1}
	val := 2
	if newLength := RemoveElement(nums, val); newLength != 5 {
		t.Error("RemoveElement failed!")
	} else {
		t.Log("RemoveElement pass!")
	}
}
