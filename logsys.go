package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type msgType uint8

// DebugMode Enable debug mode
var DebugMode bool

const (
	messageConst  msgType = 0
	message2Const msgType = 1
	warningConst  msgType = 2
	debugConst    msgType = 3
	errorConst    msgType = 0x80
)

/*
TODO: Create a more flexible HTTP error return as an example below
func HTTPError(w http.ResponseWriter, code int) {
	msg := http.StatusText(code)
	pln(errorConst, msg)
	m := make(map[string]string)
	m["status"] = "error"
	m["error"] = msg
	b, _ := json.MarshalIndent(m, "", "\t")
	http.Error(w, string(b), code)
}
*/

// Fatal show message with line break at the end and exit to OS.
func Fatal(msg ...interface{}) {
	pln(errorConst, msg...)
	os.Exit(-1)
}

// Errorln message with line break at the end.
func Errorln(msg ...interface{}) {
	pln(errorConst, msg...)
}

// Println shows message on screen with line break at the end.
func Println(msg ...interface{}) {
	pln(messageConst, msg...)
}

// Debugln shows debug message on screen with line break at the end.
// If debug mode is not active no message is displayed
func Debugln(msg ...interface{}) {
	pln(debugConst, msg...)
}

func pln(m msgType, msg ...interface{}) {
	if m == debugConst && !DebugMode {
		return
	}

	var debugInfo string
	var color string
	var prefix string

	switch m {
	case messageConst:
		color = "\x1b[37m" // White
		prefix = "msg"
	case message2Const:
		color = "\x1b[92m" // Light green
		prefix = "msg"
	case warningConst:
		color = "\x1b[93m" // Light Yellow
		prefix = "warning"
	case debugConst:
		color = "\x1b[96m" // Light Cyan
		prefix = "debug"
	case errorConst:
		color = "\x1b[91m" // Light Red
		prefix = "error"
	}

	if DebugMode {
		_, fn, line, _ := runtime.Caller(2)
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
