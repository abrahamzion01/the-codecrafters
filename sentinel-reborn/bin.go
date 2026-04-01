package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BinTodecimal(bin string) string {

	words := strings.Fields(bin)
	for i, j := range words {
		if j == "(bin)" && i > 0 {
			result, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				fmt.Println("Error")
				continue
			}
			words[i-1] = strconv.Itoa(int(result))
			words[i] = ""
		}

	}
	return strings.TrimSpace(strings.Join(strings.Fields(strings.Join(words, " ")), " "))

}
