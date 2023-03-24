package logger

import "context"

type loggerKey struct{}

func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) Logger {
	return getLogger(ctx)
}

func getLogger(ctx context.Context) Logger {
	l := ctx.Value(loggerKey{})
	if l == nil {
		return Discard
	}
	return l.(Logger)
}

func Trace(ctx context.Context, msg string, fields ...any) { getLogger(ctx).Trace(msg, fields...) }
func TraceArgs(ctx context.Context, args map[string]any, msg string, fields ...any) {
	getLogger(ctx).TraceArgs(args, msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...any) { getLogger(ctx).Debug(msg, fields...) }
func DebugArgs(ctx context.Context, args map[string]any, msg string, fields ...any) {
	getLogger(ctx).DebugArgs(args, msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...any) { getLogger(ctx).Info(msg, fields...) }
func InfoArgs(ctx context.Context, args map[string]any, msg string, fields ...any) {
	getLogger(ctx).InfoArgs(args, msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...any) { getLogger(ctx).Warn(msg, fields...) }
func WarnArgs(ctx context.Context, args map[string]any, msg string, fields ...any) {
	getLogger(ctx).WarnArgs(args, msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...any) { getLogger(ctx).Error(msg, fields...) }
func ErrorArgs(ctx context.Context, args map[string]any, msg string, fields ...any) {
	getLogger(ctx).ErrorArgs(args, msg, fields...)
}

func Err(ctx context.Context, err error, msg string, fields ...any) {
	getLogger(ctx).Err(err, msg, fields...)
}

func ErrArgs(ctx context.Context, err error, args map[string]any, msg string, fields ...any) {
	getLogger(ctx).ErrArgs(err, args, msg, fields...)
}
