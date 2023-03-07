package flights

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortFlights(t *testing.T) {
	type TestData struct {
		input    [][]string
		expected []string
	}

	testTable := []TestData{
		{
			input: [][]string{
				{"A", "B"},
			},
			expected: []string{"A", "B"},
		},
		{
			input: [][]string{
				{"A", "B"}, {"D", "A"}, {"B", "C"},
			},
			expected: []string{"D", "C"},
		},
		{
			input: [][]string{
				{"D", "A"}, {"B", "C"}, {"A", "B"},
			},
			expected: []string{"D", "C"},
		},
		{
			input: [][]string{
				{"B", "C"}, {"A", "B"}, {"D", "A"},
			},
			expected: []string{"D", "C"},
		},
	}

	assert := assert.New(t)

	for _, test := range testTable {
		fl := FlightList{Flights: test.input}
		src, dst := fl.Sorted()
		assert.Equal(test.expected[0], src)
		assert.Equal(test.expected[1], dst)
	}
}
