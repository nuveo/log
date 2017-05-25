package log

import "testing"

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Println("test log message")
		Errorln("test log message")
		Debugln("test log message")
	}
	Wait()
}
