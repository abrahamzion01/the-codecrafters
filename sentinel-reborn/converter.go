package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func ConvertNumbers(s string) string {
	reHex := regexp.MustCompile(`(-?[0-9A-Fa-f]+)\s*\(hex\)`)

	s = reHex.ReplaceAllStringFunc(s, func(m string) string {
		parts := reHex.FindStringSubmatch(m)
		if len(parts) < 2 {
			return m
		}
		n, err := strconv.ParseInt(parts[1], 16, 64)
		if err != nil {
			return m
		}
		return fmt.Sprint(n)
	})
	return s
}
