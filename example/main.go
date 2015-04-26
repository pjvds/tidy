package main

import "github.com/pjvds/tidy"

func main() {
	// configure the default logger to log to the console.
	tidy.Configure().LogFromLevel(tidy.DEBUG).To(tidy.Console).BuildDefault()

	// get a logger for the current context
	log := tidy.GetLogger()
	// V returns a noop type if current level doesn't match.
	log.V(tidy.INFO).With("foo", "bar").Write("info message")

	// log at error level
	log.Error("error")
	
	// include week field and log warning a entry.
	log.With("week", 8).Warn("warning entry")
	log.Info("info")
	log.Debug("debug")

	// include many fields at once.
	log.Withs(tidy.Fields{
		"foo": "bar",
		"baz": 42,
	}).Info("hello world")

	log.Fatal("\\o/")
}
