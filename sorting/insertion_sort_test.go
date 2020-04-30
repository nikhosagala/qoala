package sorting_test

import (
	"qoala/sorting"
	"sort"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testcases := []struct {
		numbers []int
	}{
		{[]int{10, 2, 1, 5, 3, 6}},
		{[]int{200, 300, 20, 3, 1, 100, 45, 82}},
	}

	for _, testcase := range testcases {
		if result := sorting.InsertionSort(testcase.numbers); !sort.IntsAreSorted(result) {
			t.Errorf("ints are not sorted yet %v", result)
		}
	}
}
