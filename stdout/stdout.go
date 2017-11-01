package stdout

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nuveo/log"
)

func init() {
	l := log.Logger{
		HTTPError: HTTPError,
		Fatal:     Fatal,
		Errorln:   Errorln,
		Errorf:    Errorf,
		Warningln: Warningln,
		Warningf:  Warningf,
		Println:   Println,
		Printf:    Printf,
		Debugln:   Debugln,
		Debugf:    Debugf,
	}

	log.Loggers = append(log.Loggers, l)
}

// HTTPError write lot to stdout and return json error on http.ResponseWriter with http error code.
func HTTPError(w http.ResponseWriter, code int) {
	msg := http.StatusText(code)
	Errorln(msg)
	m := make(map[string]string)
	m["status"] = "error"
	m["error"] = msg
	b, _ := json.MarshalIndent(m, "", "\t")
	http.Error(w, string(b), code)
}

// Fatal show message with line break at the end and exit to OS.
func Fatal(msg ...interface{}) {
	pln(log.ErrorLog, log.LineOut, msg...)
	os.Exit(-1)
}

// Errorln message with line break at the end.
func Errorln(msg ...interface{}) {
	pln(log.ErrorLog, log.LineOut, msg...)
}

// Errorf shows formatted error message on stdout without line break at the end.
func Errorf(msg ...interface{}) {
	pln(log.ErrorLog, log.FormattedOut, msg...)
}

// Warningln shows warning message on stdout with line break at the end.
func Warningln(msg ...interface{}) {
	pln(log.WarningLog, log.LineOut, msg...)
}

// Warningf shows formatted warning message on stdout without line break at the end.
func Warningf(msg ...interface{}) {
	pln(log.WarningLog, log.FormattedOut, msg...)
}

// Println shows message on stdout with line break at the end.
func Println(msg ...interface{}) {
	pln(log.MessageLog, log.LineOut, msg...)
}

// Printf shows formatted message on stdout without line break at the end.
func Printf(msg ...interface{}) {
	pln(log.MessageLog, log.FormattedOut, msg...)
}

// Debugln shows debug message on stdout with line break at the end.
// If debug mode is not active no message is displayed
func Debugln(msg ...interface{}) {
	pln(log.DebugLog, log.LineOut, msg...)
}

// Debugf shows debug message on stdout without line break at the end.
// If debug mode is not active no message is displayed
func Debugf(msg ...interface{}) {
	pln(log.DebugLog, log.FormattedOut, msg...)
}

func pln(m log.MsgType, o log.OutType, msg ...interface{}) {
	if m == log.DebugLog && !log.DebugMode {
		return
	}

	var debugInfo, lineBreak, output string

	if log.DebugMode {
		_, fn, line, _ := runtime.Caller(2)
		fn = filepath.Base(fn)
		debugInfo = fmt.Sprintf("%s:%d ", fn, line)
	}

	if o == log.FormattedOut {
		output = fmt.Sprintf(msg[0].(string), msg[1:]...)
	} else {
		output = fmt.Sprint(msg...)
		lineBreak = "\n"
	}

	output = fmt.Sprintf("%s%s [%s] %s%s\033[0;00m",
		log.Colors[m],
		log.Now().UTC().Format("2006/01/02 15:04:05"),
		log.Prefixes[m],
		debugInfo,
		output)

	if len(output) > log.MaxLineSize {
		output = output[:log.MaxLineSize] + "..."
	}
	output = output + lineBreak
	fmt.Print(output)
}
