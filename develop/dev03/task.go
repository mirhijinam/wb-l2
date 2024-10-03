package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* Утилита sort.
 * Отсортировать строки (man sort).
 *
 * Основное.
 * Поддержать флаги:
 * -k — указание колонки для сортировки
 * -n — сортировать по числовому значению
 * -r — сортировать в обратном порядке
 * -u — не выводить повторяющиеся строки
 *
 * Дополнительное.
 * Поддержать флаги:
 * -M — сортировать по названию месяца
 * -b — игнорировать хвостовые пробелы
 * -c — проверять отсортированы ли данные
 * -h — сортировать по числовому значению с учётом суффиксов
 *
 * Программа должна проходить все тесты.
 * Код должен проходить проверки go vet и golint.
 */

var (
	column  = flag.Int("k", 1, "указание колонки дя сортировки")
	numeric = flag.Bool("n", false, "сортировать по числовому значению")
	reverse = flag.Bool("r", false, "сортировать в обратном порядке")
	unique  = flag.Bool("u", false, "не выводить повторяющиеся строки")
)

func sortRows(rows []string, k int, n, r bool) []string {
	sort.Slice(rows, func(i, j int) bool {
		fieldsI := strings.Fields(rows[i])
		fieldsJ := strings.Fields(rows[j])

		valueI := ""
		valueJ := ""

		if len(fieldsI) > k {
			valueI = fieldsI[k]
		}
		if len(fieldsJ) > k {
			valueJ = fieldsJ[k]
		}

		if n {
			numI, errI := strconv.Atoi(valueI)
			numJ, errJ := strconv.Atoi(valueJ)

			if errI != nil && errJ != nil { // Обе нечисловые – сортируем как строки
				if r {
					return valueI > valueJ
				}
				return valueI < valueJ

			} else if errI != nil { // Первая нечисловая – идет первым
				return !r

			} else if errJ != nil { // Вторая нечисловая – идет первым
				return r
			}

			if r {
				return numI > numJ
			}

			return numI < numJ
		}

		if r {
			return valueI > valueJ
		}

		return valueI < valueJ
	})

	return rows
}

func removeDuplicates(rows []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, row := range rows {
		if !seen[row] {
			seen[row] = true
			result = append(result, row)
		}
	}
	return result
}

func mySort(data []byte, k int, n, r, u bool) (string, error) {
	rows := strings.Split(string(data), "\n")

	var nonEmptyRows []string
	for _, row := range rows {
		if len(strings.TrimSpace(row)) > 0 {
			nonEmptyRows = append(nonEmptyRows, row)
		}
	}

	if k > 0 {
		k--
		nonEmptyRows = sortRows(nonEmptyRows, k, n, r)

	} else if r {
		sort.Sort(sort.Reverse(sort.StringSlice(nonEmptyRows)))

	} else {
		sort.Strings(nonEmptyRows)
	}

	if u {
		nonEmptyRows = removeDuplicates(nonEmptyRows)
	}

	return strings.Join(nonEmptyRows, "\n") + "\n", nil
}

func main() {
	flag.Parse()

	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	result, err := mySort(data, *column, *numeric, *reverse, *unique)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
