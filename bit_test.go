package leetcode_go

import "testing"

func TestHammingWeight(t *testing.T) {
	n := 11
	if output := HammingWeight(n); output == 3 {
		t.Log("HammingWeight passed!")
	} else {
		t.Error("HammingWeight failed!")
	}
}
