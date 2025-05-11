package main

import "regexp"

func emojify(s string) string {
	replacements := map[string]string{
		"smile": "🙂",
		"grin":  "😀",
		"sad":   "😥",
		"mad":   "😠",
	}
	for word, emoji := range replacements {
		re := regexp.MustCompile(`\b` + word + `\b`)
		s = re.ReplaceAllString(s, emoji)
	}
	return s
}
