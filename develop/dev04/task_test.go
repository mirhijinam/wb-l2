package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		data     []string
		expected map[string][]string
	}{
		{
			data:     []string{"мяу", "пятак", "пятка", "тяпка", "листок", "Слиток", "слиток", "столик", "сТолик", "ток", "кот", "окт", "кто"},
			expected: map[string][]string{"пятак": {"пятак", "пятка", "тяпка"}, "листок": {"листок", "слиток", "столик"}, "ток": {"кот", "кто", "окт", "ток"}},
		},
	}

	for _, test := range tests {
		output := findAnagrams(test.data)
		assert.Equal(t, test.expected, output)
	}
}
