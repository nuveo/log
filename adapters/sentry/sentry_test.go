package sentry

import (
	"testing"

	"github.com/getsentry/raven-go"
	"github.com/nuveo/log"
)

func TestSentryLog(t *testing.T) {
	m := &MockTransport{}
	raven.DefaultClient.Transport = m

	log.Println("teste")
	if m.Count != 1 {
		t.Errorf("expected 1 but got %v", m.Count)
	}
}
