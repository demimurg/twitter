package log

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ------------------------------ CONFIGURE ------------------------------

var (
	// l is a singleton (global instance) of a zap logger
	l *zap.SugaredLogger
	// config for the global logger
	cfg *zap.Config
)

func init() {
	zapDefaultConfig := zap.NewProductionConfig()
	cfg = &zapDefaultConfig

	cfg.Level.SetLevel(zap.DebugLevel)
	cfg.OutputPaths = []string{"stdout"}
	cfg.DisableCaller = true
	cfg.DisableStacktrace = true
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	buildLogger()
}

// SetLevel will parse and set "debug/info/error" levels. Any case accepted
func SetLevel(name string) {
	level, err := zap.ParseAtomicLevel(name)
	if err != nil {
		panic(err)
	}
	ChangeConfig(func(cfg *zap.Config) {
        cfg.Level = level
    })
}

// ChangeConfig allows to override default zap logger config from init function
func ChangeConfig(override func(*zap.Config)) {
	override(cfg)
	buildLogger()
}

func buildLogger() {
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	l = logger.Sugar()
}

// ------------------------------ CONTEXT ------------------------------

// zapKey identifies zap logger in context
type zapKey string

// With will add any number of key and value pairs to context, any follow log with ctx will inherit
func With(ctx context.Context, keyVal ...any) context.Context {
	return context.WithValue(ctx, zapKey(""), from(ctx).With(keyVal...))
}

func from(ctx context.Context) *zap.SugaredLogger {
	logger := ctx.Value(zapKey(""))
	if logger == nil {
		return l
	}
	return logger.(*zap.SugaredLogger)
}

// ------------------------------ LOGGING ------------------------------

// Debug logs should be used in development environment only
func Debug(ctx context.Context, msg string, keyVal ...any) {
	from(ctx).Debugw(msg, keyVal...)
}

// Info show essential events in program
func Info(ctx context.Context, msg string, keyVal ...any) {
	from(ctx).Infow(msg, keyVal...)
}

// Error used to handle exceptional case. You can use `Error(ctx, "some msg", err)` as shorthand for `Error(ctx, "some msg", "error", err)“
func Error(ctx context.Context, msg string, keyVal ...any) {
	if len(keyVal) > 0 {
		err, ok := keyVal[0].(error)
		if ok {
			keyVal[0] = zap.Error(err)
		}
	}
	from(ctx).Errorw(msg, keyVal...)
}

// Fatal will print log to stdout and then os.Exit(1) will be called
func Fatal(ctx context.Context, msg string, keyVal ...any) {
	from(ctx).Fatalw(msg, keyVal...)
}

// Panic is Fatal replacement for cases, when you need to evaluate defer before exit
func Panic(ctx context.Context, msg string, keyVal ...any) {
	from(ctx).Panicw(msg, keyVal...)
}

// Print is very simple func that can print any value as json, can be used for temprorary debug
func Print(val any) {
	l.Debugw("temprorary debug log", "value", val)
}
