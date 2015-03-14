package tidy

import (
	"fmt"
	"io"
)

type PlainTextFormatter struct{}

func (PlainTextFormatter) FormatTo(writer io.Writer, entry Entry) error {
	buffer := NewBuffer()
	defer buffer.Free()

	// TODO: make format configurable
	//buffer.WriteString(entry.Timestamp.Format("15:04:05.000 "))
	buffer.WriteString(entry.Level.String())
	buffer.WriteString(" (")
	buffer.WriteString(entry.Module.String())
	buffer.WriteString("): ")
	buffer.WriteString(entry.Message)

	if entry.Fields.Any() {
		buffer.WriteString("\t")
		for key, value := range entry.Fields {
			buffer.Write(whitespace)
			buffer.WriteString(key)
			buffer.WriteString("=")
			buffer.WriteString(fmt.Sprint(value))
		}
	}

	buffer.Write(newline)

	_, err := buffer.WriteTo(writer)
	return err
}
