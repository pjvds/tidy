# tidy

[![Build Status](https://travis-ci.org/pjvds/tidy.svg?branch=master)](https://travis-ci.org/pjvds/tidy)

Flexible structured logging for Go that is not fast by accident.

``` go
// simple fluent configuration API
// with loggers to console, or for example
// external services as logentries.com
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
  "user": session.User,
  "url": request.Url,
).Warn("unauthorized request")
```

## Output example

``` text
08:45:10 W (module): user authenticated     → username=catty
08:45:11 W (module): unauthorized request   → username=catty url=/list/5
```
