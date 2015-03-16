package tidy_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/pjvds/tidy"
	"github.com/stretchr/testify/assert"
)

func TestPlain(t *testing.T) {
	plain := tidy.PlainTextFormatter{}
	buffer := new(bytes.Buffer)

	entry := tidy.Entry{
		Timestamp: time.Now(),
		Module:    tidy.Module("tidy_test"),
		Level:     tidy.DEBUG,
		Message:   "log message",
		Fields: tidy.Fields{
			"a": "value a",
			"b": 42,
		},
	}

	plain.FormatTo(buffer, entry)
	assert.Equal(t, buffer.String(), "DEBUG (tidy_test): log message\t a=value a b=42\n")
}
