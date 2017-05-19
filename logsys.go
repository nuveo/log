package log

import (
	"bytes"
	"fmt"
	"html/template"
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

type DebugInterface interface{}
type LogInterface interface{}

var (
	ColorMessage  = "\x1b[37m"   // White
	ColorMessage2 = "\x1b[92m"   // Light green
	ColorWarning  = "\x1b[93m"   // Light Yellow
	ColorDebug    = "\x1b[96m"   // Light Cyan
	ColorError    = "\x1b[91m"   // Light Red
	ColorReset    = "\033[0;00m" // Return terminal to default color

	PrefixMessage  = "msg"
	PrefixMessage2 = "msg"
	PrefixWarinig  = "warning"
	PrefixDebug    = "debug"
	PrefixError    = "error"

	DebugFormatter = "{{.File}}:{{.Line}}"
	LogFormatter   = "{{.Color}}{{.Time}} [{{.Prefix}}] {{.DebugInfo}} {{.Message}}{{.ColorReset}}\n"

	DebugBody DebugInterface
	LogBody   LogInterface

	/*
		"Color":      color,
			"Time":       time.Now().UTC().Format("2006/01/02 15:04:05"),
			"Prefix":     prefix,
			"DebugInfo":  debugInfo,
			"Message":    fmt.Sprint(msg...),
			"ColorReset": ColorReset,
	*/

	LogHandler func(LogData map[string]interface{}, Message ...string) (string, map[string]interface{}) = func(
		LogData map[string]interface{},
		Message ...string) (formater string, body map[string]interface{}) {

		return
	}
)

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
		color = ColorMessage
		prefix = PrefixMessage
	case message2Const:
		color = ColorMessage2
		prefix = PrefixMessage2
	case warningConst:
		color = ColorWarning
		prefix = PrefixWarinig
	case debugConst:
		color = ColorDebug
		prefix = PrefixDebug
	case errorConst:
		color = ColorError
		prefix = PrefixError
	}

	if DebugMode {
		_, fn, line, _ := runtime.Caller(2)
		fn = filepath.Base(fn)

		DebugBody = map[string]interface{}{
			"File": fn,
			"Line": line,
		}

		buff := new(bytes.Buffer)
		t, _ := template.New("debug").Parse(DebugFormatter)

		t.Execute(buff, DebugBody)

		debugInfo = string(buff.Bytes())

	}

	LogBody = map[string]interface{}{
		"Color":      color,
		"Time":       time.Now().UTC().Format("2006/01/02 15:04:05"),
		"Prefix":     prefix,
		"DebugInfo":  debugInfo,
		"Message":    fmt.Sprint(msg...),
		"ColorReset": ColorReset,
	}

	//runtime.Caller()

	LogHandler(LogBody, message)

	buff := new(bytes.Buffer)
	t, _ := template.New("log").Parse(LogFormatter)

	t.Execute(buff, LogBody)

	fmt.Print(string(buff.Bytes()))
}
