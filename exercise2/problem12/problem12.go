package problem11

import "sort"

type Ordered interface {
	~int | ~string
}

func keysAndValues[K Ordered, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	vals := make([]V, 0, len(m))
	for _, k := range keys {
		vals = append(vals, m[k])
	}
	return keys, vals
}
