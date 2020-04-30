package palindrome

import (
	"strconv"
)

func IsPalindrome(n int64) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.FormatInt(n, 10)
	bound := (len(s) / 2) + 1
	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func FromArray(numbers []int64) []int64 {
	var results []int64
	for _, number := range numbers {
		if check := IsPalindrome(number); check {
			results = append(results, number)
		}
	}
	return results
}
