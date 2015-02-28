package logentries

import (
	"bytes"
	"net"
	"time"

	"github.com/pjvds/backoff"
	"github.com/pjvds/tidy"
)

type backend struct {
	entries    chan tidy.Entry
	network    string
	address    string
	formatter  tidy.PlainTextFormatter
	token      []byte
	connection net.Conn
}

func (this *backend) do() {
	buffer := new(bytes.Buffer)
	buffer.Write(this.token)

	// reset to this point to for every entry.
	entryStart := len(this.token)

	delay := backoff.Exp(time.Millisecond, 10*time.Second)
DIAL:
	conn, err := net.Dial(this.network, this.address)

	if err != nil {
		conn.Close()
		delay.Delay()
		goto DIAL
	}

	for entry := range this.entries {
		// reset to the point after the token
		buffer.Truncate(entryStart)

		// format the message into the buffer
		this.formatter.FormatTo(buffer, entry)

		if _, err := buffer.WriteTo(conn); err != nil {
			conn.Close()
			delay.Delay()
			goto DIAL
		}
	}
}

func (this *backend) IsEnabledFor(level tidy.Level, module tidy.Module) bool {
	return true
}

func (this *backend) Log(entry tidy.Entry) {
	this.entries <- entry
}

func (this *backend) Flush() error {
	return nil
}
