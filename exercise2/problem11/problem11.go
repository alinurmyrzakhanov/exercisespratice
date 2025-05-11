package problem11

func removeDups[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	res := make([]T, 0, len(items))
	for _, v := range items {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}
