package main

import (
	"log"
	"net/http"

	"example.com/day02-http-server/handler"
)

func main() {
	http.HandleFunc("/hello", handler.SayHello)
	http.HandleFunc("/status", handler.GetStatus)
	http.HandleFunc("/date", handler.GetDate)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
