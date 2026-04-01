package main

import (
	"strings"
	"strconv"
)

func Upper(words []string) []string {
	for i := 0; i < len(words); i++ {

		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--

		} else if strings.HasPrefix(words[i], "(up,") && i > 0 && i+1 < len(words) {

			numStr := strings.TrimSuffix(words[i+1], ")")
			n, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}

			for j := 1; j <= n && i-j >= 0; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}

			words = append(words[:i], words[i+2:]...)
			i--
		}
	}
	return words
}