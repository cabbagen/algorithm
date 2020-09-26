package algorithm

// 选择排序 ->  n * n
func SelectSortFromSort(numbers []int) []int {
	for i, _ := range numbers {
		for j, _ := range numbers {
			if numbers[i] < numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

// 快速排序 ->  n * log(n)
func QuickSortFromSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	var greater, smalls []int

	for _, value := range numbers {
		if value < numbers[0] {
			smalls = append(smalls, value)
		}
		if value > numbers[0] {
			greater = append(greater, value)
		}
	}
	return append(append(QuickSortFromSort(smalls), numbers[0]), QuickSortFromSort(greater)...)
}
