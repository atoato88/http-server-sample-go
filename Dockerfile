# Build the binary
FROM golang:1.13 as builder

# Copy in the go src
WORKDIR /go/src/workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o main main.go

# Copy the operator and dependencies into a thin image
FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /go/src/workspace/main .
ENTRYPOINT ["./main"]

