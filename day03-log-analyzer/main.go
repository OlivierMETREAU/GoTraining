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
	fmt.Println(la.Counters["Hours"])
	mostFreqIp := getMostFrequentIP(la)
	fmt.Printf("Most used IP is %s, is %d occurences\n", mostFreqIp, la.Counters["IP"][mostFreqIp])
}

func getMostFrequentIP(la accesslog.LogAnalyzer) string {
	var maxKey string
	var maxVal int
	first := true

	for k, v := range la.Counters["IP"] {
		if first || v > maxVal {
			maxKey = k
			maxVal = v
			first = false
		}
	}

	return maxKey
}
