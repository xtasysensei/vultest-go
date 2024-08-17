package utils

import (
	"os/exec"
	"runtime"
	"strings"
)

var (
	Colors    = true
	Machine   = runtime.GOOS
	White     string
	Green     string
	Red       string
	Yellow    string
	Purple    string
	End       string
	Back      string
	Bold      string
	Underline string
	Blue      string
	Que       string
	Bad       string
	Good      string
	Run       string
	Grey      string
	Cyan      string
	Gray      string
	Reset     string
)

func init() {
	machineLower := strings.ToLower(Machine)

	switch {
	case machineLower == "windows":
		cmd := exec.Command("cmd", "ver")
		output, err := cmd.Output()
		if err != nil || (!strings.Contains(string(output), "Version 10") && !strings.Contains(string(output), "Version 11")) {
			Colors = false
		}
	case strings.HasPrefix(machineLower, "darwin"), strings.HasPrefix(machineLower, "ios"):
		Colors = false
	default:
		White = "\033[97m"
		Green = "\033[92m"
		Red = "\033[91m"
		Yellow = "\033[93m"
		Purple = "\033[35m"
		End = "\033[0m"
		Back = "\033[7;91m"
		Bold = "\033[1m"
		Blue = "\033[94m"
		Que = "\033[94m[?]\033[0m"
		Bad = "\033[91m[-]\033[0m"
		Good = "\033[92m[+]\033[0m"
		Run = "\033[97m[~]\033[0m"
		Grey = "\033[7;90m"
		Cyan = "\u001B[36m"
		Gray = "\033[90m"
		Reset = "\033[0m"
	}
}
