package logentries

import "github.com/pjvds/tidy"

type Config struct {
	address    string
	network    string
	token      string
	bufferSize int
}

// Configure a new logentries backend. Call `Config.Build` to build
// the backend instance.
func Configure(token string) Config {
	return Config{
		address:    "data.logentries.com:10000",
		network:    "udp",
		token:      token,
		bufferSize: 256,
	}
}

// Overrides the default log entries api endpoint address.
func (this Config) Address(address string) Config {
	this.address = address
	return this
}

// Sets the buffer size of the number of entries that can
// be held in memory. When the buffer is full, `Backend.Log` becomes blocking.
func (this Config) BufferSize(size int) Config {
	// TODO: allow to set block behavior
	this.bufferSize = size
	return this
}

// Changes the token after it has been set at construct.
func (this Config) Token(value string) Config {
	this.token = value
	return this
}

// Sets network protocol to UDP.
func (this Config) UDP() Config {
	this.network = "udp"
	return this
}

// Sets network protocol to TCP.
func (this Config) TCP() Config {
	this.network = "tcp"
	return this
}

// Build the backend based on the config.
func (this Config) Build() tidy.Backend {
	b := &backend{
		entries:   make(chan tidy.Entry, this.bufferSize),
		network:   this.network,
		address:   this.address,
		formatter: tidy.PlainTextFormatter{},
		token:     []byte(this.token),
	}
	go b.do()

	return b
}
