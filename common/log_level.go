package common

/* 日志级别 */
type LogLevel int

const (
	UnsupportedLevel LogLevel = 0
	FatalLevel       LogLevel = 1
	PanicLevel       LogLevel = 2
	ErrorLevel       LogLevel = 3
	WarnLevel        LogLevel = 4
	InfoLevel        LogLevel = 5
	DebugLevel       LogLevel = 6
)

func LevelCode(level string) LogLevel {
	switch level {
	case "Debug":
		return DebugLevel
	case "Info":
		return InfoLevel
	case "Warn":
		return WarnLevel
	case "Error":
		return ErrorLevel
	case "Panic":
		return PanicLevel
	case "Fatal":
		return FatalLevel
	default:
		return UnsupportedLevel
	}
}
