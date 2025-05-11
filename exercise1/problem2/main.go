package main

import (
	"strconv"
)

func binary(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		bit := n % 2
		result = strconv.Itoa(bit) + result
		n /= 2
	}
	return result
}
