package armory

import (
	"time"
)

type times struct{}

var Time *times

// Millisecond 获得 Unix Millisecond
func (t *times) Millisecond() int {
	return int(time.Now().UnixNano() / 1000000)
}
