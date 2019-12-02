package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	LOGLEVEL = INFO
	Debug("debug")
	Info("info")
	Error("error")
}

func BenchmarkInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Info("123")
		}
	})
}
