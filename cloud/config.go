package appengine

import (
	"github.com/pjvds/tidy"
	"golang.org/x/net/context"
	"google.golang.org/cloud"
	"google.golang.org/cloud/logging"
)

type Config struct {
	ctx       context.Context
	projectID string
	logName   string
	opts      []cloud.ClientOption
}

func Configure(ctx, projectID, logName string, opts ...cloud.ClientOption) Config {
	return Config{
		ctx:       ctx,
		projectID: projectID,
		logName:   logName,
		opts:      opts,
	}
}

// Build the backend based on the config.
func (this Config) Build() tidy.Backend {
	client, err := logging.NewClient(this.ctx, this.projectID, this.logName, this.opts...)

	// TODO: we should be able to return this error to indicate a build failure.
	if err != nil {
		panic(err)
	}

	return &backend{
		client: client,
	}
}
