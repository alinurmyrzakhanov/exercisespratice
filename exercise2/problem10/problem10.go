package problem10

func brandFactory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)
	makeBrand := func(name string) func(int) {
		if _, ok := brands[name]; !ok {
			brands[name] = 0
		}
		return func(delta int) {
			brands[name] += delta
		}
	}
	return brands, makeBrand
}
