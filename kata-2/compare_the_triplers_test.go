package main

import (
	"testing"
)

func TestCompareTriplets(t *testing.T) {
	tables := []struct {
		a      []int32
		b      []int32
		result []int32
	}{
		{[]int32{5, 6, 7}, []int32{3, 6, 10}, []int32{1, 1}},
		{[]int32{1, 2, 3}, []int32{1, 2, 3}, []int32{0, 0}},
		{[]int32{6, 8, 12}, []int32{7, 9, 15}, []int32{0, 3}},
		{[]int32{10, 15, 20}, []int32{5, 6, 7}, []int32{3, 0}},
	}

	for _, table := range tables {
		compare := CompareTriplets(table.a, table.b)
		if compare[0] != table.result[0] || compare[1] != table.result[1] {
			t.Errorf("Incorrect value: (%v %v) was incorrect, got: %v, want: %v.", table.a, table.b, compare, table.result)
		}
	}
}
