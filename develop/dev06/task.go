package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Утилита cut.
 *
 * Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные.
 *
 * Поддержать флаги:
 * -f, "fields" - выбрать поля (колонки)
 * -d, "delimiter" - использовать другой разделитель
 * -s, "separated" - только строки с разделителем
 *
 * Программа должна проходить все тесты.
 * Код должен проходить проверки go vet и golint.
 */

var (
	fields    = flag.String("f", "", "выбрать поля (колонки)")
	delimiter = flag.String("d", "\t", "использовать другой разделитель")
	separated = flag.Bool("s", false, "только строки с разделителем")
)

func cut(input []byte, fields string, delimiter string, separated bool) ([]byte, error) {
	if fields == "" {
		return nil, fmt.Errorf("не указаны поля (-f)")
	}

	fieldIndices := make([]int, 0)
	for _, field := range strings.Split(fields, ",") {
		if strings.Contains(field, "-") {
			rangeParts := strings.Split(field, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("некорректный формат диапазона: %s", field)
			}

			start, err1 := strconv.Atoi(rangeParts[0])
			end, err2 := strconv.Atoi(rangeParts[1])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("некорректный формат чисел в диапазоне: %s", field)
			}

			for i := start; i <= end; i++ {
				fieldIndices = append(fieldIndices, i)
			}
		} else {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("некорректный формат числа: %s", field)
			}

			fieldIndices = append(fieldIndices, num)
		}
	}

	lines := bytes.Split(input, []byte("\n"))
	var result []byte

	for _, line := range lines {
		if separated && !bytes.Contains(line, []byte(delimiter)) {
			continue
		}

		columns := bytes.Split(line, []byte(delimiter))
		var selectedColumns [][]byte

		for _, index := range fieldIndices {
			if index > 0 && index <= len(columns) {
				selectedColumns = append(selectedColumns, columns[index-1])
			}
		}

		if len(selectedColumns) > 0 {
			result = append(result, bytes.Join(selectedColumns, []byte(delimiter))...)
			result = append(result, '\n')
		}
	}

	return result, nil
}

func main() {
	flag.Parse()

	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при чтении файла:", err)
		os.Exit(1)
	}

	result, err := cut(data, *fields, *delimiter, *separated)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}

	fmt.Print(string(result))
}
