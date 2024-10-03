package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCut struct {
	f, d           string
	s              bool
	data           []byte
	expectedOutput []byte
	expectedError  error
}

var testCuts = []testCut{
	{
		f:              "1,2,3",
		d:              " ",
		s:              false,
		data:           []byte("First_Test_First_Line\nFirst test Second Line\nFirstTestThirdLine"),
		expectedOutput: []byte("First_Test_First_Line\nFirst test Second\nFirstTestThirdLine\n"),
		expectedError:  nil,
	},
	{
		f:              "1,2",
		d:              " ",
		s:              false,
		data:           []byte("Second_Test_First_Line\nSecond Test Second Line\nSecond_Test Third Line"),
		expectedOutput: []byte("Second_Test_First_Line\nSecond Test\nSecond_Test Third\n"),
		expectedError:  nil,
	},
	{
		f:              "3,4",
		d:              "_",
		s:              true,
		data:           []byte("Third_Test_First_Line\nThird Test Second Line\nThirdTestThird Line"),
		expectedOutput: []byte("First_Line\n"),
		expectedError:  nil,
	},
}

func TestCut(t *testing.T) {
	for _, test := range testCuts {
		output, err := cut(test.data, test.f, test.d, test.s)
		if test.expectedError != nil {
			assert.Error(t, err)
			assert.Equal(t, test.expectedError.Error(), err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expectedOutput, output)
		}
	}
}
