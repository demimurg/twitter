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

func TestLog(t *testing.T) {
	initLoggerForTest()
	ctx := context.Background()

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
		Error(ctx, "something", errors.New("bad thing happened"))
		assert.Equal(t, `{"level":"error","msg":"something","error":"bad thing happened"}`, readStdout())
	})

	t.Run("temprorary unstructured log", func(t *testing.T) {
		Print("something")
		assert.Equal(t, `{"level":"debug","msg":"temprorary debug log","value":"something"}`, readStdout())
	})
}

func TestSetLevel(t *testing.T) {
	SetLevel("info")
	initLoggerForTest()

	Debug(context.Background(), "some message")
	// use deadline to unblock read stdout call
	_ = stdout.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
	assert.Equal(t, "", readStdout())
	// drop deadline for stdout
	_ = stdout.SetReadDeadline(time.Time{})
}

var (
	stdout *os.File // fake stdout for logger
	buffer = make([]byte, 512)
)

func initLoggerForTest() {
	// copy to return after logger building,
	// to be able to see test info output
	realStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	os.Stdout, stdout = w, r
	// build logger that will use pipe to write,
	// so we can read this data in code and not in terminal
	ChangeConfig(func(c *zap.Config) {
		c.EncoderConfig.TimeKey = zapcore.OmitKey
	})
	os.Stdout = realStdout
}

func readStdout() string {
	n, err := stdout.Read(buffer)
	if err != nil {
		return ""
	}
	return strings.Trim(string(buffer[:n]), "\n")
}
