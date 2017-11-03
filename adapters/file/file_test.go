package file

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/nuveo/log"
)

func TestFileWrite(t *testing.T) {
	now = func() time.Time { return time.Unix(1498405744, 0) }

	fileWrite(
		log.ErrorLog,
		log.LineOut,
		map[string]string{"fileName": "logfile.txt"},
		"test log")
	fileWrite(
		log.DebugLog,
		log.LineOut,
		map[string]string{"fileName": "logfile.txt"},
		"test log")
	fileWrite(
		log.WarningLog,
		log.LineOut,
		map[string]string{"fileName": "logfile.txt"},
		"test log")

	b, err := ioutil.ReadFile("logfile.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = os.Remove("logfile.txt")
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	expectd := "2017/06/25 15:49:04 [error] test log\n2017/06/25 15:49:04 [warning] test log\n"
	if string(b) != expectd {
		t.Fatalf("Error expectd %q, got %q\n", expectd, string(b))
	}
}
