package apacheloganalyzer

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessLine(t *testing.T) {
	line := "83.149.9.216 - - [17/May/2015:10:05:03 +0000] \"GET /images/name.png HTTP/1.1\" 200 203023 \"http://exmaple.com/presentations/\" \"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/32.0.1700.77 Safari/537.36\""
	la := New()
	al, _ := la.DecodeLine(line)
	assert.Equal(t, "83.149.9.216", al.IP)
	assert.Equal(t, "-", al.UserID)
	assert.Equal(t, time.Date(2015, 05, 17, 10, 5, 3, 0, time.UTC), al.DateTime)
	assert.Equal(t, "GET", al.Method)
	assert.Equal(t, "/images/name.png", al.RequestUri)
	assert.Equal(t, "HTTP/1.1", al.Protocol)
	assert.Equal(t, 200, al.Status)
	assert.Equal(t, 203023, al.Size)
}
