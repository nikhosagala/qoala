package sorting

func InsertionSort(numbers []int) []int {
	// video explanation could be found here https://www.youtube.com/watch?v=OGzPmgsI-pQ
	for i := 0; i < len(numbers); i++ {
		tmp := numbers[i]
		j := i
		for j > 0 && numbers[j-1] > tmp {
			numbers[j] = numbers[j-1]
			j--
		}
		numbers[j] = tmp
	}
	return numbers
}
