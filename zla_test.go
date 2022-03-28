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
	"time"

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

func WriteFormattedLogsInAllLevels(loggerInstance stdLogger.Logger) {

	loggerInstance.Showf("%s", messages[stdLogger.LogLevel_Show])
	loggerInstance.Fatalf(nil, "%s", messages[stdLogger.LogLevel_Fatal])
	loggerInstance.Errorf(nil, "%s", messages[stdLogger.LogLevel_Error])
	loggerInstance.Warnf("%s", messages[stdLogger.LogLevel_Warning])
	loggerInstance.Infof("%s", messages[stdLogger.LogLevel_Info])
	loggerInstance.Busf("%s", messages[stdLogger.LogLevel_Business])
	loggerInstance.Msgf("%s", messages[stdLogger.LogLevel_Message])
	loggerInstance.Dbgf("%s", messages[stdLogger.LogLevel_Debug])
	loggerInstance.Qryf("%s", messages[stdLogger.LogLevel_Query])
	loggerInstance.Tracef("%s", messages[stdLogger.LogLevel_Trace])
}

func WriteErrorDetailsInAllLevels(loggerInstance stdLogger.Logger) {

	err := e2h.Trace(fmt.Errorf("This is an error"))

	loggerInstance.Fatalf(err, "%s", messages[stdLogger.LogLevel_Fatal])
	loggerInstance.Errorf(err, "%s", messages[stdLogger.LogLevel_Error])
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
	loggerInstance.SetTimestampFunc(time.Now)

	Describe("The logger only log messages lower or equal to selected level", func() {

		JustBeforeEach(func() {
			result = nil //Limpia el array
		})

		JustAfterEach(func() {
			if CurrentGinkgoTestDescription().Failed {
				PrintResult(result)
			}
		})

		Context("When the Logger Level is invalid", func() {
			It("Return error", func() {
				err := loggerInstance.SetLogLevel("unknown")
				Expect(err).ShouldNot(BeNil())
			})
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

				WriteFormattedLogsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)
				CheckResult(result, 10)
			})
		})

		Context("When and e2h error is logged", func() {
			It("Dails are wrote on the log all level messages", func() {
				_ = loggerInstance.SetLogLevel("info")
				buf.Reset()

				WriteErrorDetailsInAllLevels(loggerInstance)

				ParseResult(buf.String(), &result)

				Expect(len(result)).To(Equal(2))
				for _, v := range result {
					Expect(v["details"]).ToNot(BeNil())
				}
			})
		})

		Context("When the RefId it's set", func() {
			It("Wrote the refId", func() {
				_ = loggerInstance.SetLogLevel("info")
				buf.Reset()

				loggerInstance.WithRefID("TheRefID").Info("The log message")

				ParseResult(buf.String(), &result)

				Expect(len(result)).To(Equal(1))
				for _, v := range result {
					Expect(v["ref"]).To(BeEquivalentTo("TheRefID"))
				}
			})
		})
	})
})
