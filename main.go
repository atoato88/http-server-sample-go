package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/xyproto/simpleredis"
)

var (
	port      *string
	host      *string
	redisPort *string
	redisHost *string
	redisPool *simpleredis.ConnectionPool
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	info := HandleError(redisPool.Get(0).Do("INFO")).([]byte)
	w.Write(info)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	pong := []byte("pong")
	w.Write(pong)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Hello %s", r.RemoteAddr)
	w.Write([]byte(response))
}

func HandleError(result interface{}, err error) (r interface{}) {
	if err != nil {
		panic(err)
	}
	return result
}

func init() {
	port = flag.String("port", "8080", "listen port for serving.")
	host = flag.String("host", "", "listen hostname or IP address. If not present, use localhost.")
	redisPort = flag.String("redis-port", "6379", "port for redis server.")
	redisHost = flag.String("redis-host", "redis", "hostname or IP address for redis server.")
}

func main() {
	flag.Parse()

	redisAddr := fmt.Sprintf("%v:%v", *redisHost, *redisPort)
	redisPool = simpleredis.NewConnectionPoolHost(redisAddr)
	defer redisPool.Close()

	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	listenAddr := fmt.Sprintf("%v:%v", *host, *port)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
