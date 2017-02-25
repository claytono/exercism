package slice

const testVersion = 1

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}

	return UnsafeFirst(n, s), true
}

func All(n int, s string) []string {
	all := []string{}
	for i := 0; i <= len(s)-n; i++ {
		all = append(all, s[i:i+n])
	}

	return all
}
