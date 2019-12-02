package log

import (
	"log"
	"os"
	"path/filepath"
	"time"

	gos "github.com/golangaccount/go-libs/os"
)

func init() {
	logTime = time.Now()
	logFile = getFile(getPath(logTime))
	if logFile != nil {
		log.SetOutput(logFile)
	}
}

func getPath(tm time.Time) string {
	return filepath.Join(LOGPATH, tm.Format(DateFormat)+"."+LOGSUFFIX)
}

func setLogFile() {
	rwLock.Lock()
	defer rwLock.Unlock()
	tm := time.Now()
	if logFile == nil || logTime.Day() != tm.Day() {
		logTime = tm
		logFile = getFile(getPath(logTime))
		if logFile != nil {
			log.SetOutput(logFile)
		}
	}
}

func getFile(path string) *os.File {
	path, err := filepath.Abs(path)
	if err != nil && ISPANIC {
		panic(err)
	}
	f, err := gos.Append(path)
	if err != nil && ISPANIC {
		panic(err)
	}
	return f
}
