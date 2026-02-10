package log

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type AccessLine struct {
	IP         string
	UserID     string
	DateTime   time.Time
	Method     string
	RequestUri string
	Protocol   string
	Status     int
	Size       int
}

func DecodeLine(l string) AccessLine {
	myExp := regexp.MustCompile(`^(?P<remote_host>\S+) (?P<remote_logname>\S+) (?P<remote_user>[\S ]+) \[(?P<datetime>[^\]]+)\] \"(?P<method>[A-Z\-]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+|-)\" (?P<status>[0-9]{3}) (?P<size>[0-9]+|-) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)"`)
	result := make(map[string]string)
	match := myExp.FindStringSubmatch(l)
	if len(match) > 1 {
		for i, name := range myExp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
	}

	d, _ := time.ParseInLocation("02/Jan/2006:15:04:05 -0700", result["datetime"], time.UTC)
	status, _ := strconv.Atoi(result["status"])
	size, _ := strconv.Atoi(result["size"])
	return AccessLine{
		IP:         result["remote_host"],
		UserID:     result["remote_user"],
		DateTime:   d,
		Method:     result["method"],
		RequestUri: result["request_uri"],
		Protocol:   result["protocol"],
		Status:     status,
		Size:       size,
	}
}

func AnalyzeFile(path string) map[string]map[string]int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	counters := map[string]map[string]int{
		"IP":    {},
		"Other": {},
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		al := DecodeLine(line)
		counters["IP"][al.IP] += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return counters
}
