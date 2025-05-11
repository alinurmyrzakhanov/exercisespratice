package main

func highestDigit(n int) int {
	if n < 0 {
		n = -n
	}
	max := 0
	if n == 0 {
		return 0
	}
	for n > 0 {
		d := n % 10
		if d > max {
			max = d
		}
		n /= 10
	}
	return max
}
