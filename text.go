package logging

import (
	"io"

	"github.com/wsxiaoys/terminal"
)

var colors = []string{
	"m", // Fatal, magenta
	"r", // Error, red
	"y", // Warn, yellow
	"g", // Notice, green
	"w", // Info, white
	"c", // Debug, cyan
}

type ColoredTextFormatter struct{}

func (this ColoredTextFormatter) FormatTo(writer io.Writer, entry Entry) error {
	buffer := NewBuffer()
	defer buffer.Free()

	term := terminal.TerminalWriter{buffer}
	color := colors[entry.Level]

	term.Color(color).Print(entry.Timestamp.Format("15:04:05"))
	term.Print(entry.Level.FixedString())
	term.Print(" [").Print(entry.Module).Print("] ").Reset()
	term.Print(" ").Print(entry.Message)

	if len(entry.Fields) > 0 {
		term.Print("\t")
		for key, value := range entry.Fields {
			term.Color(color).Print(" ").Print(key).Print("=").Reset()
			term.Print(value)
		}
	}

	term.Print("\n")

	_, err := buffer.WriteTo(writer)
	return err
}
