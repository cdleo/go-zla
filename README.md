# GO-E2H

[![Go Reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/cdleo/go-e2h) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/cdleo/go-e2h/master/LICENSE) [![Build Status](https://scrutinizer-ci.com/g/cdleo/go-e2h/badges/build.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-e2h/build-status/master) [![Code Coverage](https://scrutinizer-ci.com/g/cdleo/go-e2h/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-e2h/?branch=master) [![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/cdleo/go-e2h/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-e2h/?branch=master)

GO ZeroLog Adapter (a.k.a. go-zla) is a lightweight Golang module to add a better stack trace and context information on error events.

## General

We created this abstraction keeping in mind the current VT-NET log, trying to use same levels and meanings.
That interface resides on gitlab.veritran.net/core/go-vt-platform/lib/core/interfaces/logger:
```go
type Logger interface {
	Initialize(config LogConfig) errors.StackedError
	SetLogLevel(logLevel string) errors.StackedError
	Ctx(ctx context.Context) Logger

	SetLogOutput(w io.Writer)
	SetTimeFunction(f func() time.Time)

	Show(msg string)
	Showf(msg string, v ...interface{})

	Fatal(err errors.StackedError, msg string)
	Fatalf(err errors.StackedError, msg string, v ...interface{})

	Error(err errors.StackedError, msg string)
	Errorf(err errors.StackedError, msg string, v ...interface{})

	Warn(err errors.StackedError, msg string)
	Warnf(err errors.StackedError, msg string, v ...interface{})

	Info(msg string)
	Infof(msg string, v ...interface{})

	Bus(msg string)
	Busf(msg string, v ...interface{})

	Msg(msg string)
	Msgf(msg string, v ...interface{})

	Dbg(msg string)
	Dbgf(msg string, v ...interface{})

	Qry(msg string)
	Qryf(msg string, v ...interface{})

	Trace(msg string)
	Tracef(msg string, v ...interface{})
}
```

**Log levels**
This is the list of log levels and their meanings:
- **disabled**: No log
- **show**: Very low frequency messages that should always be displayed (such as copyright)
- **fatal**: Errors that cause the application to stop
- **error**: Errors that don't stop the app
- **warning**: Alert conditions but not generating an error
- **info**: Important details, such as the version of the components when lifting
- **business**: Important business details, such as the identification of the transaction and the user who carried it out
- **message**: Details of all the Business request and response exchanged with the outside (does not include internal messages from the app, such as commands)
- **debug**: Detailed execution's information
- **query**: Detail of executed querys with their input parameters
- **trace**: Maximum level of detail, such as the values returned by the querys, http trace, etc.

Current abstraction's implementation, uses zerolog (https://github.com/rs/zerolog) as writer and go-file-rotatelogs (https://github.com/lestrrat/go-file-rotatelogs) for log rotation and older files cleanup.

**Usage**
This example program shows the initialization and the use of different levels:
```go
package logger_test

import (
	"context"
	"fmt"
	"time"

	loggerImple "gitlab.veritran.net/core/go-vt-platform/lib/core/implementations/logger"
	"gitlab.veritran.net/core/go-vt-platform/lib/core/implementations/clock"
	"gitlab.veritran.net/core/go-vt-platform/lib/core/interfaces/errors"
	"gitlab.veritran.net/core/go-vt-platform/lib/core/interfaces/logger"
)

func bar() errors.StackedError {
	return errors.WrapA(fmt.Errorf("foo"))
}

func Example_logger() {

	loggerInstance := loggerImple.NewLogger()

	//We've set this time func in order to always get the same time in the logger output
	stoppedTime := clock.NewStoppedClock(time.Date(2021, 05, 21, 9, 00, 00, 000000000, time.UTC))
	loggerInstance.SetTimeFunction(stoppedTime.CurrentInstant)

	//By default, the logger in fully initialized with level Info and writes to StdOutput
	loggerInstance.Info("Log this!")

	ctx := context.WithValue(context.Background(), logger.RequestIdKey, "Example")
	loggerInstance.Ctx(ctx).Error(bar(), "This is an error log!")

	logConfig := loggerImple.NewLogConfig()
	logConfig.LogLevel = "debug"
	logConfig.Rotation.Path = "./log"
	logConfig.Rotation.NamePattern = "%Y%m%d.log"
	logConfig.Rotation.RotationTimeInHours = 24
	logConfig.Rotation.MaxAgeInDays = 30

	//Anyway, you can re-initialize the logger adding file rotation config
	err := loggerInstance.Initialize(logConfig)
	if err != nil {
		loggerInstance.Error(err, "Unable to initialize logger")
	}

	// Output:
	// {"time":"2021-05-21T09:00:00Z","level":"INFO","message":"Log this!","where":"gitlab.veritran.net/core/go-vt-platform/lib/core_test/implementations/logger/example_test.go:27"}
	// {"time":"2021-05-21T09:00:00Z","ctx":"Example","level":"ERROR","message":"This is an error log!","where":"gitlab.veritran.net/core/go-vt-platform/lib/core_test/implementations/logger/example_test.go:30","details":{"error":"foo","stack_trace":[{"func":"gitlab.veritran.net/core/go-vt-platform/lib/core_test/implementations/logger_test.bar","caller":"gitlab.veritran.net/core/go-vt-platform/lib/core_test/implementations/logger/example_test.go:15"}]}}
}
```

## Sample

You can find a sample of the use of go-e2h project [HERE](https://github.com/cdleo/go-zla/blob/master/zla_example_test.go)

## Contributing

Comments, suggestions and/or recommendations are always welcomed. Please check the [Contributing Guide](CONTRIBUTING.md) to learn how to get started contributing.
