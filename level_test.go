package tidy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// An level allows itself and above.
func ExampleLevel_Allows() {
	fmt.Println(NOTICE.Allows(DEBUG))
	fmt.Println(NOTICE.Allows(INFO))
	fmt.Println(NOTICE.Allows(NOTICE))
	fmt.Println(NOTICE.Allows(WARN))
	fmt.Println(NOTICE.Allows(ERROR))
	fmt.Println(NOTICE.Allows(FATAL))

	// output:
	// false
	// false
	// true
	// true
	// true
	// true
}

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

func TestFatal(t *testing.T) {
	assert.Equal(t, "FATAL", FATAL.String())
}

func TestError(t *testing.T) {
	assert.Equal(t, "ERROR", ERROR.String())
}

func TestWarn(t *testing.T) {
	assert.Equal(t, "WARN", WARN.String())
}

func TestNotice(t *testing.T) {
	assert.Equal(t, "NOTICE", NOTICE.String())
}

func TestInfo(t *testing.T) {
	assert.Equal(t, "INFO", INFO.String())
}

func TestDebug(t *testing.T) {
	assert.Equal(t, "DEBUG", DEBUG.String())
}
