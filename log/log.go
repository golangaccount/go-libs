package log

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"
)

//Debug Debug
func Debug(info string) {
	write(DEBUG, info)
}

//Debugf Debugf
func Debugf(format string, parms ...interface{}) {
	write(DEBUG, fmt.Sprintf(format, parms...))
}

//Info info
func Info(info string) {
	write(INFO, info)
}

//Infof info format
func Infof(format string, parms ...interface{}) {
	write(INFO, fmt.Sprintf(format, parms...))
}

//Warn Warn
func Warn(info string) {
	write(WARN, info)
}

//Warnf Warnf
func Warnf(format string, parms ...interface{}) {
	write(WARN, fmt.Sprintf(format, parms...))
}

//Error error
func Error(info string) {
	write(ERROR, info)
}

//Errorf errorf
func Errorf(format string, parms ...interface{}) {
	write(ERROR, fmt.Sprintf(format, parms...))
}

//Fatal fatal
func Fatal(info string) {
	write(FATAL, info)
}

//Fatalf fatalf
func Fatalf(format string, parms ...interface{}) {
	write(FATAL, fmt.Sprintf(format, parms...))
}

//Recover recover
func Recover() {
	if err := recover(); err != nil {
		write(FATAL, fmt.Sprint(err))
	}
}

//RecoverStack recoverStack
func RecoverStack() {
	if err := recover(); err != nil {
		write(FATAL, strings.Replace(string(debug.Stack()), "\n", ";", -1))
	}
}

func write(l Level, str string) {
	if l < LOGLEVEL {
		return
	}
	setLogFile()
	log.Output(3, levelStr[l]+str)
}
