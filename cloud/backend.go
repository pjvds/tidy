package cloud

import (
	"github.com/pjvds/tidy"
	"google.golang.org/cloud/logging"
)

var levels = map[tidy.Level]logging.Level{
	tidy.DEBUG:  logging.Debug,
	tidy.INFO:   logging.Info,
	tidy.NOTICE: logging.Alert,
	tidy.WARN:   logging.Warning,
	tidy.ERROR:  logging.Error,
	tidy.FATAL:  logging.Emergency,
}

type backend struct {
	client *logging.Client
}

func (this *backend) Log(entry tidy.Entry) {
	buffer := this.formatter.Format(entry)
	defer buffer.Free()

	payload := entry.Fields.Clone(2)
	payload["message"] = entry.Message
	payload["module"] = entry.Module

	this.client.Log(logging.Entry{
		Level:   levels[entry.Level],
		Payload: payload,
	})
}

func (this *backend) Flush() error {
	return this.client.Flush()
}
