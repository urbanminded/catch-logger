# catch-logger

`catch-logger` defines a simple logging interface to allow logging backends be abstracted
behind a common logging interface. This is especially important for library development -
libraries should not force client applications to adopt any particular logging technology.

This is easy in Go, since interfaces are defined at the point of use, rather than the
point of declaration; all that is required to be compatible with this library is to
create a small wrapper struct that conforms to the `logger.Logger` interface, delegating
to whichever backend is in use.

## Features

  - supports trace, debug, info, warn, and error levels
  - all log functions have a secondary form that accepts an additional dictionary of structured log arguments
  - incldues built in `System` wrapper for delegating to Go's `log` package, plus a silent `Discard` logger
  - supports attaching loggers to `context.Context` instances to avoid passing loggers throughout the application

## Installation

`catch-logger` uses Go modules.

```
go get github.com/urbanminded/catch-logger
```

## Usage

### Logger Interface

`catch-logger` defines the following logger interface; applications provide their
own implementations that delegate to whatever concrete log backend they happen
to be using (e.g. zerolog).

```go
type Logger interface {
	Trace(msg string, fields ...any)
	TraceArgs(args map[string]any, msg string, fields ...any)

	Debug(msg string, fields ...any)
	DebugArgs(args map[string]any, msg string, fields ...any)

	Info(msg string, fields ...any)
	InfoArgs(args map[string]any, msg string, fields ...any)

	Warn(msg string, fields ...any)
	WarnArgs(args map[string]any, msg string, fields ...any)

	Error(msg string, fields ...any)
	ErrorArgs(args map[string]any, msg string, fields ...any)

    // Additional error logger that accepts an error argument
	Err(err error, msg string, fields ...any)
	ErrArgs(err error, args map[string]any, msg string, fields ...any)
}
```

### Contextual Logging

`catch-logger` includes package-level functions that mirror those of the `Logger`
interface, but with an additional `context.Context` parameter, to enable logging
to occur anywhere a context is available without having to pass `Logger` instances
throughout the application.

Example:

```go
package main

import logger "github.com/urbanminded/catch-logger"

func main() {
    // WithLogger() derives a new context, with attached logger, from the given context
    ctx := logger.WithLogger(context.Background(), logger.System)

    // It's now possible to log using the the package level functions:
    logger.Info(ctx, "Hello world!")
}
```
