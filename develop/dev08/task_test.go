package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestChangeDirectory(t *testing.T) {
	initialDir, _ := os.Getwd()
	changeDirectory([]string{"cd", ".."})
	newDir, _ := os.Getwd()
	if newDir == initialDir {
		t.Errorf("changeDirectory не изменил текущую директорию")
	}
	os.Chdir(initialDir)
}

func TestPrintWorkingDirectory(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printWorkingDirectory()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	if len(out) == 0 {
		t.Errorf("printWorkingDirectory ничего не вывел")
	}
}

func TestEcho(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	echo([]string{"Hello", "World"})

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	if !strings.Contains(string(out), "Hello World") {
		t.Errorf("echo не вывел ожидаемый текст")
	}
}

func TestRunExec(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	runExec([]string{"exec", "echo", "test"})

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	if !strings.Contains(string(out), "test") {
		t.Errorf("runExec не выполнил команду корректно")
	}
}

func TestExecuteCommand(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	executeCommand("echo test")

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout

	if !strings.Contains(string(out), "test") {
		t.Errorf("executeCommand не выполнил команду корректно")
	}
}
