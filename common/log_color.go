package common

type LogColor int

const (
	UnsupportedColor LogColor = 0  // 不支持
	Red              LogColor = 31 // 红色
	Green            LogColor = 32 // 绿色
	Yellow           LogColor = 33 // 黄色
	Blue             LogColor = 34 // 蓝色
	Magenta          LogColor = 35 // 洋红色
	Cyan             LogColor = 36 // 青蓝色
	White            LogColor = 37 // 白色
)

func ColorCode(color string) LogColor {
	switch color {
	case "Red":
		return Red
	case "Green":
		return Green
	case "Yellow":
		return Yellow
	case "Blue":
		return Blue
	case "Magenta":
		return Magenta
	case "Cyan":
		return Cyan
	case "White":
		return White
	default:
		return UnsupportedColor
	}
}
