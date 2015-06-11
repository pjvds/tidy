package appengine

import (
	"github.com/pjvds/tidy"
	"google.golang.org/appengine/log"
)

type backend struct {
	formatter tidy.PlainTextFormatter
}

func (this *backend) Log(entry tidy.Entry) {
	// ignore all requests that have no context, because
	// this causes the appengine logger to fail with a panic.
	if entry.Context == nil {
		return
	}

	// // The internal.Logf method of the appengine package will
	// // fail if the following context key is not set, since it
	// // expects a *context type there.
	// //
	// // See: http://git.io/vIG06
	// if netctx := entry.Context.Value(&"holds a *context"); netctx == nil {
	// 	return
	// }
	//
	//   ^
	//   |
	//   |
	//  DOES NOT WORK BECAUSE "holds a *context" IS REFERENCED IN THE INTERNAL
	//  PACKAGE!!!

	buffer := this.formatter.Format(entry)
	defer buffer.Free()

	func(message string) {
		// when the `appengine/log` package method panic, recover from it.
		defer func() {
			if recovered := recover(); recovered != nil {
				// TODO: log to internal backend logger,
				// for now just eat up.
			}
		}()

		switch entry.Level {
		case tidy.DEBUG:
			log.Debugf(entry.Context, message)
		case tidy.INFO:
			log.Infof(entry.Context, message)
		case tidy.NOTICE:
			log.Infof(entry.Context, message)
		case tidy.WARN:
			log.Warningf(entry.Context, message)
		case tidy.ERROR:
			log.Errorf(entry.Context, message)
		case tidy.FATAL:
			log.Criticalf(entry.Context, message)
		}
	}(buffer.String())
}

func (this *backend) Flush() error {
	return nil
}
