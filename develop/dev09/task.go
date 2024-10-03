package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

/* Утилита wget
 *
 * Реализовать утилиту wget с возможностью скачивать сайты целиком.
 *
 * Программа должна проходить все тесты.
 * Код должен проходить проверки go vet и golint.
 */

var (
	background = flag.Bool("b", false, "переводит загрузку в фоновый режим")
	outputFile = flag.String("o", "wget-log.txt", "имя файла, в который будет сохранен результат")
)

func getResponse(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return resp
}

func wget(url string, b bool, o string) {
	if b {
		go func() {
			resp := getResponse(url)
			defer resp.Body.Close()
			writeToFile(resp, o)
		}()
	}

	resp := getResponse(url)
	defer resp.Body.Close()
	writeToFile(resp, o)

}

func writeToFile(resp *http.Response, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	io.Copy(writer, resp.Body)

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Использование: wget [-b] [-o filename] URL")
		os.Exit(1)
	}

	url := flag.Arg(0)

	wget(url, *background, *outputFile)

	if *background {
		fmt.Println("Загрузка запущена в фоновом режиме")
	} else {
		fmt.Println("Загрузка завершена")
	}
}
