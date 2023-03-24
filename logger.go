package logger

const (
	logLevelTrace = 0
	logLevelDebug = 1
	logLevelInfo  = 2
	logLevelWarn  = 3
	logLevelError = 4
)

type levelLogger interface {
	Log(level int, msg string, fields ...any)
	LogArgs(level int, args map[string]any, msg string, fields ...any)
}

type Logger interface {
	Trace(msg string, fields ...any)
	TraceArgs(args map[string]any, msg string, fields ...any)

	Debug(msg string, fields ...any)
	DebugArgs(args map[string]any, msg string, fields ...any)

	Info(msg string, fields ...any)
	InfoArgs(args map[string]any, msg string, fields ...any)

	Warn(msg string, fields ...any)
	WarnArgs(args map[string]any, msg string, fields ...any)

	Error(msg string, fields ...any)
	ErrorArgs(args map[string]any, msg string, fields ...any)

	Err(err error, msg string, fields ...any)
	ErrArgs(err error, args map[string]any, msg string, fields ...any)
}

func wrap(ll levelLogger) Logger {
	return &levelWrapper{ll}
}

type levelWrapper struct {
	ll levelLogger
}

func (w *levelWrapper) Trace(msg string, fields ...any) {
	w.ll.Log(logLevelTrace, msg, fields...)
}

func (w *levelWrapper) TraceArgs(args map[string]any, msg string, fields ...any) {
	w.ll.LogArgs(logLevelTrace, args, msg, fields...)
}

func (w *levelWrapper) Debug(msg string, fields ...any) {
	w.ll.Log(logLevelDebug, msg, fields...)
}

func (w *levelWrapper) DebugArgs(args map[string]any, msg string, fields ...any) {
	w.ll.LogArgs(logLevelDebug, args, msg, fields...)
}

func (w *levelWrapper) Info(msg string, fields ...any) {
	w.ll.Log(logLevelInfo, msg, fields...)
}

func (w *levelWrapper) InfoArgs(args map[string]any, msg string, fields ...any) {
	w.ll.LogArgs(logLevelInfo, args, msg, fields...)
}

func (w *levelWrapper) Warn(msg string, fields ...any) {
	w.ll.Log(logLevelWarn, msg, fields...)
}

func (w *levelWrapper) WarnArgs(args map[string]any, msg string, fields ...any) {
	w.ll.LogArgs(logLevelWarn, args, msg, fields...)
}

func (w *levelWrapper) Error(msg string, fields ...any) {
	w.ll.Log(logLevelError, msg, fields...)
}

func (w *levelWrapper) ErrorArgs(args map[string]any, msg string, fields ...any) {
	w.ll.LogArgs(logLevelError, args, msg, fields...)
}

func (w *levelWrapper) Err(err error, msg string, fields ...any) {
	w.ll.LogArgs(logLevelError, map[string]any{
		"error": err,
	}, msg, fields...)
}

func (w *levelWrapper) ErrArgs(err error, args map[string]any, msg string, fields ...any) {
	args["error"] = err
	w.ll.LogArgs(logLevelError, args, msg, fields...)
}
