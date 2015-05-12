package appengine

import (
	"github.com/pjvds/tidy"
	"google.golang.org/appengine/log"
)

type backend struct {
	formatter tidy.PlainTextFormatter
}

func (this *backend) IsEnabledFor(level tidy.Level, module tidy.Module) bool {
	// todo: make this configurable
	return true
}

func (this *backend) Log(entry tidy.Entry) {
	// ignore all requests that have no context, because
	// this causes the appengine logger to fail with a panic.
	if entry.Context == nil {
		return
	}

	buffer := this.formatter.Format(entry)
	defer buffer.Free()

	switch entry.Level {
	case tidy.DEBUG:
		log.Debugf(entry.Context, buffer.String())
	case tidy.INFO:
		log.Infof(entry.Context, buffer.String())
	case tidy.NOTICE:
		log.Infof(entry.Context, buffer.String())
	case tidy.WARN:
		log.Warningf(entry.Context, buffer.String())
	case tidy.ERROR:
		log.Errorf(entry.Context, buffer.String())
	case tidy.FATAL:
		log.Criticalf(entry.Context, buffer.String())
	}
}

func (this *backend) Flush() error {
	return nil
}
