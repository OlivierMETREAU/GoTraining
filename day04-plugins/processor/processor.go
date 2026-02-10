package processor

import (
	"strings"
)

type Processor interface {
	Process(s string) string
}

type Uppercase struct{}

func (Uppercase) Process(s string) string {
	return strings.ToUpper(s)
}

type Lowercase struct{}

func (Lowercase) Process(s string) string {
	return strings.ToLower(s)
}

type Rot13 struct{}

func (Rot13) Process(s string) string {
	out := make([]rune, len(s))
	for i, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			out[i] = 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			out[i] = 'a' + (r-'a'+13)%26
		default:
			out[i] = r
		}
	}
	return string(out)
}
