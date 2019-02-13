package times

import (
	"time"
)

// Millisecond 获得 Unix Millisecond
func Millisecond() int {
	return int(time.Now().UnixNano() / 1000000)
}
