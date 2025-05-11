package problem5

import "sort"

func products(pr map[string]float64, min float64) []string {
	type kv struct {
		key   string
		price float64
	}
	arr := make([]kv, 0)
	for k, v := range pr {
		if v >= min {
			arr = append(arr, kv{k, v})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].price == arr[j].price {
			return arr[i].key < arr[j].key
		}
		return arr[i].price > arr[j].price
	})
	res := make([]string, len(arr))
	for i, kv := range arr {
		res[i] = kv.key
	}
	return res
}
