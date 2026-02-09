package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDate(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/date", nil)
	w := httptest.NewRecorder()
	GetDate(w, req)
	today := time.Now().Format("2006-01-28")
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, today, string(body))
}

func TestSayHello(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/hello", nil)
	w := httptest.NewRecorder()
	SayHello(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Hello from the day02-http-server in Go.", string(body))
}

func TestGetStatus(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/status", nil)
	w := httptest.NewRecorder()
	GetStatus(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "This is a place holder for the status.", string(body))
}
