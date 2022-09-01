package search

import "testing"

type testStruct struct {
	data int
}

func (t testStruct) CompareTo(other testStruct) int {
	return other.data - t.data
}

func TestComparableSlice(t *testing.T) {
	tests := []struct {
		input    []testStruct
		element  testStruct
		expected int
	}{
		{
			[]testStruct{
				{1},
				{2},
				{3},
				{4},
				{5},
				{6},
			},
			testStruct{3},
			2,
		},
		{
			[]testStruct{
				{10},
				{20},
				{30},
				{40},
				{50},
				{60},
			},
			testStruct{30},
			2,
		},
		{
			[]testStruct{
				{10},
				{20},
				{30},
				{40},
				{50},
				{60},
			},
			testStruct{70},
			-1,
		},
	}

	for _, test := range tests {
		if actual := ComparableSlice(test.input, test.element); actual != test.expected {
			t.Errorf("expected index %d, but got %d, for slice %v, with element %d", test.expected, actual, test.input, test.element)
		}
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		input    []int
		element  int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 8},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 9, -1},
	}

	for _, test := range tests {
		if actual := Slice(test.input, test.element); actual != test.expected {
			t.Errorf("expected index %d, but got %d, for slice %v, with element %d", test.expected, actual, test.input, test.element)
		}
	}
}
