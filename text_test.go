package logging

import (
	"io/ioutil"
	"testing"
	"time"
)

var smallFields = Fields{
	"foo": "bar",
	"baz": 42,
}

type LengthRecorder struct {
	Length int
}

func (this *LengthRecorder) Write(data []byte) (int, error) {
	length := len(data)
	this.Length = length

	return length, nil
}

type DevNullWriter struct {
}

func (*DevNullWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

func BenchmarkColoredTextFormatterParallel(b *testing.B) {
	entry := Entry{
		Module:    Module("benchmark"),
		Timestamp: time.Now(),
		Level:     NOTICE,
		Message:   "bazinga",
		Fields: Fields{
			"foo": "bar",
			"baz": 42,
		},
	}

	formatter := &ColoredTextFormatter{}
	lengthRecorder := &LengthRecorder{}

	if err := formatter.FormatTo(lengthRecorder, entry); err != nil {
		b.Fatalf("failed to write: %v", err.Error())
	}
	length := int64(lengthRecorder.Length)

	b.ResetTimer()
	b.SetBytes(length)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			formatter.FormatTo(ioutil.Discard, entry)
		}
	})
}

func BenchmarkColoredTextFormatter(b *testing.B) {
	entry := Entry{
		Module:    Module("benchmark"),
		Timestamp: time.Now(),
		Level:     NOTICE,
		Message:   "bazinga",
		Fields: Fields{
			"foo": "bar",
			"baz": 42,
		},
	}

	formatter := &ColoredTextFormatter{}
	lengthRecorder := &LengthRecorder{}
	devNull := &DevNullWriter{}

	if err := formatter.FormatTo(lengthRecorder, entry); err != nil {
		b.Fatalf("failed to write: %v", err.Error())
	}
	length := int64(lengthRecorder.Length)

	b.ResetTimer()
	b.SetBytes(length)

	for n := 0; n < b.N; n++ {
		formatter.FormatTo(devNull, entry)
	}
}
