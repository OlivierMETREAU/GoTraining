package handler

import "time"

func GetDate() string {
	return time.Now().Format("2006-01-28")
}

func SayHello() string {
	return "Hello from the day02-http-server in Go."
}

func GetStatus() string {
	return "This is a place holder for the status."
}
