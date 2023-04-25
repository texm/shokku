package db

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

var durationUnits = map[time.Duration]string{
	time.Nanosecond:  "elapsed_ns",
	time.Microsecond: "elapsed_us",
	time.Millisecond: "elapsed_ms",
	time.Second:      "elapsed",
	time.Minute:      "elapsed_min",
	time.Hour:        "elapsed_hr",
}

type Logger struct{}

func (l Logger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l Logger) Error(ctx context.Context, msg string, opts ...interface{}) {
	zerolog.Ctx(ctx).Error().Msg(fmt.Sprintf(msg, opts...))
}

func (l Logger) Warn(ctx context.Context, msg string, opts ...interface{}) {
	zerolog.Ctx(ctx).Warn().Msg(fmt.Sprintf(msg, opts...))
}

func (l Logger) Info(ctx context.Context, msg string, opts ...interface{}) {
	zerolog.Ctx(ctx).Info().Msg(fmt.Sprintf(msg, opts...))
}

func (l Logger) Trace(ctx context.Context, begin time.Time, f func() (string, int64), err error) {
	zl := zerolog.Ctx(ctx)
	var event *zerolog.Event

	if err != nil {
		event = zl.Debug()
	} else {
		event = zl.Trace()
	}

	durationKey, found := durationUnits[zerolog.DurationFieldUnit]
	if !found {
		zl.Error().
			Dur("zerolog.DurationFieldUnit", zerolog.DurationFieldUnit).
			Msg("unknown value for zerolog.DurationFieldUnit")
		durationKey = "elapsed_"
	}

	event.Dur(durationKey, time.Since(begin))

	sql, rows := f()
	if sql != "" {
		event.Str("sql", sql)
	}
	if rows > -1 {
		event.Int64("rows", rows)
	}

	event.Send()

	return
}
