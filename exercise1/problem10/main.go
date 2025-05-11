package main

import (
	"fmt"
	"strconv"
)

func sum(a, b string) (string, error) {
	x, err := strconv.Atoi(a)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", a)
	}
	y, err := strconv.Atoi(b)
	if err != nil {
		return "", fmt.Errorf("string: %s cannot be converted", b)
	}
	return strconv.Itoa(x + y), nil
}
