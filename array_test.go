package leetcode_go

import (
	"testing"
	"fmt"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3,2,2,3}
	val := 3
	newLength := RemoveElement(nums,val)
	fmt.Println(newLength)
}