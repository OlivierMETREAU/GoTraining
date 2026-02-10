package loganalyzer

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strconv"

	sharedtypes "example.com/day03-log-analyzer/sharedtypes"
)

type LineDecoder interface {
	DecodeLine(string) (sharedtypes.AccessLine, error)
}

type LogAnalyzer struct {
	decoder          LineDecoder
	filePath         string
	Counters         map[string]map[string]int
	totalSize        int
	numberOfRequests int
}

func New(ld LineDecoder, fp string) LogAnalyzer {
	return LogAnalyzer{
		decoder:  ld,
		filePath: fp,
		Counters: map[string]map[string]int{
			"IP":         {},
			"UserID":     {},
			"Method":     {},
			"RequestUri": {},
			"Protocol":   {},
			"Status":     {},
			"Hours":      {},
		},
		totalSize:        0,
		numberOfRequests: 0,
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
		al, err := la.decoder.DecodeLine(line)
		if err == nil {
			la.incrementCounters(al)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}

func (la *LogAnalyzer) incrementCounters(al sharedtypes.AccessLine) {
	la.Counters["IP"][al.IP] += 1
	la.Counters["UserID"][al.UserID] += 1
	la.Counters["Method"][al.Method] += 1
	la.Counters["RequestUri"][al.RequestUri] += 1
	la.Counters["Protocol"][al.Protocol] += 1
	la.Counters["Status"][http.StatusText(al.Status)] += 1
	la.Counters["Hours"][strconv.Itoa(al.DateTime.Hour())] += 1
	la.totalSize += al.Size
	la.numberOfRequests += 1
}
