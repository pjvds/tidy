package tidy

import (
	"testing"
	"time"
)

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

// smallFields is a small size data set for benchmarking
var smallFields = Fields{
	"foo":   "bar",
	"baz":   "qux",
	"one":   "two",
	"three": "four",
}

// largeFields is a large size data set for benchmarking
var largeFields = Fields{
	"foo":       "bar",
	"baz":       "qux",
	"one":       "two",
	"three":     "four",
	"five":      "six",
	"seven":     "eight",
	"nine":      "ten",
	"eleven":    "twelve",
	"thirteen":  "fourteen",
	"fifteen":   "sixteen",
	"seventeen": "eighteen",
	"nineteen":  "twenty",
	"a":         "b",
	"c":         "d",
	"e":         "f",
	"g":         "h",
	"i":         "j",
	"k":         "l",
	"m":         "n",
	"o":         "p",
	"q":         "r",
	"s":         "t",
	"u":         "v",
	"w":         "x",
	"y":         "z",
	"this":      "will",
	"make":      "thirty",
	"entries":   "yeah",
}

// func BenchmarkColoredTextFormatterParallel(b *testing.B) {
// 	entry := Entry{
// 		Module:    Module("benchmark"),
// 		Timestamp: time.Now(),
// 		Level:     NOTICE,
// 		Message:   "bazinga",
// 		Fields: Fields{
// 			"foo": "bar",
// 			"baz": 42,
// 		},
// 	}

// 	formatter := &ColoredTextFormatter{}
// 	lengthRecorder := &LengthRecorder{}

// 	if err := formatter.FormatTo(lengthRecorder, entry); err != nil {
// 		b.Fatalf("failed to write: %v", err.Error())
// 	}
// 	length := int64(lengthRecorder.Length)

// 	b.ResetTimer()
// 	b.SetBytes(length)

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			formatter.FormatTo(ioutil.Discard, entry)
// 		}
// 	})
// }

func BenchmarkSmallTextColoredTextFormatter(b *testing.B) {
	entry := Entry{
		Module:    NewModule("benchmark"),
		Timestamp: time.Now(),
		Level:     INFO,
		Message:   "message",
		Fields:    smallFields,
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

func BenchmarkLargeTextColoredTextFormatter(b *testing.B) {
	entry := Entry{
		Module:    NewModule("benchmark"),
		Timestamp: time.Now(),
		Level:     INFO,
		Message:   "message",
		Fields:    largeFields,
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
