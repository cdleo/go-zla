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

	/*logConfig := loggerImple.NewLogConfig()
	logConfig.LogLevel = "debug"
	logConfig.Rotation.Path = "./log"
	logConfig.Rotation.NamePattern = "%Y%m%d.log"
	logConfig.Rotation.RotationTimeInHours = 24
	logConfig.Rotation.MaxAgeInDays = 30
	writer, err := fs.NewFileRotationWriter(config.Rotation)
	if err != nil {
		return errors.TraceA(err)
	}
	l.SetLogOutput(writer)*/

	// Output:
	// {"time":"2021-05-21T06:00:00-03:00","level":"INFO","message":"Log this!","where":"aca"}
	// {"time":"2021-05-21T06:00:00-03:00","ref":"ad7ec2d7-d92d-4d02-a937-e0c477611ffd","level":"ERROR","message":"This is an error log!","where":"aca","details":{"error":"foo","stack_trace":[{"func":"github.com/cdleo/go-zla_test.bar","caller":"zla_example_test.go:12"}]}}

}
