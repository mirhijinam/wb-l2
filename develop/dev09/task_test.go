package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestWget(t *testing.T) {
	server1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test content"))
	}))
	defer server1.Close()

	t.Run("regular", func(t *testing.T) {
		outputFile := "test_output.txt"
		wget(server1.URL, false, outputFile)

		content, err := os.ReadFile(outputFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		if string(content) != "test content" {
			t.Errorf("expect '%s', but got '%s'", "test content", string(content))
		}

		// os.Remove(outputFile)
	})

	server2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		for i := 0; i < 10; i++ {
			w.Write([]byte("test content"))
		}
	}))
	defer server2.Close()
	t.Run("background", func(t *testing.T) {
		outputFile := "test_output_background.txt"
		wget(server2.URL, true, outputFile)

		time.Sleep(time.Second)

		content, err := os.ReadFile(outputFile)
		if err != nil {
			t.Fatalf("failed to read file: %v", err)
		}

		expectedContent := strings.Repeat("test content", 10)
		if string(content) != expectedContent {
			t.Errorf("expect '%s', but got '%s'", expectedContent, string(content))
		}

		// os.Remove(outputFile)
	})
}
