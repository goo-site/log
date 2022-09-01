package log

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/goo-site/log/common"
)

var lg logger = logger{level: common.InfoLevel}

type logger struct {
	mu      sync.Mutex
	level   common.LogLevel // 日志等级
	writers []Writer        // 日志输出
}

type Writer interface {
	// 自由设置输出格式
	Format(prefix string, timestamp int64, filepath string, line int, s string) string

	// 将string写入目的io
	Write(s string)
}

func AddWriter(writer ...Writer) {
	lg.writers = append(lg.writers, writer...)
}

func SetLogLevel(level string) {
	levelcode := common.LevelCode(level)
	if levelcode == common.UnsupportedLevel {
		fmt.Printf("[system] unsupported log level\n")
		return
	}
	lg.level = levelcode
}

func (lg *logger) output(prefix string, s string) {
	_, filepath, line, ok := runtime.Caller(2)
	if !ok {
		filepath = "???"
		line = 0
	}

	lg.mu.Lock()
	defer lg.mu.Unlock()

	now := time.Now().Unix()
	for _, writer := range lg.writers {
		s = writer.Format(prefix, now, filepath, line, s)
		writer.Write(s)
	}
}

func Debug(format string, args ...interface{}) {
	if lg.level < common.InfoLevel {
		return
	}
	lg.output("Debug", fmt.Sprintf(format, args...))
}

func Info(format string, args ...interface{}) {
	if lg.level < common.InfoLevel {
		return
	}
	lg.output("Info", fmt.Sprintf(format, args...))
}

func Warn(format string, args ...interface{}) {
	if lg.level < common.InfoLevel {
		return
	}
	lg.output("Warn", fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	if lg.level < common.InfoLevel {
		return
	}
	lg.output("Error", fmt.Sprintf(format, args...))
}

func Panic(format string, args ...interface{}) {
	if lg.level < common.InfoLevel {
		return
	}
	lg.output("Panic", fmt.Sprintf(format, args...))
	panic(fmt.Sprintf(format, args...))
}

func Fatal(format string, args ...interface{}) {
	if lg.level < common.InfoLevel {
		return
	}
	lg.output("[FATAL] ", fmt.Sprintf(format, args...))
	os.Exit(1)
}
