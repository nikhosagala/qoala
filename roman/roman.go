package roman

import (
	"errors"
	"regexp"
	"strings"
)

var romanNumber = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func IsRoman(roman string) bool {
	if roman == "" {
		return false
	}

	romanb := []byte(strings.ToUpper(roman))
	// regex got from here
	// https://stackoverflow.com/questions/267399/how-do-you-match-only-valid-roman-numerals-with-a-regular-expression
	check, _ := regexp.Match("^M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", romanb)
	return check
}

func FromString(roman string) (int, error) {
	if !IsRoman(roman) {
		return 0, errors.New("non roman number provided")
	}
	previousValue := 1000 // assign max value from roman number
	romanb := []byte(strings.ToUpper(roman))
	totalValue := 0
	for _, rDigit := range romanb {
		value := romanNumber[string(rDigit)]
		if previousValue < value {
			totalValue -= previousValue                     // Undo addition from previous step
			totalValue = totalValue + value - previousValue // Compound digit addition
		} else {
			totalValue += value
		}
		previousValue = value
	}
	return totalValue, nil
}