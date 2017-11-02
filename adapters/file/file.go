package file

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/nuveo/log"
)

var now = time.Now

func init() {
	log.SetAdapter("file", log.AdapterPod{
		Adapter: fileWrite,
		Config:  map[string]string{"fileName": "file.log"},
	})
}

func fileWrite(m log.MsgType, o log.OutType, config map[string]string, msg ...interface{}) {
	if m == log.DebugLog && !log.DebugMode {
		return
	}

	var debugInfo, lineBreak, output string

	if log.DebugMode {
		_, fn, line, _ := runtime.Caller(5)
		fn = filepath.Base(fn)
		debugInfo = fmt.Sprintf("%s:%d ", fn, line)
	}

	if o == log.FormattedOut {
		output = fmt.Sprintf(msg[0].(string), msg[1:]...)
	} else {
		output = fmt.Sprint(msg...)
		lineBreak = "\n"
	}

	output = fmt.Sprintf("%s [%s] %s%s",
		now().UTC().Format(log.TimeFormat),
		log.Prefixes[m],
		debugInfo,
		output)

	if len(output) > log.MaxLineSize {
		output = output[:log.MaxLineSize] + "..."
	}
	output = output + lineBreak
	fmt.Print(output)

	f, err := os.OpenFile(config["fileName"], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(output); err != nil {
		panic(err)
	}
}
