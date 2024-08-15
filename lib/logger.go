package lib

import (
	"fmt"
)

const (
	N = "\033[0m"
	W = "\033[1;37m"
	B = "\033[1;34m"
	M = "\033[1;35m"
	R = "\033[1;31m"
	G = "\033[1;32m"
	Y = "\033[1;33m"
	C = "\033[1;36m"
)

func Info(text string) {
	fmt.Println("[" + Y + "*" + N + "] [" + G + "INFO" + N + "] " + text)
}

func Warning(text string) {
	fmt.Println("[" + Y + "+" + N + "] [" + Y + "WARNING" + N + "] " + text)
}

func High(text string) {
	fmt.Println("[" + Y + "-" + N + "] [" + R + "CRITICAL" + N + "] " + text)
}
