package currenttime

import (
	"time"
)

// GetCurrentTime return a string containing the current time
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02T15:04:05.999-07:00")
}
