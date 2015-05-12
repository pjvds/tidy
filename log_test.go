package tidy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestLoggerContext(t *testing.T) {
	logger := NewLogger(GetModuleFromCaller(0), Console.Build())

	ctx := context.Background()
	logger = logger.Context(ctx)

	assert.NotNil(t, logger.context)
}
