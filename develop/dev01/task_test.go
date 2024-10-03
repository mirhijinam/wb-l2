package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestGetTimeSuccess(t *testing.T) {
	ntpTimeFunc = func(_ string) (time.Time, error) {
		return time.Date(2024, time.September, 20, 12, 34, 56, 0, time.UTC), nil
	}

	expected := "2024-09-20 12:34:56 +0000 UTC"
	result, err := getTime()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Fatalf("error: %s", fmt.Errorf("expected: %s, got: %s", expected, result))
	}
}

func TestGetTimeError(t *testing.T) {
	ntpTimeFunc = func(_ string) (time.Time, error) {
		return time.Time{}, errors.New("NTP server unreachable")
	}

	_, err := getTime()
	if err == nil {
		t.Fatalf("unexpected success!")
	}

	expectedError := "NTP server unreachable"
	if err.Error() != expectedError {
		t.Fatalf("error: %s", fmt.Errorf("expected: %s, got: %s", expectedError, err.Error()))
	}
}
