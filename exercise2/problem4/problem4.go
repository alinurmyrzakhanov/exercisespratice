package problem4

import "strings"

func mapping(letters []string) map[string]string {
	m := make(map[string]string, len(letters))
	for _, l := range letters {
		if len(l) == 1 {
			m[l] = strings.ToUpper(l)
		}
	}
	return m
}
