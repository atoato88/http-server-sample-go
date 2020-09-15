package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestInfo(t *testing.T) {
	c := exec.Command("curl", "http://localhost:8080/info")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		t.Error("error occurs when call /info")
	}
}

func TestHello(t *testing.T) {
	c := exec.Command("curl", "http://localhost:8080/hello")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		t.Error("error occurs when call /hello")
	}
}

func TestPing(t *testing.T) {
	c := exec.Command("curl", "http://localhost:8080/ping")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		t.Error("error occurs when call /hello")
	}
}
