package sharedtypes

import "time"

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
