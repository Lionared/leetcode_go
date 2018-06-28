package leetcode_go

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	if found := twoSum(nums, target); found[0] == 2 && found[1] == 7 {
		t.Log("twoSum Test Passed!")
	} else {
		t.Error("two Sum Test Failed!")
	}
}

func Test_RemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	if newLength := RemoveElement(nums, val); newLength != 5 {
		t.Error("RemoveElement Test Failed!")
	} else {
		t.Log("RemoveElement Test Pass!")
	}
}
