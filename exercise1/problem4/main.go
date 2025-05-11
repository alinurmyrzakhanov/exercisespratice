package main

func detectWord(s string) string {
	result := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result += string(r)
		}
	}
	return result
}
