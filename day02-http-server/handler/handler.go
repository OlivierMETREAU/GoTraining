package handler

import (
	"fmt"
	"net/http"
	"time"
)

func GetDate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("2006-01-28"))
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a place holder for the status.")
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from the day02-http-server in Go.")
}
