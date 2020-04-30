package palindrome_test

import (
	"qoala/palindrome"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		number int64
		result bool
	}{
		{101, true},
		{808, true},
		{809, false},
		{111, true},
		{102, false},
	}

	for _, testcase := range testcases {
		if isPalindrome := palindrome.IsPalindrome(testcase.number); isPalindrome != testcase.result {
			t.Errorf("expected result from %d is %v but returned %v", testcase.number, testcase.result, isPalindrome)
		}
	}
}

func TestFromArray(t *testing.T) {
	testcases := []struct {
		numbers []int64
		result  int
	}{
		{[]int64{808, 809, 111, 102}, 2},
		{[]int64{908, 809, 122, 102}, 0},
	}

	for _, testcase := range testcases {
		if palindromes := palindrome.FromArray(testcase.numbers); len(palindromes) != testcase.result {
			t.Errorf("expected result from %d is %v but returned %v", testcase.numbers, testcase.result, palindromes)
		}
	}
}
