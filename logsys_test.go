package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
	"time"
)

func validate(key string, logFunc func(msg ...interface{}), valueExpected string) (err error) {
	rescueStdout := os.Stdout
	defer func() { os.Stdout = rescueStdout }()

	r, w, err := os.Pipe()
	if err != nil {
		return
	}
	os.Stdout = w

	logFunc("log test")

	err = w.Close()
	if err != nil {
		return
	}

	out, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	if string(out) != valueExpected {
		err = fmt.Errorf("Error, '%s' printed %q, expected %q", key, string(out), valueExpected)
	}
	return
}

func TestLog(t *testing.T) {
	now = func() time.Time { return time.Unix(1498405744, 0) }
	DebugMode = false

	data := []struct {
		key           string
		logFunc       func(msg ...interface{})
		expectedValue string
	}{
		{"Println", Println, "\x1b[37m2017/06/25 15:49:04 [msg] log test\x1b[0;00m\n"},
		{"Errorln", Errorln, "\x1b[91m2017/06/25 15:49:04 [error] log test\x1b[0;00m\n"},
		{"Warningln", Warningln, "\x1b[93m2017/06/25 15:49:04 [warning] log test\x1b[0;00m\n"},
		{"Debugln", Debugln, ""},
	}

	for _, v := range data {
		err := validate(v.key, v.logFunc, v.expectedValue)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	DebugMode = true

	rescueStdout := os.Stdout
	defer func() { os.Stdout = rescueStdout }()

	r, w, err := os.Pipe()
	if err != nil {
		return
	}
	os.Stdout = w

	Debugln("log test")

	os.Stdout = rescueStdout

	err = w.Close()
	if err != nil {
		t.Fatal(err.Error())
	}

	out, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	rstr := "\x1b\\[96m2017/06/25 15:49:04 \\[debug\\] logsys_test.go:\\d+ log test\x1b\\[0;00m\n"
	match, err := regexp.Match(rstr, out)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !match {
		t.Fatalf("Error, 'Debugln' printed %q, not match with expected", string(out))
	}

}
