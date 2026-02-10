package loganalyzer

import (
	"testing"

	apacheloganalyzer "example.com/day03-log-analyzer/apacheloganalyzer"
	"github.com/stretchr/testify/assert"
)

func TestAnalyzeFile(t *testing.T) {
	la := New(apacheloganalyzer.New(), "./../apache.log")
	la.AnalyzeFile()
	assert.Equal(t, 6, la.Counters["IP"]["105.235.130.196"])
}
