package appengine

import (
	"github.com/pjvds/tidy"
	"google.golang.org/appengine/internal"
)

type backend struct {
	formatter tidy.PlainTextFormatter
}

var levelMap = []int64{
	4, // FATAL
	3, // ERROR
	2, // WARN
	1, // NOTICE
	1, // INFO
	0, // DEBUG
}

func (this *backend) IsEnabledFor(level tidy.Level, module tidy.Module) bool {
	// todo: make this configurable
	return true
}

func (this *backend) Log(entry tidy.Entry) {
	buffer := this.formatter.Format(entry)
	defer buffer.Free()

	internal.Logf(entry.Context, levelMap[entry.Level], buffer.String())
}

func (this *backend) Flush() error {
	return nil
}
