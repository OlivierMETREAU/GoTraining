package loganalyzer

import (
	"bufio"
	"log"
	"os"

	sharedtypes "example.com/day03-log-analyzer/sharedtypes"
)

type LineDecoder interface {
	DecodeLine(string) sharedtypes.AccessLine
}

type LogAnalyzer struct {
	decoder  LineDecoder
	filePath string
	Counters map[string]map[string]int
}

func New(ld LineDecoder, fp string) LogAnalyzer {
	return LogAnalyzer{
		decoder:  ld,
		filePath: fp,
		Counters: map[string]map[string]int{
			"IP":    {},
			"Other": {},
		},
	}
}

func (la *LogAnalyzer) AnalyzeFile() {
	file, err := os.Open(la.filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		al := la.decoder.DecodeLine(line)
		la.Counters["IP"][al.IP] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}
