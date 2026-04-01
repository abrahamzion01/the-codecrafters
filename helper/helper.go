package helper

import (
	"strconv"
	"strings"
)

func Upper(s string) string {
	ssplit := strings.Fields(s)
	for i := 0; i < len(ssplit); i++ {
		if strings.HasPrefix(ssplit[i], "(up") {

			cleanedi := strings.Trim(ssplit[i], "()")
			sepi := strings.Split(cleanedi, ",")

			count := 1

			if len(sepi) > 1 {
				num, err := strconv.Atoi(strings.TrimSpace(sepi[1]))
				if err == nil {
					count = num
				}
			}
			for j := 1; j <= count; j++ {
				if i-j < 0 {

					break
				}
				ssplit[i-j] = strings.TrimSpace(ssplit[i-j])

				ssplit[i-j] = strings.ToUpper(ssplit[i-j])

			}
			ssplit = append(ssplit[:i], ssplit[i+1:]...)
			i--
		}
	}
	result := strings.Join(ssplit, " ")
	return result
}

// func FixArticle(s string) string {

// 	ssplit := strings.Fields(s)
// 	vowels := "aeiouAEIOU"
// 	for i := 1; i < len(ssplit); i++ {
// 		if strings.ContainsAny(string(ssplit[i][0]), vowels) {

// 			if string(ssplit[i-1]) == "a" {
// 				ssplit[i-1] = "an"

// 			}
// 		}

// 	}
// 	return strings.Join(ssplit, " ")

// }

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

P
