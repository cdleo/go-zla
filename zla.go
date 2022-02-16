package zla

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/cdleo/go-e2h"
	e2hformat "github.com/cdleo/go-e2h/formatter"

	stdLogger "github.com/cdleo/go-facades/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type zlogAdapter struct {
	level           stdLogger.LogLevel
	refID           string
	formatter       e2hformat.Formatter
	formatterParams e2hformat.Params
}

func NewLogger() (stdLogger.Logger, error) {

	zerolog.SetGlobalLevel(zerolog.NoLevel)

	formatter, err := e2hformat.NewFormatter(e2hformat.Format_JSON)
	if err != nil {
		return nil, e2h.Trace(err)
	}

	_, b, _, _ := runtime.Caller(0)
	hideThisPath := filepath.Dir(b) + string(os.PathSeparator)
	params := e2hformat.Params{
		Beautify:         false,
		InvertCallstack:  false,
		PathHidingMethod: e2hformat.HidingMethod_FullBaseline,
		PathHidingValue:  hideThisPath,
	}

	l := &zlogAdapter{
		refID:           "",
		formatter:       formatter,
		formatterParams: params,
	}

	err = l.SetLogLevel(stdLogger.LogLevel_Info.String())
	if err != nil {
		return nil, e2h.Trace(err)
	}

	l.SetOutput(os.Stdout)
	l.SetTimestampFunc(time.Now)
	return l, nil
}

func (l *zlogAdapter) SetLogLevel(level string) error {

	var err error
	l.level, err = stdLogger.NewLogLevel(level)
	if err != nil {
		return e2h.Trace(err)
	}
	return nil
}

func (l *zlogAdapter) SetOutput(w io.Writer) {
	log.Logger = log.Logger.Output(w)
}

func (l *zlogAdapter) SetTimestampFunc(f func() time.Time) {
	zerolog.TimestampFunc = f
}

func (l *zlogAdapter) WithRefID(refID string) stdLogger.Logger {
	return &zlogAdapter{
		level:           l.level,
		formatter:       l.formatter,
		formatterParams: l.formatterParams,
		refID:           refID,
	}
}

func (l *zlogAdapter) logMsg(msgLevel stdLogger.LogLevel, err error, format string, v ...interface{}) {

	if l.level.IsLogAllowed(msgLevel) {
		var levelMsgHook LevelMsgHook
		levelMsgHook.where = "aca" //utils.GetWhere(3)
		levelMsgHook.level = strings.ToUpper(msgLevel.String())
		if v == nil {
			levelMsgHook.message = format
		} else {
			levelMsgHook.message = fmt.Sprintf(format, v...)
		}

		//El hook de timestamp ya viene del contexto, no se puede eliminar ni modificar el orden
		var hooked zerolog.Logger
		//Se agrega el contexto si existe
		//Se agrega nivel de log, el mensaje y desde donde se llamó
		if l.refID != "" {
			hooked = log.Hook(ContextHook{ref: l.refID}).Hook(levelMsgHook)
		} else {
			hooked = log.Hook(levelMsgHook)
		}

		//Se agregan los detalles del error, si hay uno
		if err != nil {
			hooked = hooked.Hook(ErrorHook{err: err, params: l.formatterParams, formatter: l.formatter})
		}

		//Dispara el log
		hooked.Log().Send()
	}
}

func (l *zlogAdapter) Show(msg string) {
	l.logMsg(stdLogger.LogLevel_Show, nil, msg)
}
func (l *zlogAdapter) Showf(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Show, nil, format, v...)
}

func (l *zlogAdapter) Fatal(err error, msg string) {
	l.logMsg(stdLogger.LogLevel_Fatal, err, msg)
}
func (l *zlogAdapter) Fatalf(err error, format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Fatal, err, format, v...)
}

func (l *zlogAdapter) Error(err error, msg string) {
	l.logMsg(stdLogger.LogLevel_Error, err, msg)
}
func (l *zlogAdapter) Errorf(err error, format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Error, err, format, v...)
}

func (l *zlogAdapter) Warn(msg string) {
	l.logMsg(stdLogger.LogLevel_Warning, nil, msg)
}
func (l *zlogAdapter) Warnf(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Warning, nil, format, v...)
}

func (l *zlogAdapter) Info(msg string) {
	l.logMsg(stdLogger.LogLevel_Info, nil, msg)
}
func (l *zlogAdapter) Infof(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Info, nil, format, v...)
}

func (l *zlogAdapter) Bus(msg string) {
	l.logMsg(stdLogger.LogLevel_Business, nil, msg)
}
func (l *zlogAdapter) Busf(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Business, nil, format, v...)
}

func (l *zlogAdapter) Msg(msg string) {
	l.logMsg(stdLogger.LogLevel_Message, nil, msg)
}
func (l *zlogAdapter) Msgf(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Message, nil, format, v...)
}

func (l *zlogAdapter) Dbg(msg string) {
	l.logMsg(stdLogger.LogLevel_Debug, nil, msg)
}
func (l *zlogAdapter) Dbgf(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Debug, nil, format, v...)
}

func (l *zlogAdapter) Qry(msg string) {
	l.logMsg(stdLogger.LogLevel_Query, nil, msg)
}
func (l *zlogAdapter) Qryf(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Query, nil, format, v...)
}

func (l *zlogAdapter) Trace(msg string) {
	l.logMsg(stdLogger.LogLevel_Trace, nil, msg)
}
func (l *zlogAdapter) Tracef(format string, v ...interface{}) {
	l.logMsg(stdLogger.LogLevel_Trace, nil, format, v...)
}
