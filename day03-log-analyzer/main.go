package main

import (
	"fmt"
	"os"

	apacheloganalyzer "example.com/day03-log-analyzer/apacheloganalyzer"
	accesslog "example.com/day03-log-analyzer/loganalyzer"
)

func main() {
	la := accesslog.New(apacheloganalyzer.New(), os.Args[1])
	la.AnalyzeFile()
	fmt.Println(la.Counters["Status"])
}
