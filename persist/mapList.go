package persist

func ColList(validList []string) map[string]int {
	cols := make(map[string]int)
	for k, val := range validList {
		cols[val] = k
	}
	return cols
}
