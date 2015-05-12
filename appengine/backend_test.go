package appengine

import (
	"testing"

	"github.com/pjvds/tidy"
	"google.golang.org/appengine/internal"

	"golang.org/x/net/context"
)

func TestBackendInvokesInternalLogger(t *testing.T) {
	logger := tidy.Configure().LogFromLevel(tidy.DEBUG).To(Configure()).MustBuild()

	invoked := false
	ctx := internal.WithLogOverride(context.Background(), func(level int64, format string, args ...interface{}) {
		invoked = true
		return
	})

	logger.Context(ctx).Debug("foobar")

	if !invoked {
		t.Fatalf("internal logger never called")
	}
}

func TestBackendDoesNotPanicOnNilContext(t *testing.T) {
	logger := tidy.Configure().LogFromLevel(tidy.DEBUG).To(Configure()).MustBuild()
	logger.Context(nil).Debug("foobar")
}
