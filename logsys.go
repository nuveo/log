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

type LogInterface map[string]interface{}

const DefaultLogFormatter = "{{.Color}}{{.Time}} [{{.Prefix}}] {{.File}}:{{.Line}} {{.Message}}{{.ColorReset}}\n"

func DefaultLogHandler(logData LogInterface) (string, LogInterface) {

	logData["Time"] = logData["Time"].(time.Time).UTC().Format("2006/01/02 15:04:05")
	logData["File"] = filepath.Base(logData["File"].(string))

	return DefaultLogFormatter, logData
}

func DefaultDebugLogHandler(logData LogInterface) (string, LogInterface) {

	logData["Time"] = logData["Time"].(time.Time).UTC().Format("2006/01/02 15:04:05")

	return DefaultLogFormatter, logData
}

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

	LogHandler func(logData LogInterface) (string, LogInterface)
)

func init() {
	LogHandler = DefaultLogHandler
}

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

	var line int
	var file string
	_, file, line, _ = runtime.Caller(2)

	logBody := LogInterface{
		"Color":      color,
		"Time":       time.Now(),
		"Prefix":     prefix,
		"Message":    fmt.Sprint(msg...),
		"ColorReset": ColorReset,
		"Line":       line,
		"File":       file,
	}

	var logFormater string
	logFormater, logBody = LogHandler(logBody)

	t, err := template.New("log").Parse(logFormater)
	if err != nil {
		panic(err)
	}

	buff := new(bytes.Buffer)
	err = t.Execute(buff, logBody)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(buff.Bytes()))
}
