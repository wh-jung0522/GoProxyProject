package main

import (
	"log"
	"net/http"
	"proxy/protocol"
)

func main() {
	http.HandleFunc("/", protocol.ApiHandler)
	log.Fatal(http.ListenAndServe(":8887", nil))
}
