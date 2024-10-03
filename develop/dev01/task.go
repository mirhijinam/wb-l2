package main

/* Базовая задача 
 * Создать программу печатающую точное время с использованием NTP библиотеки.
 * Инициализировать как go module.
 * Использовать библиотеку https://github.com/beevik/ntp.
 * Написать программу печатающую текущее время / точное время 
 * с использованием этой библиотеки.
 *
 * Программа должна быть оформлена с использованием как go module.
 * Программа должна корректно обрабатывать ошибки библиотеки: 
 * распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
 * Программа должна проходить проверки go vet и golint.
 */

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

var ntpTimeFunc = ntp.Time // переменная для функции, которую можно замокать в тестах

func getTime() (string, error) {
	time, err := ntpTimeFunc("pool.ntp.org")
	if err != nil {
		return "", err
	}

	return time.String(), nil
}

func main() {
	time, err := getTime()
	if err != nil {
		log.Fatalf("failed to get precise time: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("precise time: %s\n", time)
}
