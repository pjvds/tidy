package logentries

import (
	"bytes"
	"net"
	"time"

	"github.com/pjvds/tidy"
)

const (
	network = "udp"
	address = "data.logentries.com"
	port    = 20000
)

type Backend struct {
	entries   chan tidy.Entry
	network   string
	address   string
	formatter tidy.PlainTextFormatter
	token     []byte
}

func New(token string) tidy.Backend {
	backend := &Backend{
		entries:   make(chan tidy.Entry, 5000),
		network:   "tcp",
		address:   "data.logentries.com:10000",
		formatter: tidy.PlainTextFormatter{},
		token:     []byte(token),
	}

	go backend.do()
	return backend
}

func (this *Backend) do() {
	buffer := new(bytes.Buffer)
	buffer.Write(this.token)

	entryStart := len(this.token)
CONNECT:
	conn, err := net.Dial(this.network, this.address)
	if err != nil {
		time.Sleep(3 * time.Second)
		goto CONNECT
	}

	for entry := range this.entries {
		buffer.Truncate(entryStart)
		this.formatter.FormatTo(buffer, entry)

		for {
			if _, err := conn.Write(buffer.Bytes()); err != nil {
				if neterr, ok := err.(net.Error); ok && neterr.Temporary() {
					continue
				}
				goto CONNECT
			}
			break
		}
	}
}

func (this *Backend) IsEnabledFor(level tidy.Level, module tidy.Module) bool {
	return true
}

func (this *Backend) Log(entry tidy.Entry) {
	this.entries <- entry
}

func (this *Backend) Flush() error {
	return nil
}
