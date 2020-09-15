package main

import (
	"testing"
)

func Test_initDefault(t *testing.T) {
	expectedHost := ""
	expectedPort := "8080"
	expectedRedisHost := "redis"
	expectedRedisPort := "6379"

	if *host != expectedHost {
		t.Fatalf("server host isn't default value. expacted:%v actual:%v", expectedHost, *host)
	}
	if *port != expectedPort {
		t.Fatalf("server port isn't default value. expacted:%v actual:%v", expectedPort, *port)
	}
	if *redisHost != expectedRedisHost {
		t.Fatalf("redis server host isn't default value. expacted:%v actual:%v", expectedRedisHost, *redisHost)
	}
	if *redisPort != expectedRedisPort {
		t.Fatalf("redis server port isn't default value. expacted:%v actual:%v", expectedRedisPort, *redisPort)
	}
}
