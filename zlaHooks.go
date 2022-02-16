package zla

import (
	"github.com/rs/zerolog"

	e2hformat "github.com/cdleo/go-e2h/formatter"
)

type ContextHook struct {
	ref string
}

func (h ContextHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("ref", h.ref)
}

type LevelMsgHook struct {
	level   string
	where   string
	message string
}

func (h LevelMsgHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.Str("level", h.level)
	e.Str("message", h.message)
	e.Str("where", h.where)
}

type ErrorHook struct {
	err       error
	params    e2hformat.Params
	formatter e2hformat.Formatter
}

func (h ErrorHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.RawJSON("details", []byte(h.formatter.Format(h.err, h.params)))
}
