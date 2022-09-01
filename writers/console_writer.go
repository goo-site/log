package writers

import (
	"fmt"
	"path"
	"time"
)

type ConsoleWriter struct {
}

func (c *ConsoleWriter) Format(prefix string, timestamp int64, filepath string, line int, s string) string {
	color := PrefixColor(prefix)
	filename := path.Base(filepath)
	s1 := fmt.Sprintf("\033[%d;%dm[%s] [%s]\033[0m", HighLight, color, prefix, time.Unix(timestamp, 0).Format("15:04:05"))
	s2 := fmt.Sprintf("\033[%d;%d;%dm%s:%d\033[0m", HighLight, UnderLine, color, filename, line)
	s3 := fmt.Sprintf("\033[%d;%dm%s\033[0m", HighLight, color, s)
	return fmt.Sprintf("%s %s %s\n", s1, s2, s3)
}

func (c *ConsoleWriter) Write(s string) {
	fmt.Printf("%s", s)
}

/* 输出颜色 */
type Color int

const (
	Red     Color = 31 // 红色
	Green   Color = 32 // 绿色
	Yellow  Color = 33 // 黄色
	Blue    Color = 34 // 蓝色
	Magenta Color = 35 // 洋红色
	Cyan    Color = 36 // 青蓝色
	White   Color = 37 // 白色
)

func PrefixColor(prefix string) Color {
	switch prefix {
	case "Debug":
		return Blue
	case "Info":
		return Green
	case "Warn":
		return Yellow
	case "Error":
		return Red
	case "Panic":
		return Red
	case "Fatal":
		return Red
	default:
		return White
	}
}

/* 输出样式 */
type StringAttr int

const (
	Default   StringAttr = 0 // 默认设置
	HighLight StringAttr = 1 // 高亮
	UnderLine StringAttr = 4 // 下划线
)
