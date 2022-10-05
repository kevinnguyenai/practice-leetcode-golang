package isSubsequences

func isSubsequence(s string, t string) bool {
	m := len(s)
	n := len(t)
	val := isSubsequences(s, t, m, n)
	return val
}

func isSubsequences(s string, t string, m int, n int) bool {
	if m == 0 {
		return true
	}
	if n == 0 {
		return false
	}
	if s[m-1] == t[n-1] {
		return isSubsequences(s, t, m-1, n-1)
	}
	return isSubsequences(s, t, m, n-1)
}
