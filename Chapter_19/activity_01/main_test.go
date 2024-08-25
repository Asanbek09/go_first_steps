package main

import (
	"testing"

	"gotest.tools/assert"
)

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		inputs []int
		want   int
	}{
		{
			name:   "Test Case 1",
			inputs: []int{5, 6},
			want:   11,
		}, {
			name:   "Test Case 2",
			inputs: []int{11, 7},
			want:   18,
		}, {
			name:   "Test Case 3",
			inputs: []int{1, 8},
			want:   9,
		}, {
			name:   "Test Case 4 (international failure)",
			inputs: []int{2, 3},
			want:   0, //This should be 5, intentionally incorrect to demonstrate failure
		},
	}
	//for _, test := range tests {
	//	got := add(test.inputs[0], test.inputs[1])
	//	assert.Equal(t, test.want, got)
	//}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got := add(test.inputs[0], test.inputs[1])
			assert.Equal(t, test.want, got)
		})
	}
}
