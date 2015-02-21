package main

import "github.com/pjvds/logging"

func main() {
	log := logging.NewLogger("main")
	log.Fatal("fatal")
	log.Error("error")
	log.WithField("week", 8).Warn("warning entry")
	log.Info("info")
	log.Debug("debug")

	log.WithFields(logging.Fields{
		"foo": "bar",
		"baz": 42,
	}).Info("hello world")

	log.Fatal("\\o/")
}
