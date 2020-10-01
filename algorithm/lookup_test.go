package algorithm

import "testing"

var lookupNumbers []int = []int{ 1, 18, 19, 20, 21, 109, 111, 200, 302 }

func TestSampleFindFromLookup(t *testing.T) {
	number1, isExist1 := SampleFindFromLookup(lookupNumbers, 20)

	if number1 != 3 || !isExist1 {
		t.Errorf("lookup fail by SampleFindFromLookup\n input: %v\n output: %d, %v", lookupNumbers, number1, isExist1)
	}

	number2, isExist2 := SampleFindFromLookup(lookupNumbers, 25)

	if number2 != -1 || isExist2 {
		t.Errorf("lookup fail by SampleFindFromLookup\n input: %v\n output: %d, %v", lookupNumbers, number2, isExist2)
	}
}

func TestDichotomyFromLookup(t *testing.T) {
	number1, isExist1 := DichotomyFromLookup(lookupNumbers, 20, 0, len(lookupNumbers))

	if number1 != 3 || !isExist1 {
		t.Errorf("lookup fail by DichotomyFromLookup\n input: %v\n output: %d, %v", lookupNumbers, number1, isExist1)
	}

	number2, isExist2 := DichotomyFromLookup(lookupNumbers, 25, 0, len(lookupNumbers))

	if number2 != -1 || isExist2 {
		t.Errorf("lookup fail by DichotomyFromLookup\n input: %v\n output: %d, %v", lookupNumbers, number2, isExist2)
	}
}