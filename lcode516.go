package main

func longestPalindromeSubseq(s string) int {

	lcs := make([][]int, len(s)+1)

	for i := range lcs {
		lcs[i] = make([]int, len(s)+1)
	}

	for i := 1; i <= len(s); i++ {
		lcs[i] = make([]int, len(s)+1)
		for j := 1; j <= len(s); j++ {
			if s[i-1] == s[len(s)-j] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	return lcs[len(lcs)-1][len(lcs[0])-1]
}
