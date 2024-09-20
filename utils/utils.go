package utils

import "regexp"

func RemoveAnsiStyle(s string) string {
	return regexp.MustCompile(`\x1b[[\d;]*m`).ReplaceAllString(s, "")
}
