package logs

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

var logFiles []*os.File

func init() {
	go clippingLog()
}

func clippingLog() {
	for {
		var step int64 = 24 * 3600
		// step = 5
		gap := (time.Now().Unix() + 8*3600) % step
		if gap == 0 {
			for _, logFile := range logFiles {
				logPath := logFile.Name()
				fi, err := logFile.Stat()
				if fi.Size() == 0 || err != nil {
					continue
				}
				// println(logFile.Name())
				rotateFormat := "20060102"
				// rotateFormat = "20060102-150405"
				dirPath := filepath.Dir(logPath)
				baseName := filepath.Base(logPath)
				ext := filepath.Ext(logPath)
				name := strings.TrimSuffix(baseName, ext)

				fix := time.Duration(-1 * step * int64(time.Second))
				rotateFlag := time.Now().Add(fix).Format(rotateFormat)
				rotatePath := filepath.Join(dirPath, name+"-"+rotateFlag+ext)

				os.Rename(logPath, rotatePath)

				oldF := *logFile
				newF, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				*logFile = *newF
				oldF.Close()
			}
		}
		time.Sleep(time.Second)
	}
}

// DailyRotateLog DailyRotateLog
func DailyRotateLog(path *string) *os.File {
	f, _ := os.OpenFile(*path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logFiles = append(logFiles, f)
	return f
}
