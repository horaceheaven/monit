package main

import (
	"time"
)

var StartTime time.Time

func init() {
	StartTime = time.Now()
}

func UpTime() time.Duration {
	return time.Since(StartTime)
}
