package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type msgType uint8

const (
	MessageLog  msgType = 0
	Message2Log msgType = 1
	WarningLog  msgType = 2
	DebugLog    msgType = 3
	ErrorLog    msgType = 4
)

var (

	// DebugMode Enable debug mode
	DebugMode bool

	Colors = []string{
		MessageLog:  "\x1b[37m", // White
		Message2Log: "\x1b[92m", // Light green
		WarningLog:  "\x1b[93m", // Light Yellow
		DebugLog:    "\x1b[96m", // Light Cyan
		ErrorLog:    "\x1b[91m", // Light Red
	}

	Prefixes = []string{
		MessageLog:  "msg",
		Message2Log: "msg",
		WarningLog:  "warning",
		DebugLog:    "debug",
		ErrorLog:    "error",
	}
)

// Fatal show message with line break at the end and exit to OS.
func Fatal(msg ...interface{}) {
	pln(ErrorLog, msg...)
	os.Exit(-1)
}

// Errorln message with line break at the end.
func Errorln(msg ...interface{}) {
	pln(ErrorLog, msg...)
}

// Println shows message on screen with line break at the end.
func Println(msg ...interface{}) {
	pln(MessageLog, msg...)
}

// Debugln shows debug message on screen with line break at the end.
// If debug mode is not active no message is displayed
func Debugln(msg ...interface{}) {
	pln(DebugLog, msg...)
}

func pln(m msgType, msg ...interface{}) {
	if m == DebugLog && !DebugMode {
		return
	}

	var debugInfo string

	if DebugMode {
		_, fn, line, _ := runtime.Caller(2)
		fn = filepath.Base(fn)
		debugInfo = fmt.Sprintf("%s:%d ", fn, line)
	}

	fmt.Printf("%s%s [%s] %s%s\033[0;00m\n",
		Colors[m],
		time.Now().UTC().Format("2006/01/02 15:04:05"),
		Prefixes[m],
		debugInfo,
		fmt.Sprint(msg...))
}
