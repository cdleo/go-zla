package zla_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cdleo/go-commons/logger"
	"github.com/cdleo/go-zla"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Tests Suite")
}

type logLine map[string]interface{}
type logResult []logLine

var messages = []string{"Disabled", "Show", "Fatal", "Error", "Warn", "Info", "Business", "Message", "Debug", "Query", "Trace"}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func ParseResult(buf string, result *logResult) {
	scanner := bufio.NewScanner(strings.NewReader(buf))
	for scanner.Scan() {
		var line logLine
		_ = json.Unmarshal(scanner.Bytes(), &line)
		*result = append(*result, line)
	}
}

func PrintResult(result logResult) {
	fmt.Printf("Result elements %d \n", len(result))
	for i, v := range result {
		fmt.Printf("Line %d => %s \n", i, v)
	}
}

func WriteLogsInAllLevels(loggerInstance logger.Logger) {

	loggerInstance.Show(messages[logger.LogLevel_Show])
	loggerInstance.Fatal(nil, messages[logger.LogLevel_Fatal])
	loggerInstance.Error(nil, messages[logger.LogLevel_Error])
	loggerInstance.Warn(messages[logger.LogLevel_Warning])
	loggerInstance.Info(messages[logger.LogLevel_Info])
	loggerInstance.Bus(messages[logger.LogLevel_Business])
	loggerInstance.Msg(messages[logger.LogLevel_Message])
	loggerInstance.Dbg(messages[logger.LogLevel_Debug])
	loggerInstance.Qry(messages[logger.LogLevel_Query])
	loggerInstance.Trace(messages[logger.LogLevel_Trace])
}

func CheckResult(result logResult, expected int) {
	Expect(len(result)).To(Equal(expected))
	for i, v := range result {
		Expect(v["message"]).To(BeEquivalentTo(messages[i+1]))
	}
}

var _ = Describe("Testing: LOGGER", func() {
	var (
		loggerInstance logger.Logger
		result         logResult
		buf            bytes.Buffer
	)

	loggerInstance, _ = zla.NewLogger()
	loggerInstance.SetOutput(&buf)

	Describe("The logger only log messages lower or equal to selected level", func() {

		JustBeforeEach(func() {
			result = nil //Limpia el array
		})

		JustAfterEach(func() {
			if CurrentGinkgoTestDescription().Failed {
				PrintResult(result)
			}
		})

		Context("When the Logger Level is disabled", func() {
			It("Wrote nothing", func() {
				_ = loggerInstance.SetLogLevel("disabled")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)
				ParseResult(buf.String(), &result)

				CheckResult(result, 0)
			})
		})

		Context("When the Logger Level is set to Show", func() {
			It("Only wrote Show level messages", func() {
				_ = loggerInstance.SetLogLevel("show")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 1)
			})
		})

		Context("When the Logger Level is set to Fatal", func() {
			It("Only wrote Fatal and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("fatal")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 2)
			})
		})

		Context("When the Logger Level is set to Error", func() {
			It("Only wrote Error and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("error")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 3)
			})
		})

		Context("When the Logger Level is set to Warning", func() {
			It("Only wrote Warning and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("warn")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 4)
			})
		})

		Context("When the Logger Level is set to Info", func() {
			It("Only wrote Info and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("info")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 5)
			})
		})

		Context("When the Logger Level is set to Business", func() {
			It("Only wrote Business and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("business")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 6)
			})
		})

		Context("When the Logger Level is set to Message", func() {
			It("Only wrote Message and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("message")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 7)
			})
		})

		Context("When the Logger Level is set to Debug", func() {
			It("Only wrote Debug and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("debug")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 8)
			})
		})

		Context("When the Logger Level is set to Query", func() {
			It("Only wrote Query and lower level messages", func() {
				_ = loggerInstance.SetLogLevel("query")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 9)
			})
		})

		Context("When the Logger Level is set to Trace", func() {
			It("Show all level messages", func() {
				_ = loggerInstance.SetLogLevel("trace")
				buf.Reset()

				WriteLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 10)
			})
		})
	})
})
