//+build !windows,!plan9
package syslog

import (
	"bytes"
	. "log/syslog"

	"github.com/pjvds/tidy"
)

type backend struct {
	entries   chan tidy.Entry
	writer    *Writer
	formatter tidy.PlainTextFormatter
}

func (this *backend) do() {
	buffer := new(bytes.Buffer)

	for entry := range this.entries {
		buffer.Reset()

		if err := this.formatter.FormatTo(buffer, entry); err != nil {
			// todo: add internal debug logging
			// not much we can do about it
			continue
		}

		switch entry.Level {
		case tidy.FATAL:
			this.writer.Crit(buffer.String())
		case tidy.ERROR:
			this.writer.Err(buffer.String())
		case tidy.WARN:
			this.writer.Warning(buffer.String())
		case tidy.NOTICE:
			this.writer.Notice(buffer.String())
		case tidy.INFO:
			this.writer.Info(buffer.String())
		case tidy.DEBUG:
			this.writer.Debug(buffer.String())
		default:
			// todo: add internal debug logging
			// not much we can do about it
		}
	}
}

func (this *backend) IsEnabledFor(level tidy.Level, module tidy.Module) bool {
	// todo: make this configurable
	return true
}

func (this *backend) Log(entry tidy.Entry) {
	this.entries <- entry
}

func (this *backend) Flush() error {
	return nil
}
