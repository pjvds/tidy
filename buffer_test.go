package tidy_test

import (
	"testing"

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

	assert.Equal(t, writeAndReturnString(1), "01")
	buffer.Reset()

	assert.Equal(t, writeAndReturnString(9), "09")
	buffer.Reset()

	assert.Equal(t, writeAndReturnString(10), "10")
	buffer.Reset()

	assert.Equal(t, writeAndReturnString(12), "12")
	buffer.Reset()

	assert.Equal(t, writeAndReturnString(55), "55")
	buffer.Reset()

}
