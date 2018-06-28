package leetcode_go

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	if newLength := RemoveElement(nums, val); newLength != 5 {
		t.Error("RemoveElement Test Failed!")
	} else {
		t.Log("RemoveElement Test Pass!")
	}
}
