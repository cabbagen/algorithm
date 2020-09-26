package algorithm

// 简单查找 -> n
func SampleFindFromLookup(numbers []int, target int) (int, bool) {
	for index, value := range numbers {
		if value == target {
			return index, true 
		}
	}
	return -1, false
}

// 二分法查找 -> log(n)
func DichotomyFromLookup(numbers []int, target, start, end int) (int, bool)  {
	halfIndex := (end - start) / 2 + start

	if numbers[halfIndex] == target {
		return halfIndex, true
	}
	if numbers[halfIndex] < target {
		return DichotomyFromLookup(numbers, target, halfIndex, end)
	}
	if numbers[halfIndex] > target {
		return DichotomyFromLookup(numbers, target, start, halfIndex)
	}

	return -1, false
}

