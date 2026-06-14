package main

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{"two longest", []int{1, 3, 5, 4, 7}, 2},      // [1,3,4,7] and [1,3,5,7]
		{"all equal", []int{2, 2, 2, 2, 2}, 5},        // LIS length 1, five of them
		{"strictly increasing", []int{1, 2, 3, 4, 5}, 1},
		{"strictly decreasing", []int{5, 4, 3, 2, 1}, 5}, // each singleton
		{"single", []int{1}, 1},
		{"two with tie", []int{1, 2, 4, 3, 5, 4, 7, 2}, 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := findNumberOfLIS(c.nums); got != c.want {
				t.Errorf("findNumberOfLIS(%v) = %d, want %d", c.nums, got, c.want)
			}
		})
	}
}
