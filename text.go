package tidy

import (
	"fmt"
	"io"
	"sort"
)

var colors = [][]byte{
	[]byte("\033[35m"), // Fatal, magenta
	[]byte("\033[31m"), // Error, red
	[]byte("\033[33m"), // Warn, yellow
	[]byte("\033[32m"), // Notice, green
	[]byte(""),         // Info, default
	[]byte("\033[36m"), // Debug, cyan
}

var reset = []byte("\033[0m")
var newline = []byte("\n")
var whitespace = []byte(" ")
var colon = []byte(":")

type ColoredTextFormatter struct{}

func (this ColoredTextFormatter) Format(entry Entry) *FreeableBuffer {
	buffer := NewBuffer()

	color := colors[entry.Level]
	buffer.Write(color)

	buffer.WriteTwoDigits(entry.Timestamp.Hour())
	buffer.Write(colon)
	buffer.WriteTwoDigits(entry.Timestamp.Minute())
	buffer.Write(colon)
	buffer.WriteTwoDigits(entry.Timestamp.Second())

	buffer.WriteString(chars[entry.Level])
	buffer.WriteString(" ⟨")
	buffer.WriteString(entry.Module.String())
	buffer.WriteString("⟩")
	buffer.WriteString(": ")
	buffer.Write(reset)
	buffer.WriteString(entry.Message)

	if entry.Fields.Any() {
		buffer.Write(color)
		buffer.WriteString("\t→")
		keys := make([]string, 0, len(entry.Fields))

		for key, _ := range entry.Fields {
			keys = append(keys, key)
		}

		sort.Strings(keys)

		for _, key := range keys {
			value := entry.Fields[key]

			buffer.Write(whitespace)
			buffer.Write(color)
			buffer.WriteString(key)
			buffer.Write(reset)
			buffer.WriteString("=")

			valueAsString := fmt.Sprint(value)

			if len(valueAsString) > 1000 {
				valueAsString = valueAsString[0:1000] + "<truncated>"
			}

			buffer.WriteString(valueAsString)
		}
	}

	buffer.Write(newline)
	return buffer
}

func (this ColoredTextFormatter) FormatTo(writer io.Writer, entry Entry) error {
	buffer := this.Format(entry)
	defer buffer.Free()

	_, err := buffer.WriteTo(writer)
	return err
}

var (
	Console consoleBackendBuilder
)

type consoleBackendBuilder struct {
}

func (this consoleBackendBuilder) Build() Backend {
	return &ColoredConsoleBackend{}
}
