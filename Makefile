BIN:=http-server-sample
TAG:=latest

.PHONY: all build run clean test
all: run

build:
	go build -o $(BIN) ./main.go 

build-container:
	docker build -t $(BIN):$(TAG) . 

run:
	go run ./main.go

clean:
	rm $(BIN)

test:
	go test . -v -cover

integration-test:
	docker-compose up -d
	go test ./test/integration -v 
	docker-compose down
