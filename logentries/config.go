package logentries

import "github.com/pjvds/tidy"

type Config struct {
	address string
	network string
	token   string
	failure func(error)
}

func Configure(token string) Config {
	return Config{
		address: "data.logentries.com:10000",
		network: "udp",
		token:   token,
	}
}

func (this Config) Token(value string) Config {
	this.token = value
	return this
}

func (this Config) UDP() Config {
	this.network = "udp"
	return this
}

func (this Config) TCP() Config {
	this.network = "tcp"
	return this
}

func (this Config) Build() tidy.Backend {
	b := &backend{
		entries:   make(chan tidy.Entry, 10000),
		network:   this.network,
		address:   this.address,
		formatter: tidy.PlainTextFormatter{},
		token:     []byte(this.token),
	}
	go b.do()

	return b
}
