package handlers

import (
	"strings"

	"github.com/xtasysensei/vultest/cmd/utils"
)

func GeneratePayload(eff int) (string, error) {
	payloads := []string{
		"prompt(5000/200)",
		"alert(6000/3000)",
		"alert(document.cookie)",
		"prompt(document.cookie)",
		"console.log(5000/3000)",
	}

	if eff == 1 {
		return "<script/>" + payloads[utils.RandRange(0, 4)] + `<\script\>`, nil
	} else if eff == 2 {
		return `<\script/>` + payloads[utils.RandRange(0, 4)] + `<\\script>`, nil

	} else if eff == 3 {
		return `<\script\>` + payloads[utils.RandRange(0, 4)] + `<//script>`, nil
	} else if eff == 4 {
		return `<script>` + payloads[utils.RandRange(0, 4)] + `<\script/>`, nil

	} else if eff == 5 {
		return `<script>` + payloads[utils.RandRange(0, 4)] + `<//script>`, nil

	} else if eff == 6 {
		return `<script>` + payloads[utils.RandRange(0, 4)] + `</script>`, nil

	}

	return "", nil
}

func ParseHtml(html string) {
	doc, err := html.Parse(strings.NewReader(html))
}
