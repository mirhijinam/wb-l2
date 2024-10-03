package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/* Взаимодействие с ОС.
 *
 * Необходимо реализовать собственный шелл:
 * - встроенные команды: cd, pwd, echo, kill, ps
 * - поддержать fork, exec команды
 * - конвеер на пайпах
 *
 * Реализовать утилиту netcat (nc) клиент:
 * - принимать данные из stdin и отправлять в соединение (tcp/udp)
 *
 * Программа должна проходить все тесты.
 * Код должен проходить проверки go vet и golint.
 */

func main() {
	runShell()
}

func runShell() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		executeCommand(input)
	}
}

func executeCommand(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		changeDirectory(args)
	case "pwd":
		printWorkingDirectory()
	case "echo":
		echo(args[1:])
	case "kill":
		killProcess(args)
	case "ps":
		listProcesses()
	case "nc":
		netcat(args)
	case "fork":
		runFork(args)
	case "exec":
		runExec(args)
	default:
		fmt.Println("unknown command")
	}
}

func changeDirectory(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: cd <directory>")
		return
	}

	err := os.Chdir(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(dir)
	}
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func killProcess(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: kill <pid>")
		return
	}

	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	err = process.Kill()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func listProcesses() {
	processes, err := os.ReadDir("/proc")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, process := range processes {
		fmt.Println(process.Name())
	}
}

func netcat(args []string) {
	if len(args) < 3 {
		fmt.Println("usage: nc <host> <port>")
		return
	}
	host, port := args[1], args[2]

	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer conn.Close()

	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}

func runFork(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: fork <command>")
		return
	}

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		fmt.Printf("failed to start command: %v\n", err)
		return
	}

	fmt.Printf("process started with pid: %d\n", cmd.Process.Pid)

	go func() {
		err := cmd.Wait()
		if err != nil {
			fmt.Printf("command finished with error: %v\n", err)
		}
	}()
}

func runExec(args []string) {
	if len(args) < 2 {
		fmt.Println("usage: exec <command>")
		return
	}

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("failed to run command: %v\n", err)
	}
}
