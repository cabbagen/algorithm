package algorithm

// 冒泡排序 -> n * n
func BubbleSortFromSort(numbers []int) []int {
	for i:= 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

// 选择排序 ->  n * n
func SelectSortFromSort(numbers []int) []int {
	var i, j, min int
	
	for i = 0; i < len(numbers); i++ {
		min = i
		
		for j = i; j < len(numbers); j++ {
			if numbers[min] > numbers[j] {
				min = j
			}
		}

		if i != min {
			numbers[i], numbers[min] = numbers[min], numbers[i]
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

func IsSorted(numbers []int) bool {
	var result bool = true

	for result && len(numbers) > 1 {
		first, second := numbers[0:2][0], numbers[0:2][1]

		result, numbers = second - first >= 0, numbers[1:]
	}

	return result
}