package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	reader, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error Reading File:", err)
		return ""
	}
	return string(reader)

}
func writeFile(filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Println("error writing file:", err)
		return
	}
}

func binTodecimal(bin string) string {

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
