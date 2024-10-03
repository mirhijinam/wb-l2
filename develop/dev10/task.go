package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/* Утилита telnet
 *
 * Реализовать примитивный telnet клиент.
 *
 * Примеры вызова:
 * go-telnet --timeout=10s host port
 * go-telnet mysite.ru 8080
 * go-telnet --timeout=3s 1.1.1.1 123
 *
 * Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
 * После подключения STDIN программы должен записываться в сокет,
 * а данные полученные из сокета должны выводиться в STDOUT.
 * Опционально в программу можно передать таймаут на подключение к серверу
 * (через аргумент --timeout, по умолчанию 10s).
 *
 * При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
 * Если сокет закрывается со стороны сервера, программа должна также завершаться.
 * При подключении к несуществующему сервер, программа должна завершаться через timeout.
 */

var (
	timeout = flag.Duration("timeout", 10*time.Second, "таймаут на подключение к серверу")
)

func telnet(host string, port string, timeout time.Duration) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), timeout)
	if err != nil {
		fmt.Println("connection error:", err)
		return
	}
	fmt.Println("connected to", host, ":", port)

	ctx, cancel := context.WithCancel(context.Background())

	// stdin->conn
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("end of input.")
				} else {
					fmt.Println("error reading from stdin:", err)
				}
				break
			}

			_, err = conn.Write([]byte(line))
			if err != nil {
				fmt.Println("error writing to server:", err)
				break
			}
		}

		cancel()
		fmt.Println("writer goroutine finished")
	}()

	// conn->stdout
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil && err != io.EOF {
			fmt.Println("error reading from server:", err)
		}

		cancel()
		fmt.Println("reader goroutine finished")
	}()

	<-ctx.Done()

	conn.Close()
	fmt.Println("connection closed")
}

func main() {
	flag.Parse()

	// Проверка аргументов
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("usage: go run task.go [--timeout=<duration>] <host> <port>")
		return
	}

	host, port := args[0], args[1]

	fmt.Printf("attempting to connect to %s:%s with timeout %v\n", host, port, *timeout)

	telnet(host, port, *timeout)
}
