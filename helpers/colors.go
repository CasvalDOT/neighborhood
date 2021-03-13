package helpers

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

// Colorize ...
func Colorize(content string, color string) string {
	switch color {
	case "red":
		color = red
		break
	case "green":
		color = green
		break
	case "blue":
		color = blue
		break
	case "yellow":
		color = yellow
		break
	case "purple":
		color = purple
		break
	case "gray":
		color = gray
		break
	case "white":
		color = white
		break
	case "cyan":
		color = cyan
		break
	default:
		color = white
	}

	return color + content + reset
}
