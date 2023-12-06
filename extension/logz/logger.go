package logz

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"demo/extension/contextz"
)

var (
	l *slog.Logger

	Group    = slog.Group
	String   = slog.String
	Int64    = slog.Int64
	Int      = slog.Int
	Uint64   = slog.Uint64
	Float64  = slog.Float64
	Bool     = slog.Bool
	Time     = slog.Time
	Duration = slog.Duration
	Any      = slog.Any
)

func init() {
	l = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func Debug(ctx context.Context, msg string, args ...any) {
	l.DebugContext(ctx, prefixWithModuleName(ctx, msg), args...)
}

func DebugNoCtx(msg string, args ...any) {
	l.Debug(msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	l.InfoContext(ctx, prefixWithModuleName(ctx, msg), args...)
}

func InfoNoCtx(msg string, args ...any) {
	l.Info(msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	l.WarnContext(ctx, prefixWithModuleName(ctx, msg), args...)
}

func WarnNoCtx(msg string, args ...any) {
	l.Warn(msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	l.ErrorContext(ctx, prefixWithModuleName(ctx, msg), args...)
}

func ErrorNoCtx(msg string, args ...any) {
	l.Error(msg, args...)
}

func Err(err error) slog.Attr {
	return slog.Any("error", err)
}

func prefixWithModuleName(ctx context.Context, msg string) string {
	module := contextz.ModuleName(ctx)
	if module == "" {
		return msg
	}
	return fmt.Sprintf("[%s] %s", module, msg)
}
