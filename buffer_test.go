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
