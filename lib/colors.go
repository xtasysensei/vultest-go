package lib

import (
	"os/exec"
	"runtime"
	"strings"
)

var (
	colors  bool = true
	machine      = runtime.GOOS
	white   string
	green   string
	red     string
	yellow  string
	purple  string
	end     string
	back    string
	bold    string
	blue    string
	info    string
	que     string
	bad     string
	good    string
	run     string
	grey    string
	cyan    string
	gray    string
	reset   string
)

func init() {
	if machine == "windows" {
		cmd := exec.Command("cmd", "ver")
		output, err := cmd.Output()
		if err != nil {
			colors = false
		}
		versionInfo := strings.TrimSpace(string(output))
		if strings.Contains(versionInfo, "Version 10") || strings.Contains(versionInfo, "Version 11") {
			colors = true
		} else {
			colors = false
		}
	} else if strings.HasPrefix(strings.ToLower(machine), "darwin") || strings.HasPrefix(strings.ToLower(machine), "ios") {
		colors = false
	} else {
		white = "\033[97m"
		green = "\033[92m"
		red = "\033[91m"
		yellow = "\033[93m"
		purple = "\033[35m"
		end = "\033[0m"
		back = "\033[7;91m"
		bold = "\033[1m"
		blue = "\033[94m"
		info = "\033[93m[!]\033[0m"
		que = "\033[94m[?]\033[0m"
		bad = "\033[91m[-]\033[0m"
		good = "\033[92m[+]\033[0m"
		run = "\033[97m[~]\033[0m"
		grey = "\033[7;90m"
		cyan = "\u001B[36m"
		gray = "\033[90m"
		reset = "\033[0m"
	}
}
