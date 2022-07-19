package search_test

import (
	"testing"

	"github.com/snprajwal/godsalgo/search"
)

type testSearch struct {
	arr      []int64
	key      int64
	expected int
}

var tests []testSearch = []testSearch{
	{
		arr:      []int64{1, 2, 3, 4, 5, 6, 7, 8},
		key:      3,
		expected: 2,
	},
	{
		arr:      []int64{1, 2, 3, 4, 5, 6, 7, 8},
		key:      6,
		expected: 5,
	},
	{
		arr:      []int64{1, 2, 3, 4, 5, 6, 7, 8},
		key:      9,
		expected: -1,
	},
}

func TestLinear(t *testing.T) {
	for _, test := range tests {
		got := search.Linear(test.arr, test.key)
		if got != test.expected {
			t.Errorf("Got: %d\nExpected: %d", got, test.expected)
		}
	}
}

func TestBinary(t *testing.T) {
	for _, test := range tests {
		got := search.Binary(test.arr, test.key)
		if got != test.expected {
			t.Errorf("Got: %d\nExpected: %d", got, test.expected)
		}
	}
}

func TestBinaryRecusive(t *testing.T) {
	for _, test := range tests {
		got := search.BinaryRecursive(test.arr, test.key)
		if got != test.expected {
			t.Errorf("Got: %d\nExpected: %d", got, test.expected)
		}
	}
}
