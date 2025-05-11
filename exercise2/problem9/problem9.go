package problem9

func factory(factor int) func(...int) []int {
	return func(vals ...int) []int {
		res := make([]int, len(vals))
		for i, v := range vals {
			res[i] = v * factor
		}
		return res
	}
}
