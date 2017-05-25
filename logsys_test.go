package log

import "testing"

func BenchmarkHello(b *testing.B) {
	DebugMode = true
	for i := 0; i < b.N; i++ {
		Println("test log message")
		Errorln("test log message")
		Debugln("test log message")
	}
}
