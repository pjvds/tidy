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
		Module:    tidy.NewModule("tidy_test"),
		Level:     tidy.DEBUG,
		Message:   "log message",
		Fields: tidy.Fields{
			"foo": "bar",
		},
	}

	plain.FormatTo(buffer, entry)
	assert.Equal(t, "DEBUG (tidy_test): log message\t foo=bar\n", buffer.String())
}
