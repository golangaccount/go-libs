package log

import (
	"os"
	"sync"
	"time"
)

//Level log日志等级
type Level int

const (
	//DEBUG 调试
	DEBUG Level = iota
	//INFO 一般信息
	INFO
	//WARN 警告
	WARN
	//ERROR 错误
	ERROR
	//FATAL 致命
	FATAL
)

var (
	levelStr = []string{"[DEBUG] ", "[INFO] ", "[WARN] ", "[ERROR] ", "[FATAL] "}
	//LOGLEVEL 写入log文件中的log日志级别
	LOGLEVEL Level = INFO
	//DateFormat 日期格式
	DateFormat string = "2006-01-02"
	//LOGPATH 日志地址
	LOGPATH string = "logs"
	//LOGSUFFIX log文件后缀
	LOGSUFFIX string = "log"
	//ISPANIC 日志文件创建失败，写失败是否进行panic
	ISPANIC bool = false
)

var logFile *os.File
var logTime time.Time
var rwLock sync.RWMutex
