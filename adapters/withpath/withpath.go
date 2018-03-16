package withpath

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/nuveo/log"
)

func init() {
	log.AddAdapter("withpath", log.AdapterPod{
		Adapter: withPathLog,
		Config:  map[string]interface{}{},
	})
}

var now = time.Now

func withPathLog(m log.MsgType, o log.OutType, config map[string]interface{}, msg ...interface{}) {
	var debugInfo, lineBreak, output string
	_, fn, line, _ := runtime.Caller(5)
	fn = filepath.Base(fn)
	debugInfo = fmt.Sprintf("%s:%d ", fn, line)

	if o == log.FormattedOut {
		output = fmt.Sprintf(msg[0].(string), msg[1:]...)
	} else {
		output = fmt.Sprint(msg...)
		lineBreak = "\n"
	}

	output = fmt.Sprintf("%s%s [%s] %s%s\033[0;00m",
		log.Colors[m],
		now().UTC().Format(log.TimeFormat),
		log.Prefixes[m],
		debugInfo,
		output)

	if len(output) > log.MaxLineSize {
		output = output[:log.MaxLineSize] + "..."
	}
	output = output + lineBreak
	fmt.Print(output)
}
