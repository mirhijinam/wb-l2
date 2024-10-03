package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	defer listener.Close()
	fmt.Println("server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		conn.Write([]byte("echo:" + text + "\n"))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading from connection:", err)
	}
}
