package output

const (
	// Output colors
	red      string = "\033[31m"
	yellow   string = "\033[33m"
	endColor string = "\033[0m"
)

func Red(text string) string {
	return red + text + endColor
}

func Yellow(text string) string {
	return yellow + text + endColor
}
