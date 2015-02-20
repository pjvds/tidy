package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllowsDebug(t *testing.T) {
	assert.True(t, DEBUG.Allows(FATAL))
	assert.True(t, DEBUG.Allows(ERROR))
	assert.True(t, DEBUG.Allows(WARN))
	assert.True(t, DEBUG.Allows(INFO))
	assert.True(t, DEBUG.Allows(DEBUG))
}

func TestAllowsInfo(t *testing.T) {
	assert.True(t, INFO.Allows(FATAL))
	assert.True(t, INFO.Allows(ERROR))
	assert.True(t, INFO.Allows(WARN))
	assert.True(t, INFO.Allows(INFO))
	assert.False(t, INFO.Allows(DEBUG))
}

func TestAllowsWarn(t *testing.T) {
	assert.True(t, WARN.Allows(FATAL))
	assert.True(t, WARN.Allows(ERROR))
	assert.True(t, WARN.Allows(WARN))
	assert.False(t, WARN.Allows(INFO))
	assert.False(t, WARN.Allows(DEBUG))
}

func TestAllowsError(t *testing.T) {
	assert.True(t, ERROR.Allows(FATAL))
	assert.True(t, ERROR.Allows(ERROR))
	assert.False(t, ERROR.Allows(WARN))
	assert.False(t, ERROR.Allows(INFO))
	assert.False(t, ERROR.Allows(DEBUG))
}

func TestAllowsFatal(t *testing.T) {
	assert.True(t, FATAL.Allows(FATAL))
	assert.False(t, FATAL.Allows(ERROR))
	assert.False(t, FATAL.Allows(WARN))
	assert.False(t, FATAL.Allows(INFO))
	assert.False(t, FATAL.Allows(DEBUG))
}
