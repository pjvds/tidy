package tidy_test

import (
	"testing"
	"time"

	"github.com/pjvds/tidy"
	"github.com/stretchr/testify/assert"
)

func TestNewBufferIsNotNil(t *testing.T) {
	buffer := tidy.NewBuffer()
	assert.NotNil(t, buffer)
}

func TestNewBufferIsAlwaysEmpty(t *testing.T) {
	for i := 0; i < 10000; i++ {
		buffer := tidy.NewBuffer()
		assert.Equal(t, 0, buffer.Len())

		buffer.WriteString("hello world")
		buffer.Free()
	}
}

func TestBufferWriteTwoDigits(t *testing.T) {
	buffer := tidy.NewBuffer()
	defer buffer.Free()

	writeAndReturnString := func(n int) string {
		buffer.WriteTwoDigits(n)
		return string(buffer.Bytes())
	}

	assert.Equal(t, "01", writeAndReturnString(1))
	buffer.Reset()

	assert.Equal(t, "09", writeAndReturnString(9))
	buffer.Reset()

	assert.Equal(t, "10", writeAndReturnString(10))
	buffer.Reset()

	assert.Equal(t, "12", writeAndReturnString(12))
	buffer.Reset()

	assert.Equal(t, "55", writeAndReturnString(55))
	buffer.Reset()

	assert.Equal(t, "68", writeAndReturnString(168))
	buffer.Reset()
}

func BenchmarkClockWritingWithWriteTwoDigits(b *testing.B) {
	b.ReportAllocs()

	buffer := tidy.NewBuffer()
	hour, minute, second := time.Now().Clock()

	for n := 0; n < b.N; n++ {
		buffer.Grow(7)
		buffer.WriteTwoDigits(hour)
		buffer.WriteRune(':')
		buffer.WriteTwoDigits(minute)
		buffer.WriteRune(':')
		buffer.WriteTwoDigits(second)
	}
}

func BenchmarkClockWritingWithTimeFormat(b *testing.B) {
	b.ReportAllocs()

	buffer := tidy.NewBuffer()
	clock := time.Now()

	for n := 0; n < b.N; n++ {
		buffer.Grow(7)
		buffer.WriteString(clock.Format("15:04:05"))
	}
}
