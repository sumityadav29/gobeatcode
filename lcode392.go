package main

func isSubsequence(s string, t string) bool {
	return isSubsequenceFromIndex(s, t, 0, 0)
}

func isSubsequenceFromIndex(s string, t string, si int, ti int) bool {
	if si >= len(s) {
		return true
	}

	if ti >= len(t) {
		return false
	}

	if s[si] == t[ti] {
		return isSubsequenceFromIndex(s, t, si+1, ti+1)
	}

	return isSubsequenceFromIndex(s, t, si, ti+1)
}

// Go strings are immutable byte slices
