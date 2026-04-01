package main

import (
	"strconv"
	"strings"
)

func ToCap(dswayy string) string {
	s := strings.Fields(dswayy)
	for i := 0; i < len(s); i++ {
		if s[i] == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s = append(s[:i], s[i+1:]...)
		} else if strings.HasPrefix(s[i], "(cap") {
			s[i+1] = strings.TrimSuffix(s[i+1], ")")
			n, err := strconv.Atoi(s[i+1])
			if err != nil {
				continue
			}
			for j := 1; j <= n && i-j >= 0; j++ {
				s[i-j] = strings.Title(s[i-j])
			}
			s = append(s[:i], s[i+2:]...)
			i--
		}
	}
	return strings.Join(s, " ")
}
