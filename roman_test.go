package roman_test

import (
	roman "qoala"
	"testing"
)

func TestIsRoman(t *testing.T) {
	testcases := []struct {
		roman  string
		result bool
	}{
		{"IJKL", false},
		{"IIII", false},
		{"IVI", false},
		{"VIV", false},
		{"IXIII", false},
		{"CHECK", false},
		{"III", true},
		{"IV", true},
		{"IX", true},
		{"MCMXCIV", true},
		{"IIV", false},
	}
	for _, testcase := range testcases {
		if check := roman.IsRoman(testcase.roman); check != testcase.result {
			t.Errorf("expected result from %s is %v but returned %v", testcase.roman, testcase.result, check)
		}
	}
}

func TestFromString(t *testing.T) {
	testcases := []struct {
		roman  string
		result int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, testcase := range testcases {
		if result, err := roman.FromString(testcase.roman); result != testcase.result {
			t.Errorf("expected result from %s is %v but returned %v, %s", testcase.roman, testcase.result, result, err)
		}
	}
}
