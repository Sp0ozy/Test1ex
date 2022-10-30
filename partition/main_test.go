package main

import "testing"

func TestSort(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		n    int
		want int
	}{
		{[]int{}, 10, 0},
		{[]int{0, 120, 21}, 1, 21},
		{[]int{1, 2, 54, 101, 4, 1}, 3, 4},
		{[]int{1, 2, 54, 101, 4, 1}, 5, 101},
		{[]int{1, 2, 54, 101, 4, 1, 5}, 8, 0},
		{[]int{1, -2, 54, -101, -4, 1, -5}, 6, 54},
		{[]int{1, 2, -54, 101, 4, 1, -5}, 1, -5},
	} {
		if got := find(tc.s, tc.n); got != tc.want {
			t.Errorf("find(%v,%v) = %v, want %v", tc.s, tc.n, got, tc.want)
		}
	}
}
