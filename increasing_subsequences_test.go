package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// captureStdout runs f and returns everything it printed to stdout.
// Needed because AllIncreasingSubsequences prints its results instead of returning them.
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// isStrictlyIncreasing reports whether s is strictly increasing.
func isStrictlyIncreasing(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i] <= s[i-1] {
			return false
		}
	}
	return true
}

func TestAllIncreasingSubsequences(t *testing.T) {
	cases := []struct {
		name      string
		nums      []int
		wantLines int // number of subsequences printed
	}{
		{"three sorted", []int{1, 2, 3}, 7}, // every non-empty subset is increasing
		{"single", []int{5}, 1},
		{"all decreasing", []int{3, 2, 1}, 3}, // only the singletons
		{"empty", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := captureStdout(func() {
				AllIncreasingSubsequences(c.nums, []int{}, 0)
			})

			lines := 0
			for _, l := range strings.Split(strings.TrimSpace(out), "\n") {
				if strings.TrimSpace(l) != "" {
					lines++
				}
			}
			if lines != c.wantLines {
				t.Errorf("AllIncreasingSubsequences(%v) printed %d lines, want %d\noutput:\n%s",
					c.nums, lines, c.wantLines, out)
			}
		})
	}
}

func TestGetLongestIncreasingSubsequences(t *testing.T) {
	cases := []struct {
		name    string
		nums    []int
		wantLen int
	}{
		{"classic", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"mixed", []int{0, 1, 0, 3, 2, 3}, 4},
		{"all equal", []int{7, 7, 7, 7}, 1},
		{"single", []int{5}, 1},
		{"empty", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GetLongestIncreasingSubsequence(c.nums, []int{}, 0)

			if len(got) != c.wantLen {
				t.Errorf("GetLongestIncreasingSubsequences(%v) = %v (len %d), want len %d",
					c.nums, got, len(got), c.wantLen)
			}
			if !isStrictlyIncreasing(got) {
				t.Errorf("GetLongestIncreasingSubsequences(%v) = %v is not strictly increasing",
					c.nums, got)
			}
		})
	}
}

func TestGetLengthOfLIS(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{"classic", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"mixed", []int{0, 1, 0, 3, 2, 3}, 4},
		{"all equal", []int{7, 7, 7, 7}, 1},
		{"all decreasing", []int{5, 4, 3, 2, 1}, 1},
		{"all increasing", []int{1, 2, 3, 4, 5}, 5},
		{"single", []int{5}, 1},
		{"empty", []int{}, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := GetLengthOfLIS(c.nums); got != c.want {
				t.Errorf("GetLengthOfLIS(%v) = %d, want %d", c.nums, got, c.want)
			}
		})
	}
}

// TestGetAllIncreasingSubsequences is a placeholder. The current
// GetAllIncreasingSubsequences has no return value and accumulates into a
// pass-by-value [][]int, so the collected result is discarded and cannot be
// asserted. Unskip this once the function returns its [][]int result.
func TestGetAllIncreasingSubsequences(t *testing.T) {
	t.Skip("GetAllIncreasingSubsequences does not return its result yet; nothing to assert")
}
