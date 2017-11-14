package sentry

import (
	"testing"

	"github.com/getsentry/raven-go"
	"github.com/nuveo/log"
)

func TestSentryLog(t *testing.T) {
	testCases := []struct {
		name          string
		fn            func(msg ...interface{})
		expectedCount int
		debugMode     bool
	}{
		{"log sentry", log.Errorln, 1, false},
		{"do not log debug", log.Debugln, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := &MockTransport{}
			raven.DefaultClient.Transport = m

			tc.fn("teste")
			if m.Count != tc.expectedCount {
				t.Errorf("expected %v, but got %v", tc.expectedCount, m.Count)
			}
		})
	}
}
