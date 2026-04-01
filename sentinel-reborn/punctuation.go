package main

func FixPunctuation(s string) string {
	r := []rune(s)
	var res []rune

	isPunct := func(c rune) bool {
		return c == '.' || c == ',' || c == '!' || c == '?' || c == ':' || c == ';'
	}

	for i := 0; i < len(r); i++ {

		if isPunct(r[i]) {

			for len(res) > 0 && res[len(res)-1] == ' ' {
				res = res[:len(res)-1]
			}
			res = append(res, r[i])
			if i+1 < len(r) {

				j := i + 1
				for j < len(r) && r[j] == ' ' {
					j++
				}

				if j < len(r) && r[j] != '\'' && !isPunct(r[j]) {
					res = append(res, ' ')
				}
			}

			continue
		}

		res = append(res, r[i])
	}

	return string(res)
}
