package main

import "testing"

func TestExpressionMatter(t *testing.T) {
	tables := []struct {
		a      int
		b      int
		c      int
		result int
	}{
		{2, 1, 2, 6},
		{2, 1, 1, 4},
		{1, 1, 1, 3},
		{1, 2, 3, 9},
		{1, 3, 1, 5},
		{2, 2, 2, 8},
		{5, 1, 3, 20},
		{3, 5, 7, 105},
		{5, 6, 1, 35},
		{1, 6, 1, 8},
		{2, 6, 1, 14},
		{6, 7, 1, 48},
		{2, 10, 3, 60},
		{1, 8, 3, 27},
		{9, 7, 2, 126},
		{1, 1, 10, 20},
		{9, 1, 1, 18},
		{10, 5, 6, 300},
		{1, 10, 1, 12},
	}
	for _, table := range tables {
		result := ExpressionMatter(table.a, table.b, table.c)
		if result != table.result {
			t.Errorf("Incorrect value: (%v %v %v) was incorrect, got: %v, want: %v.", table.a, table.b, table.c, result, table.result)
		}
	}
}
