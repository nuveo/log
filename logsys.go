package logSys

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

type MsgType uint8

const (
	Message MsgType = 0
	Debug   MsgType = 1
	Error   MsgType = 0x80
)

var DebugMode = false

func Println(m MsgType, msg ...interface{}) {
	if m == Debug && !DebugMode {
		return
	}

	var debugInfo string
	var color string
	var prefix string

	switch m {
	case Message:
		color = "\x1b[37m" // White
		prefix = "msg"
	case Debug:
		color = "\x1b[93m" // Light Yellow
		prefix = "debug"
	case Error:
		color = "\x1b[91m" // Light Red
		prefix = "error"
	}

	if DebugMode {
		_, fn, line, _ := runtime.Caller(1)
		fn = filepath.Base(fn)
		debugInfo = fmt.Sprintf("%s:%d ", fn, line)
	}

	fmt.Printf("%s%s [%s] %s%s\033[0;00m\n",
		color,
		time.Now().UTC().Format("2006/01/02 15:04:05"),
		prefix,
		debugInfo,
		fmt.Sprint(msg...))
}
