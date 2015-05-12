package appengine

import "github.com/pjvds/tidy"

type Config struct {
}

func Configure(token string) Config {
	return Config{}
}

// Build the backend based on the config.
func (this Config) Build() tidy.Backend {
	return &backend{}
}
