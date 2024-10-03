package main

import "testing"

func TestUnpack(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{`a4bc2d5e`, `aaaabccddddde`},
		{`abcd`, `abcd`},
		{`45`, `""`},
		{``, `""`},
		{`qwe\4\5`, `qwe45`},
		{`qwe\\5`, `qwe\\\\\`},
		{`qwe\4\5rty`, `qwe45rty`},
		{`qwe\45rty`, `qwe44444rty`},
	}

	for _, test := range cases {
		res, err := Unpack(test.input)
		if err != nil {
			if test.expected != "\"\"" {
				t.Errorf("input:\"%s\"; unexpected error:\"%v\"", test.input, err)
			}
		}
		if res != test.expected {
			t.Errorf("input:\"%s\"; expected:\"%s\"; got:\"%s\"", test.input, test.expected, res)
		}
	}
}
