BIN:=http-server-sample
TAG:=latest

.PHONY: all build run clean
all: run

build:
	go build -o $(BIN) ./main.go 

build-container:
	docker build -t $(BIN):$(TAG) . 

run:
	go run ./main.go

clean:
	rm $(BIN)

