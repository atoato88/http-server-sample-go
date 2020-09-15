package main

import (
	"net/http"
	"os"
	"os/exec"
	"testing"
)

func TestInfo(t *testing.T) {
	c := exec.Command("curl", "--silent", "http://localhost:8080/info")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		t.Error("error occurs when call /info")
	}
}

func TestHello(t *testing.T) {
	c := exec.Command("curl", "--silent", "http://localhost:8080/hello")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		t.Error("error occurs when call /hello")
	}
}

func TestPing(t *testing.T) {
	c := exec.Command("curl", "--silent", "http://localhost:8080/ping")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		t.Error("error occurs when call /ping")
	}
}

func TestInfo2(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/info")
	if err != nil {
		t.Error("error occurs when call /info")
	}
	if resp.StatusCode != 200 {
		t.Error("status code isn't 200 when call /info")
	}
}

func TestHello2(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		t.Error("error occurs when call /hello")
	}
	if resp.StatusCode != 200 {
		t.Error("status code isn't 200 when call /hello")
	}
}

func TestPing2(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		t.Error("error occurs when call /info")
	}
	if resp.StatusCode != 200 {
		t.Error("status code isn't 200 when call /ping")
	}
}
