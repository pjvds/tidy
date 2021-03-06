package appengine

import (
	"testing"

	"github.com/pjvds/tidy"
	"google.golang.org/appengine/internal"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestBackendInvokesInternalLogger(t *testing.T) {
	logger := tidy.Configure().LogFromLevel(tidy.DEBUG).To(Configure()).MustBuild()

	var invoked bool
	var passedLevel int64
	var passedFormat string
	var passedArgs []interface{}

	ctx := internal.WithLogOverride(context.Background(), func(level int64, format string, args ...interface{}) {
		invoked = true
		passedLevel = level
		passedFormat = format
		passedArgs = args
		return
	})

	logger.Context(ctx).Debug("foobar")

	if !invoked {
		t.Fatalf("internal logger never called")
	}

	assert.Equal(t, int64(0), passedLevel)
	assert.Equal(t, "DEBUG (appengine): foobar\t location=backend_test.go:29", passedFormat)
}

func TestBackendDoesNotPanicOnNilContext(t *testing.T) {
	logger := tidy.Configure().LogFromLevel(tidy.DEBUG).To(Configure()).MustBuild()
	logger.Context(nil).Debug("foobar")
}

func TestBackendDoesNotPanicOnInvalidContext(t *testing.T) {
	logger := tidy.Configure().LogFromLevel(tidy.DEBUG).To(Configure()).MustBuild()
	logger.Context(context.Background()).Debug("foobar")
}
