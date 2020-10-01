package algorithm

import "testing"

var sortNumbers []int = []int{1, 20, 18, 12, 12, 21, 109, 22, 211, 29}

func TestBubbleSortFromSort (t *testing.T) {
	sorted := BubbleSortFromSort(sortNumbers)

	if IsSorted(sorted) {
		t.Log("sort success by BubbleSortFromSort\n")
		return
	}
	t.Errorf("sort fail by BubbleSortFromSort\n, input is %v\n, output is %v\n", sortNumbers, sorted)
}

func TestSelectSortFromSort(t *testing.T) {
	sorted := SelectSortFromSort(sortNumbers)

	if IsSorted(sorted) {
		t.Log("sort success by SelectSortFromSort\n")
		return
	}
	t.Errorf("sort fail by SelectSortFromSort\n, input is %v\n, output is %v\n", sortNumbers, sorted)
}

func TestQuickSortFromSort(t *testing.T) {
	sorted := QuickSortFromSort(sortNumbers)

	if IsSorted(sorted) {
		t.Log("sort success by QuickSortFromSort\n")
		return
	}
	t.Errorf("sort fail by QuickSortFromSort\n, input is %v\n, output is %v\n", sortNumbers, sorted)
}
