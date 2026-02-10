package apacheloganalyzer

import (
	"regexp"
	"strconv"
	"time"

	sharedtypes "example.com/day03-log-analyzer/sharedtypes"
)

type ApacheLogAnalyzer struct {
	regex *regexp.Regexp
}

func New() ApacheLogAnalyzer {
	return ApacheLogAnalyzer{
		regex: regexp.MustCompile(`^(?P<remote_host>\S+) (?P<remote_logname>\S+) (?P<remote_user>[\S ]+) \[(?P<datetime>[^\]]+)\] \"(?P<method>[A-Z\-]+) (?P<request_uri>[^ \"]+) (?P<protocol>HTTP/[0-9.]+|-)\" (?P<status>[0-9]{3}) (?P<size>[0-9]+|-) "(?P<referer>[^\"]*)" "(?P<user_agent>[^\"]*)"`),
	}
}

func (ala ApacheLogAnalyzer) DecodeLine(l string) sharedtypes.AccessLine {

	result := make(map[string]string)
	match := ala.regex.FindStringSubmatch(l)
	if len(match) > 1 {
		for i, name := range ala.regex.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
	}

	d, _ := time.ParseInLocation("02/Jan/2006:15:04:05 -0700", result["datetime"], time.UTC)
	status, _ := strconv.Atoi(result["status"])
	size, _ := strconv.Atoi(result["size"])
	return sharedtypes.AccessLine{
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
