package tidy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestLoggerContext(t *testing.T) {
	logger := NewLogger(GetModuleFromCaller(0), defaulBackend)

	ctx := context.Background()
	logger = logger.Context(ctx)

	assert.NotNil(t, logger.context)
}

func TestLoggerStacktrace(t *testing.T) {
	logger := NewLogger(GetModuleFromCaller(0), defaulBackend).WithStacktrace()

	stacktrace, ok := logger.fields["stacktrace"]
	assert.True(t, ok, "logger is missing stacktrace field")
	assert.Equal(t, "foobar", stacktrace)
}
