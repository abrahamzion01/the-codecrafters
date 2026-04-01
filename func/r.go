package main

import (
	// "context"
	"fmt"
	"os"

	// "regexp"
	"strings"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file", err)
	}
	return string(data)
}

func writeFile(filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0664)
	if err != nil {
		fmt.Println("Error writing file", err)
	}
}

func fixAToAn(text string) string {
	words := strings.Fields(text)
	consonantSound := []string{
		"university", "unicorn", "use", "user", "unit", "euro", "one",
	}
	vowelSound := []string{
		"hour", "honor", "heir", "honest",
	}

	for i, word := range words {
		if word == "a" || word == "A" {
			if i+1 < len(words) {
				next := strings.ToLower(words[i+1])
				first := next[0]
				isVowelStart := strings.ContainsRune("aeiou", rune(first))
				for _, cs := range consonantSound {
					if strings.HasPrefix(next, cs) {
						isVowelStart = false
						break
					}
				}
				for _, vs := range vowelSound {
					if strings.HasPrefix(next, vs) {
						isVowelStart = true
						break
					}
				}
				if isVowelStart {
					if word == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}

