package logging

import (
	"fmt"
	"io"
)

var colors = [][]byte{
	[]byte("\033[35m"), // Fatal, magenta
	[]byte("\033[31m"), // Error, red
	[]byte("\033[33m"), // Warn, yellow
	[]byte("\033[32m"), // Notice, green
	[]byte("\033[37m"), // Info, white
	[]byte("\033[36m"), // Debug, cyan
}

var reset = []byte("\033[0m")
var newline = []byte("\n")
var whitespace = []byte(" ")

type ColoredTextFormatter struct{}

func (this ColoredTextFormatter) FormatTo(writer io.Writer, entry Entry) error {
	buffer := NewBuffer()
	defer buffer.Free()

	color := colors[entry.Level]

	buffer.Write(color)
	buffer.WriteString(entry.Timestamp.Format("15:04:05 "))
	buffer.WriteString(entry.Level.String())
	buffer.WriteString(" ⟨")
	buffer.WriteString(entry.Module.String())
	buffer.WriteString("⟩: ")
	buffer.Write(reset)
	buffer.WriteString(entry.Message)

	if entry.Fields.Any() {
		buffer.Write(color)
		buffer.WriteString("\t→")
		for key, value := range entry.Fields {
			buffer.Write(whitespace)
			buffer.Write(color)
			buffer.WriteString(key)
			buffer.Write(reset)
			buffer.WriteString("=")
			buffer.WriteString(fmt.Sprint(value))
		}
	}

	buffer.Write(newline)

	_, err := buffer.WriteTo(writer)
	return err
}
