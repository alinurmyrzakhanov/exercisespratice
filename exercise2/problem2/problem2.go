package problem2

import "strings"

func capitalize(names []string) []string {
	out := make([]string, len(names))
	for i, name := range names {
		low := strings.ToLower(name)
		if len(low) > 0 {
			out[i] = strings.ToUpper(low[:1]) + low[1:]
		} else {
			out[i] = low
		}
	}
	return out
}
