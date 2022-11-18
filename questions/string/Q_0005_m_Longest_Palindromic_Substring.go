package main

func main() {

}

func longestPalindrome(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		odd := palindrome(s, i, i)
		even := palindrome(s, i, i+1)

		if len(odd) > len(res) {
			res = odd
		}
		if len(even) > len(res) {
			res = even
		}
	}

	return res
}

func palindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return s[left+1 : right]
}
