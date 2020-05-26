package main

import "testing"


func TestAdd (t *testing.T){
	if actual := Add(1, 0); actual != 1 {
		t.Errorf("Add(%d, %d) expected %d, Actual %d", 1, 0, 1, actual)
	}
	if actual := Add(2, 3); actual != 5 {
		t.Errorf("Add(%d, %d) expected %d, Actual %d", 2, 3, 5, actual)
	}
}

func TestSum(t *testing.T) {
	type testCase  struct {
		input int
		expected int
	}
	testCases := []testCase{
		{1,1},
		{2,3},
		{10,55},
		{20,210},
	}
	for _, tcase := range testCases {
		if actual := Sum(tcase.input) ; actual != tcase.expected {
			t.Errorf("Add(%d) expected %d, Actual %d", tcase.input, tcase.expected,  actual)
		}
	}

}