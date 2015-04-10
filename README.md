# tidy

[![Build Status](https://travis-ci.org/pjvds/tidy.svg?branch=master)](https://travis-ci.org/pjvds/tidy)

Flexible structured logging for Go that is not fast by accident.

``` go
// simple fluent configuration API
tidy.Configure()
  .LogFromLevel(tidy.DEBUG).To(tidy.Console)
  .LogFromLevel(tidy.INFO).To(logentries.Token(token))
  .BuildDefault()

// get a logger for the current context
log := tidy.GetLogger()

// log entry with single field
log.With("user", session.User).Debug("user authenticated")

// log entry with multiple fields
log.Withs(tidy.Fields{
  "username": params.ByName("username"),
  "hash": params.ByName("hash"),
  "url": request.Url,
).Warn("unauthorized request")
```
