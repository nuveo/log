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
type LogInterface map[string]interface{}

const (
	MessageLog  msgType = 0
	Message2Log msgType = 1
	WarningLog  msgType = 2
	DebugLog    msgType = 3
	ErrorLog    msgType = 4

	ColorReset = "\033[0;00m" // Return terminal to default color

	DefaultLogFormatter = "{{.Color}}{{.Time}} [{{.Prefix}}] {{.File}}:{{.Line}} {{.Message}}{{.ColorReset}}\n"
)

var (
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

	LogHandler func(logData LogInterface) (string, LogInterface)
)

func init() {
	if LogHandler == nil {
		LogHandler = DefaultLogHandler
	}
}

func DefaultLogHandler(logData LogInterface) (string, LogInterface) {

	logData["Time"] = logData["Time"].(time.Time).UTC().Format("2006/01/02 15:04:05")
	logData["File"] = filepath.Base(logData["File"].(string))

	return DefaultLogFormatter, logData
}

func DefaultDebugLogHandler(logData LogInterface) (string, LogInterface) {

	logData["Time"] = logData["Time"].(time.Time).UTC().Format("2006/01/02 15:04:05")

	return DefaultLogFormatter, logData
}

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

	_, file, line, _ := runtime.Caller(2)

	logBody := LogInterface{
		"Color":      Colors[m],
		"ColorReset": ColorReset,
		"File":       file,
		"Line":       line,
		"Message":    fmt.Sprint(msg...),
		"MsgType":    m,
		"Prefix":     Prefixes[m],
		"Time":       time.Now(),
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
