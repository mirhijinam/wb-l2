package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySort(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		n        bool
		r        bool
		u        bool
		data     []byte
		expected string
	}{
		{
			name:     "Сортировка по умолчанию",
			k:        0,
			n:        false,
			r:        false,
			u:        false,
			data:     []byte("dddd ebcd 21234\naaaa pfor 9874\nssss odod 2921\nuuuu vfjv 8212\nbbbb olws 6714\noooo ikik 9120\nnnnn yhji 2912\n"),
			expected: "aaaa pfor 9874\nbbbb olws 6714\ndddd ebcd 21234\nnnnn yhji 2912\noooo ikik 9120\nssss odod 2921\nuuuu vfjv 8212\n",
		},
		{
			name:     "Числовая сортировка по первой колонке",
			k:        1,
			n:        true,
			r:        false,
			u:        false,
			data:     []byte("10 dddd ebcd 12343\n100 aaaa pfor 9874\n2 ssss odod 2921\n32 nnnn yhji 2912\n5 uuuu vfjv 8212\n50 bbbb olws 6714\n6 oooo ikik 9120\n"),
			expected: "2 ssss odod 2921\n5 uuuu vfjv 8212\n6 oooo ikik 9120\n10 dddd ebcd 12343\n32 nnnn yhji 2912\n50 bbbb olws 6714\n100 aaaa pfor 9874\n",
		},
		{
			name:     "Обратная сортировка",
			k:        0,
			n:        false,
			r:        true,
			u:        false,
			data:     []byte("dddd ebcd 21234\naaaa pfor 9874\nssss odod 2921\nuuuu vfjv 8212\nbbbb olws 6714\noooo ikik 9120\nnnnn yhji 2912\n"),
			expected: "uuuu vfjv 8212\nssss odod 2921\noooo ikik 9120\nnnnn yhji 2912\ndddd ebcd 21234\nbbbb olws 6714\naaaa pfor 9874\n",
		},
		{
			name:     "Сортировка по четвертой колонке",
			k:        4,
			n:        false,
			r:        false,
			u:        false,
			data:     []byte("dddd ebcd 21234 5\naaaa pfor 9874 1\nssss odod 2921 3\nuuuu vfjv 8212 4\nbbbb olws 6714 2\noooo ikik 9120 6\nnnnn yhji 2912 7\n"),
			expected: "aaaa pfor 9874 1\nbbbb olws 6714 2\nssss odod 2921 3\nuuuu vfjv 8212 4\ndddd ebcd 21234 5\noooo ikik 9120 6\nnnnn yhji 2912 7\n",
		},
		{
			name:     "Числовая сортировка по четвертой колонке",
			k:        4,
			n:        true,
			r:        false,
			u:        false,
			data:     []byte("dddd ebcd 21234 5\naaaa pfor 9874 1\nssss odod 2921 3\nuuuu vfjv 8212 4\nbbbb olws 6714 2\noooo ikik 9120 6\nnnnn yhji 2912 7\n"),
			expected: "aaaa pfor 9874 1\nbbbb olws 6714 2\nssss odod 2921 3\nuuuu vfjv 8212 4\ndddd ebcd 21234 5\noooo ikik 9120 6\nnnnn yhji 2912 7\n",
		},
		{
			name:     "Удаление дубликатов",
			k:        0,
			n:        false,
			r:        false,
			u:        true,
			data:     []byte("a\nb\na\nc\nb\n"),
			expected: "a\nb\nc\n",
		},
		{
			name:     "Сортировка с отсутствующей колонкой",
			k:        3,
			n:        false,
			r:        false,
			u:        false,
			data:     []byte("a b c\nd e\nf g h\n"),
			expected: "d e\na b c\nf g h\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := mySort(test.data, test.k, test.n, test.r, test.u)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, output)
		})
	}
}
