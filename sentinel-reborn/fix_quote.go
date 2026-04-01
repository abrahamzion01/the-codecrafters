package main

import (
	"strings"
)

func fixSingleQuotes(text string) string {
	text = strings.Trim(text, "'")

	text = strings.TrimSpace(text)

	return "'" + text + "'"
}
