package peggy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	type test struct {
		test     []int
		expected []int
	}
	tests := []test{
		{
			test: []int{
				4, 30, 50,
			},
			expected: []int{
				12, 1,
			},
		},
		{
			test: []int{
				4, 17, 50,
			},
			expected: []int{
				-1, -1,
			},
		},
		{
			test: []int{
				4, 30, 50, 80, 100,
			},
			expected: []int{
				-1, -1,
			},
		},
		{
			test: []int{
				4, 17,
			},
			expected: []int{
				26, 3,
			},
		},
		{
			test: []int{
				18, 40, 100, 150,
			},
			expected: []int{
				8, 1,
			},
		},
		{
			test: []int{
				20, 40, 60, 100,
			},
			expected: []int{
				-1, -1,
			},
		},
		{
			test: []int{
				20, 40, 60, 90,
			},
			expected: []int{
				-1, -1,
			},
		},
		{
			test: []int{
				20, 40, 60, 80,
			},
			expected: []int{
				40, 3,
			},
		},
	}

	for _, test := range tests {
		answer := Calculate(test.test)
		require.Equal(t, test.expected, answer)
	}
}
