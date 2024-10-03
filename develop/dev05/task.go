package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

/* Утилита grep.
 *
 * Реализовать утилиту фильтрации (man grep).
 *
 * Поддержать флаги:
 * -A, "after"       - печатать N строк после строки, совпадающей с паттерном
 * -B, "before"      - печатать N строк до строки, совпадающей с паттерном
 * -C, "context"     - печатать по N строк вокруг строки, совпадающей с паттерном
 * -c, "count"       - вывести количество строк, совпадающих с паттерном
 * -i, "ignore-case" - игнорировать регистр при соотнесении строк с паттерном
 * -v, "invert"      - печатать строки, не совпадающие с паттерном
 * -F, "fixed"       - печатать строки, точно совпадающие с паттерном
 * -n, "line-number" - печатать номер строки, совпадающей с паттерном
 *
 * Программа должна проходить все тесты.
 * Код должен проходить проверки go vet и golint.
 */

var (
	after      = flag.Int("A", 0, "печатать N строк после строки, совпадающей с паттерном")
	before     = flag.Int("B", 0, "печатать N строк до строки, совпадающей с паттерном")
	context    = flag.Int("C", 0, "печатать по N строк вокруг строки, совпадающей с паттерном")
	count      = flag.Bool("c", false, "вывести количество строк, совпадающих с паттерном")
	ignoreCase = flag.Bool("i", false, "игнорировать регистр при соотнесении строк с паттерном")
	invert     = flag.Bool("v", false, "печатать строки, не совпадающие с паттерном")
	fixed      = flag.Bool("f", false, "печатать строки, точно совпадающие с паттерном")
	lineNum    = flag.Bool("n", false, "печатать номер строки, совпадающей с паттерном")
)

func grep(
	data []byte,
	pattern string,
	A int,
	B int,
	C int,
	c bool,
	i bool,
	v bool,
	F bool,
	n bool,
) (string, error) {
	lines := strings.Split(string(data), "\n")
	result := []string{}
	matchCount := 0

	for lineNum, line := range lines {
		isMatch := false

		if i {
			isMatch = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		} else if F {
			isMatch = line == pattern
		} else {
			isMatch = strings.Contains(line, pattern)
		}

		if v {
			isMatch = !isMatch
		}

		if isMatch {
			matchCount++

			if !c {
				if n {
					result = append(result, fmt.Sprintf("%d:%s", lineNum+1, line))
				} else {
					result = append(result, line)
				}

				for i := 1; i <= A && lineNum+i < len(lines); i++ {
					result = append(result, lines[lineNum+i])
				}

				for i := 1; i <= B && lineNum-i >= 0; i++ {
					result = append([]string{lines[lineNum-i]}, result...)
				}

				if C > 0 {
					for i := 1; i <= C && lineNum+i < len(lines); i++ {
						result = append(result, lines[lineNum+i])
					}
					for i := 1; i <= C && lineNum-i >= 0; i++ {
						result = append([]string{lines[lineNum-i]}, result...)
					}
				}
			}
		}
	}

	if c {
		return fmt.Sprintf("%d", matchCount), nil
	}

	return strings.Join(result, "\n"), nil
}

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		fmt.Println("Использование: программа [флаги] паттерн файл")
		flag.PrintDefaults()
		return
	}

	pattern := flag.Arg(0)
	filename := flag.Arg(1)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения файла: %v\n", err)
		os.Exit(1)
	}

	result, err := grep(data, pattern, *after, *before, *context, *count, *ignoreCase, *invert, *fixed, *lineNum)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка выполнения grep: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
