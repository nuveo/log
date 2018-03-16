package withpath

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
	"time"

	"github.com/nuveo/log"
)

func getOutput(logFunc func(msg ...interface{}), msg ...interface{}) ([]byte, error) {
	rescueStdout := os.Stdout
	defer func() { os.Stdout = rescueStdout }()

	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	os.Stdout = w

	logFunc(msg...)

	err = w.Close()
	if err != nil {
		return nil, err
	}

	out, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func validate(key string, logFunc func(msg ...interface{}), valueExpected string, msg ...interface{}) (err error) {
	out, err := getOutput(logFunc, msg...)
	if err != nil {
		return
	}
	match, err := regexp.Match(valueExpected, out)
	if err != nil {
		return
	} else if !match {
		err = fmt.Errorf("Error, '%s' printed %q, expected %q", key, string(out), valueExpected)
	}
	return
}

func Test_withPathLog(t *testing.T) {
	now = func() time.Time { return time.Unix(1498405744, 0) }
	log.DebugMode = false

	data := []struct {
		key           string
		logFunc       func(msg ...interface{})
		expectedValue string
	}{
		{"Println", log.Println, "\x1b\\[37m2017/06/25 15:49:04 \\[msg\\] log test\x1b\\[0;00m\n"},
		{"Errorln", log.Errorln, "\x1b\\[91m2017/06/25 15:49:04 \\[error\\] log test\x1b\\[0;00m\n"},
		{"Warningln", log.Warningln, "\x1b\\[93m2017/06/25 15:49:04 \\[warning\\] log test\x1b\\[0;00m\n"},
		{"Debugln", log.Debugln, "\x1b\\[93m2017/06/25 15:49:04 \\[warning\\] log test\x1b\\[0;00m\n"},
	}
	formattedData := []struct {
		key           string
		logFunc       func(msg ...interface{})
		expectedValue string
	}{
		{"Printf", log.Printf, "\x1b\\[37m2017/06/25 15:49:04 \\[msg\\] formatted log 1.12\x1b\\[0;00m"},
		{"Errorf", log.Errorf, "\x1b\\[91m2017/06/25 15:49:04 \\[error\\] formatted log 1.12\x1b\\[0;00m"},
		{"Warningf", log.Warningf, "\x1b\\[93m2017/06/25 15:49:04 \\[warning\\] formatted log 1.12\x1b\\[0;00m"},
		{"Debugf", log.Debugf, ""},
	}
	for _, v := range data {
		err := validate(v.key, v.logFunc, v.expectedValue, "log test")
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	for _, v := range formattedData {
		err := validate(v.key, v.logFunc, v.expectedValue, "%s %s %.2f", "formatted", "log", 1.1234)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	log.DebugMode = true

	err := validate("Debugln", log.Debugln, "\x1b\\[96m2017/06/25 15:49:04 \\[debug\\] log_test.go:\\d+ log test\x1b\\[0;00m\n", "log test")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = validate("Debugf", log.Debugf, "\x1b\\[96m2017/06/25 15:49:04 \\[debug\\] log_test.go:\\d+ formatted log 1.12\x1b\\[0;00m", "%s %s %.2f", "formatted", "log", 1.1234)
	if err != nil {
		t.Fatal(err.Error())
	}
}
