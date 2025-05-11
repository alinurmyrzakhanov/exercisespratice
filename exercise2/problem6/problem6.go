package problem6

func sumOfTwo(a, b []int, v int) bool {
	set := make(map[int]struct{}, len(a))
	for _, x := range a {
		set[x] = struct{}{}
	}
	for _, y := range b {
		if _, ok := set[v-y]; ok {
			return true
		}
	}
	return false
}
