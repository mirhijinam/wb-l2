package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/* Задача на распаковку.
 * Создать Go функцию, осуществляющую примитивную распаковку строки,
 * содержащую повторяющиеся символы / руны.
 *
 * Пример:
 * - "a4bc2d5e" => "aaaabccddddde"
 * - "abcd" => "abcd"
 * - "45" => "" (некорректная строка)
 * - "" => ""
 *
 * Дополнительное задание: поддержка escape - последовательностей
 * - qwe\4\5 => qwe45 (*)
 * - qwe\45 => qwe44444 (*)
 * - qwe\\5 => qwe\\\\\ (*)
 *
 * В случае если была передана некорректная строка функция должна
 * возвращать ошибку.
 * Написать unit-тесты.
 *
 * Функция должна проходить все тесты.
 * Код должен проходить проверки go vet и golint.
 *
 * Примечание: при решении считаю, что строка состоит только из
 * букв латинского алфавита, чисел и обратного слэша.
 */

var isdig = unicode.IsDigit
var islet = unicode.IsLetter

func atoi(r rune) int {
	d, err := strconv.Atoi(string(r))
	if err != nil {
		return -1
	}

	return d
}

func repeat(r rune, n int) []rune {
	res := make([]rune, 0)

	for i := 0; i < n; i++ {
		res = append(res, r)
	}

	return res
}

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "\"\"", nil
	}

	workstr := []rune(str)

	res := make([]rune, 0, len(str))

	for i := 0; i < len(workstr); {
		if workstr[i] == '\\' {
			i++

			if i == len(workstr) {
				return "\"\"", errors.New("incorrect string")
			}

			if i+1 == len(workstr) ||
				islet(workstr[i+1]) ||
				workstr[i+1] == '\\' {
				res = append(res, workstr[i])
			} else {
				res = append(res, repeat(workstr[i], atoi(workstr[i+1]))...)
				i++
			}
		}

		if islet(workstr[i]) {
			if i == len(workstr) {
				return "\"\"", errors.New("incorrect string")
			}

			if i == len(workstr)-1 || islet(workstr[i+1]) || workstr[i+1] == '\\' {
				res = append(res, workstr[i])
			} else {
				res = append(res, repeat(workstr[i], atoi(workstr[i+1]))...)
				i++
			}
		}

		i++
		if i < len(workstr) && isdig(rune(str[i])) {
			return "\"\"", errors.New("incorrect string")
		}

	}

	return string(res), nil
}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(Unpack(input))
}
