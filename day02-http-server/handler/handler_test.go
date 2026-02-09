package handler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDate(t *testing.T) {
	output := GetDate()
	today := time.Now().Format("2006-01-28")
	assert.Equal(t, today, output)
}

func TestSayHello(t *testing.T) {
	assert.Equal(t, "Hello from the day02-http-server in Go.", SayHello())
}

func TestGetStatus(t *testing.T) {
	assert.Equal(t, "This is a place holder for the status.", GetStatus())
}
