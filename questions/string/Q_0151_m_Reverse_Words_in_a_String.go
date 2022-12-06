package main

func main() {

}

func reverseWords(s string) string {
	idx := len(s) - 1

	bytes := make([]byte, 0)
	for idx >= 0 {
		//skip leading blank
		for idx >= 0 && s[idx] == ' ' {
			idx--
		}

		end := idx
		for idx >= 0 && s[idx] != ' ' {
			idx--
		}

		for i := idx + 1; i <= end; i++ {
			bytes = append(bytes, s[i])
		}

		//check if need a blank between words,
		//if we meet another non blank char, then there is another word
		for idx >= 0 && s[idx] == ' ' {
			idx--
		}
		//not reach the end, so there is another word
		if idx >= 0 {
			bytes = append(bytes, ' ')
		}
	}

	return string(bytes)
}
