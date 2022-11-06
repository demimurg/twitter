package log

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	stdout *os.File
	ctx    = context.Background()
	buffer = make([]byte, 512)
)

func init() {
	realStdout := os.Stdout // copy to return after logger building
	r, w, _ := os.Pipe()
	os.Stdout, stdout = w, r
	ChangeConfig(func(c *zap.Config) {
		c.EncoderConfig.TimeKey = zapcore.OmitKey
	})
	os.Stdout = realStdout
}

func readStdout() string {
	n, _ := stdout.Read(buffer)
	return strings.Trim(string(buffer[:n]), "\n")
}

func TestLogger(t *testing.T) {
	t.Run("simple info log", func(t *testing.T) {
		Info(ctx, "some message")
		assert.Equal(t, `{"level":"info","msg":"some message"}`, readStdout())
	})

	t.Run("debug log with key and value", func(t *testing.T) {
		Debug(ctx, "some message",
			"key", "value")
		assert.Equal(t, `{"level":"debug","msg":"some message","key":"value"}`, readStdout())
	})

	t.Run("context with key value pairs", func(t *testing.T) {
		ctx := With(ctx, "key", "value")
		Info(ctx, "first message")
		assert.Equal(t, `{"level":"info","msg":"first message","key":"value"}`, readStdout())
		Info(ctx, "second message")
		assert.Equal(t, `{"level":"info","msg":"second message","key":"value"}`, readStdout())
	})

	t.Run("log with object", func(t *testing.T) {
		type Object struct{ FieldA, FieldB string }
		Info(ctx, "something",
			"object", Object{FieldA: "a", FieldB: "b"})
		assert.Equal(t, `{"level":"info","msg":"something","object":{"FieldA":"a","FieldB":"b"}}`, readStdout())
	})

	t.Run("log with duration field", func(t *testing.T) {
		Info(ctx, "something",
			"duration", time.Second+100*time.Millisecond)
		assert.Equal(t, `{"level":"info","msg":"something","duration":"1.1s"}`, readStdout())
	})

	t.Run("log error with shorthand", func(t *testing.T) {
		Error(ctx, "something", errors.New("bad thing happend"))
		assert.Equal(t, `{"level":"error","msg":"something","error":"bad thing happend"}`, readStdout())
	})
}
