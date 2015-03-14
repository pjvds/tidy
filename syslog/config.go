package syslog

import (
	. "log/syslog"

	"github.com/pjvds/tidy"
)

type Config struct {
	// The network (udp, tcp), or empty for local syslog socket. If the network is set
	// you also need to set the address.
	network string
	// The remote address, or empty for local syslog socket. If the network is set
	// you also need to set the address.
	address string

	// The TAG field value for the syslog entries. If empty the process name is used.
	tag string

	// The buffer size of the channel.
	bufferSize int
}

// Configure a new syslog backend. Call `Config.Build` to build
// the backend instance.
func Configure() Config {
	return Config{}
}

// Overrides the default local syslog address. Usually only used when logging to
// a remote syslog server.
func (this Config) Address(network string, address string) Config {
	this.network = network
	this.address = address
	return this
}

// Overrides the default tag that is set to the process name (os.Args[0]).
func (this Config) Tag(tag string) Config {
	this.tag = tag
	return this
}

// Sets the buffer size of the number of entries that can
// be held in memory. When the buffer is full, `Backend.Log` becomes blocking.
func (this Config) BufferSize(size int) Config {
	// TODO: allow to set block behavior
	this.bufferSize = size
	return this
}

// Build the backend based on the config.
func (this Config) Build() tidy.Backend {
	// Only one call to Dial is necessary.
	// On write failures, the syslog client will attempt to reconnect to the server and write again.
	writer, err := Dial(this.network, this.address, LOG_DEBUG, this.tag)

	// Dial only errors when there are invalid parameters, not when it can connect.
	if err != nil {
		panic(err)
	}

	b := &backend{
		entries:   make(chan tidy.Entry, this.bufferSize),
		formatter: tidy.PlainTextFormatter{},
		writer:    writer,
	}
	go b.do()

	return b
}
