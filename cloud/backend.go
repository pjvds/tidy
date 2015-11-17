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
	sync   bool
}

func (this *backend) Log(entry tidy.Entry) {
	payload := entry.Fields.Clone(2)
	payload["message"] = entry.Message
	payload["module"] = entry.Module

	cloudEntry := logging.Entry{
		Level:   levels[entry.Level],
		Payload: payload,
	}

	var err error
	if this.sync {
		err = this.client.LogSync(cloudEntry)
	} else {
		err = this.client.Log(cloudEntry)
	}

	if err != nil {
		panic(err)
	}
}

func (this *backend) Flush() error {
	return this.client.Flush()
}
