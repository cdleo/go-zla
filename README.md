# GO-E2H

[![Go Reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/cdleo/go-e2h) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/cdleo/go-e2h/master/LICENSE) [![Build Status](https://scrutinizer-ci.com/g/cdleo/go-e2h/badges/build.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-e2h/build-status/master) [![Code Coverage](https://scrutinizer-ci.com/g/cdleo/go-e2h/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-e2h/?branch=master) [![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/cdleo/go-e2h/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-e2h/?branch=master)

GO ZeroLog Adapter (a.k.a. go-zla) is a lightweight Golang module to add a better stack trace and context information on error events.
As it's name says, this implementation uses zerolog (https://github.com/rs/zerolog) as writer.

## General

The logger contract resides on the go-commons repository: [github.com/cdleo/go-commons/logger/logger.go](https://github.com/cdleo/go-commons/logger/logger.go):
```go
type Logger interface {
	//Sets the current log level. (e.g. "debug")
	SetLogLevel(level string) error
	//Sets the log writer. (e.g. os.Stdout)
	SetOutput(w io.Writer)
	//Sets the function to write the log's timestamp. (e.g. time.Now)
	SetTimestampFunc(f func() time.Time)

	//Includes the ref field on the related log msg call
	WithRefID(refID string) Logger

	Show(msg string)
	Showf(msg string, v ...interface{})

	Fatal(err error, msg string)
	Fatalf(err error, msg string, v ...interface{})

	Error(err error, msg string)
	Errorf(err error, msg string, v ...interface{})

	Warn(msg string)
	Warnf(msg string, v ...interface{})

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

**Usage**
This example program shows the initialization and the use of different levels:
```go
package zla_test

import (
	"fmt"
	"time"

	"github.com/cdleo/go-e2h"
	"github.com/cdleo/go-zla"
)

func bar() error {
	return e2h.Trace(fmt.Errorf("foo"))
}

func Example_logger() {

	logger, _ := zla.NewLogger()

	//We've set this time func in order to always get the same time in the logger output
	mockedDateTime := time.Date(2021, 05, 21, 9, 00, 00, 000000000, time.UTC)
	logger.SetTimestampFunc(mockedDateTime.Local)

	//By default, the logger in fully initialized with level Info and writes to StdOutput
	logger.Info("Log this!")

	reqId := "ad7ec2d7-d92d-4d02-a937-e0c477611ffd"
	logger.WithRefID(reqId).Error(bar(), "This is an error log!")

	// Output:
	// {"time":"2021-05-21T06:00:00-03:00","level":"INFO","message":"Log this!","where":"zla_example_test.go:24"}
	// {"time":"2021-05-21T06:00:00-03:00","ref":"ad7ec2d7-d92d-4d02-a937-e0c477611ffd","level":"ERROR","message":"This is an error log!","where":"zla_example_test.go:27","details":{"error":"foo","stack_trace":[{"func":"github.com/cdleo/go-zla_test.bar","caller":"zla_example_test.go:12"}]}}
}
```

## Sample

You can find a sample of the use of go-e2h project [HERE](https://github.com/cdleo/go-zla/blob/master/zla_example_test.go)

## Contributing

Comments, suggestions and/or recommendations are always welcomed. Please check the [Contributing Guide](CONTRIBUTING.md) to learn how to get started contributing.
