package main

import (
	"fmt"
	"testing"
)

func TestLinear(t *testing.T) {
	m := GameMap{
		Settings: &MapSettings{},
	}
	m.Settings.SizeHorizontal = 5
	m.Settings.SizeVertical = 7
	m.Settings.Map = []int{
		7, 1, 1, 1, 1,
		1, 5, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 2, 2, 2, 1,
		1, 1, 1, 1, 9,
	}
	tests := []struct {
		name     string
		row      int
		column   int
		expected int
	}{
		{
			name:     "simple",
			row:      1,
			column:   1,
			expected: 5,
		},
		{
			name:     "first",
			row:      0,
			column:   0,
			expected: 7,
		},
		{
			name:     "last",
			row:      6,
			column:   4,
			expected: 9,
		},
	}
	for _, test := range tests {
		if res := m.createLinearFromRowAndColumn(test.row, test.column); m.Settings.Map[res] != test.expected {
			t.Errorf(fmt.Sprintf("%s failed, expected %d, got %d, index %d", test.name, test.expected, m.Settings.Map[res], res))
		}
	}

}
