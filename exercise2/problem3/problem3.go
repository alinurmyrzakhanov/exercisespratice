package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, corner dir) [][]int {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			var v int
			switch corner {
			case ul:
				v = i + j
			case ur:
				v = i + (n - 1 - j)
			case ll:
				v = (n - 1 - i) + j
			case lr:
				v = (n - 1 - i) + (n - 1 - j)
			default:
				v = i + j
			}
			m[i][j] = v
		}
	}
	return m
}
