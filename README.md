# GO-ZLA

[![Go Reference](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://pkg.go.dev/github.com/cdleo/go-zla) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/cdleo/go-zla/master/LICENSE) [![Build Status](https://scrutinizer-ci.com/g/cdleo/go-zla/badges/build.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-zla/build-status/master) [![Code Coverage](https://scrutinizer-ci.com/g/cdleo/go-zla/badges/coverage.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-zla/?branch=master) [![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/cdleo/go-zla/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/cdleo/go-zla/?branch=master)

GO ZeroLog Adapter (a.k.a. go-zla) is a lightweight Golang module that adapt the use of the zerolog logger to a standard interface provided by the go-facades project.

## General

Details of the standard interface could be found [HERE](https://github.com/cdleo/go-facades/blob/master/README.md)

**Usage**
This example program shows the initialization and the use of some levels and features:
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
	// {"time":"2021-05-21T06:00:00-03:00","level":"INFO","message":"Log this!","where":"aca"}
	// {"time":"2021-05-21T06:00:00-03:00","ref":"ad7ec2d7-d92d-4d02-a937-e0c477611ffd","level":"ERROR","message":"This is an error log!","where":"aca","details":{"error":"foo","stack_trace":[{"func":"github.com/cdleo/go-zla_test.bar","caller":"zla_example_test.go:12"}]}}

}
```

## Sample

You can find a sample of the use of go-zla project [HERE](https://github.com/cdleo/go-zla/blob/master/zla_example_test.go)

## Contributing

Comments, suggestions and/or recommendations are always welcomed. Please check the [Contributing Guide](CONTRIBUTING.md) to learn how to get started contributing.
