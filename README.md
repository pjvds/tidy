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

## Git Commit Guidelines

We have very precise rules over how our git commit messages can be formatted.  This leads to **more
readable messages** that are easy to follow when looking through the **project history**.

### Commit Message Format
Each commit message consists of a **header**, a **body** and a **footer**.  The header has a special
format that includes a **type**, a **scope** and a **subject**:

```
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

Any line of the commit message cannot be longer 100 characters! This allows the message to be easier
to read on github as well as in various git tools.

### Type
Must be one of the following:

* **feat**: A new feature
* **fix**: A bug fix
* **docs**: Documentation only changes
* **style**: Changes that do not affect the meaning of the code (white-space, formatting, missing
  semi-colons, etc)
* **refactor**: A code change that neither fixes a bug or adds a feature
* **perf**: A code change that improves performance
* **test**: Adding missing tests
* **chore**: Changes to the build process or auxiliary tools and libraries such as documentation
  generation

### Scope
The scope could be anything specifying place of the commit change. For example `logger`,
`text`, `console`, `appengine`, `backend`, etc...

### Subject
The subject contains succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize first letter
* no dot (.) at the end

### Body
Just as in the **subject**, use the imperative, present tense: "change" not "changed" nor "changes"
The body should include the motivation for the change and contrast this with previous behavior.

### Footer
The footer might reference GitHub issues that this commit relates to or **Closes**.