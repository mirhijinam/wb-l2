package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type taskTest struct {
	input          []byte
	pattern        string
	after          int
	before         int
	context        int
	count          bool
	ignoreCase     bool
	invert         bool
	fixed          bool
	lineNum        bool
	expectedOutput string
	expectedError  error
}

var taskTests = []taskTest{
	{
		input:          []byte("apple\nbanana\ncherry\ndate\neggplant"),
		pattern:        "an",
		expectedOutput: "banana\neggplant",
	},
	{
		input:          []byte("Apple\nBanana\nCherry\nDate\nEggplant"),
		pattern:        "A",
		ignoreCase:     true,
		expectedOutput: "Apple\nBanana\nDate\nEggplant",
	},
	{
		input:          []byte("one\ntwo\nthree\nfour\nfive"),
		pattern:        "o",
		after:          1,
		lineNum:        true,
		expectedOutput: "1:one\ntwo\n2:two\nthree\n4:four\nfive",
	},
	{
		input:          []byte("red\ngreen\nblue\nyellow\npurple"),
		pattern:        "e",
		count:          true,
		expectedOutput: "5",
	},
	{
		input:          []byte("cat\ndog\nfish\nbird\nsnake"),
		pattern:        "a",
		invert:         true,
		expectedOutput: "dog\nfish\nbird",
	},
}

func TestTask(t *testing.T) {
	for _, test := range taskTests {
		output, err := grep(test.input, test.pattern, test.after, test.before, test.context,
			test.count, test.ignoreCase, test.invert, test.fixed, test.lineNum)
		assert.Equal(t, test.expectedError, err)
		assert.Equal(t, test.expectedOutput, output)
	}
}
