package tidy_test

import (
	"testing"

	"github.com/pjvds/tidy"
	"github.com/stretchr/testify/assert"
)

type EntryRecorderBackend struct {
	Entries []tidy.Entry
}

func (this *EntryRecorderBackend) IsEnabledFor(level tidy.Level, module tidy.Module) bool {
	// todo: make this configurable
	return true
}

func (this *EntryRecorderBackend) Build() tidy.Backend {
	return this
}

func (this *EntryRecorderBackend) FirstEntry() tidy.Entry {
	return this.Entries[0]
}

func (this *EntryRecorderBackend) Log(entry tidy.Entry) {
	this.Entries = append(this.Entries, entry)
}

func (this *EntryRecorderBackend) Flush() error {
	return nil
}

func TestConfigSetsModule(t *testing.T) {
	recoder := new(EntryRecorderBackend)
	logger := tidy.Configure().LogFromLevel(tidy.DEBUG).To(recoder).MustBuild()

	logger.Debug("foobar")

	assert.Equal(t, "tidy_test", recoder.FirstEntry().Module.String())
}
