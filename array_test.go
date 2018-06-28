package leetcode_go

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	//nums := []int{2, 7, 11, 15}
	nums := []int{5, 7, 2, 4, 3, 6, 1}
	target := 5
	if found := twoSum(nums, target); found[0] == 2 && found[1] == 4 {
		t.Log("twoSum passed!")
	} else {
		t.Error("twoSum failed!")
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	//nums1 := []int{1,3}
	//nums2 := []int{2}
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{2, 3, 4, 5}
	if median := findMedianSortedArrays(nums1, nums2); median == 3 {
		t.Log("findMedianSortedArrays passed!")
	} else {
		t.Error("findMedianSortedArrays failed!")
	}
}

func Test_RemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	if newLength := RemoveElement(nums, val); newLength != 5 {
		t.Error("RemoveElement failed!")
	} else {
		t.Log("RemoveElement pass!")
	}
}
