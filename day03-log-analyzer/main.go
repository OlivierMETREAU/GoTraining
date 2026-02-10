package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

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
	generateAsciiHistogram(la)
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

func generateAsciiHistogram(la accesslog.LogAnalyzer) {
	hours := make([]int, 0, len(la.Counters["Hours"]))
	var histo = map[int]int{}
	max := 0
	for h := range la.Counters["Hours"] {
		intHour, _ := strconv.Atoi(h)
		histo[intHour] = la.Counters["Hours"][h]
		hours = append(hours, intHour)
		if la.Counters["Hours"][h] > max {
			max = la.Counters["Hours"][h]
		}
	}
	sort.Ints(hours)
	divider := max / 10

	for h := range histo {
		histo[h] /= divider
	}
	max = 10

	// Print histogram from top to bottom
	for level := max; level > 0; level-- {
		for _, h := range hours {
			if histo[h] >= level {
				fmt.Print(" █ ") // ASCII-safe block alternative: "#"
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

	// Print hour labels
	for _, h := range hours {
		fmt.Printf("%2d ", h)
	}
	fmt.Println()
}
