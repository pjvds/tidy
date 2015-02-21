package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBufferIsNotNil(t *testing.T) {
	buffer := NewBuffer()
	assert.NotNil(t, buffer)
}

func TestNewBufferIsAlwaysEmpty(t *testing.T) {
	for i := 0; i < 10000; i++ {
		buffer := NewBuffer()
		assert.Equal(t, 0, buffer.Len())

		buffer.WriteString("hello world")
		buffer.Free()
	}
}
