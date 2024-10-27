package main

import (
	"fmt"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	err := os.Setenv("ARMANDO_HTTP_PORT", "9086")
	if err != nil {
		t.Fatal(err)
	}

	v := getEnv("ARMANDO_HTTP_PORT", "8080")
	if v != "9086" {
		t.Fatal(
			fmt.Sprintf("value is not correct, expected: %s, received: %s",
				"9086", v))
	}

	v = getEnv("ARMANDO_HTTP_HOST", "localhost")
	if v != "localhost" {
		t.Fatal(
			fmt.Sprintf("value is not correct, expected: %s, received: %s",
				"9086", v))
	}
}
