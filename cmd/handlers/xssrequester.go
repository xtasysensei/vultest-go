package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"strings"

	"github.com/anaskhan96/soup"
	"github.com/spf13/cobra"
	"github.com/xtasysensei/vultest/cmd/utils"
)

type ScanPayload struct {
	VulnType  string
	TimeFound time.Time
	UrlInfo   UrlInfo
}
type UrlInfo struct {
	BadURL     string
	BadPayload string
}

func generatePayload(eff int) (string, error) {
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

// func generatePayload(eff int) (string, error) {
// 	payloads := []string{
// 		"prompt(5000/200)",
// 		"alert(6000/3000)",
// 		"alert(document.cookie)",
// 		"prompt(document.cookie)",
// 		"console.log(5000/3000)",
// 	}

// 	switch eff {
// 	case 1:
// 		return "<script>" + payloads[utils.RandRange(0, 4)] + "</script>", nil
// 	case 2:
// 		return "<script>" + payloads[utils.RandRange(0, 4)] + "</script>", nil
// 	case 3:
// 		return "<script>" + payloads[utils.RandRange(0, 4)] + "</script>", nil
// 	case 4:
// 		return "<script>" + payloads[utils.RandRange(0, 4)] + "</script>", nil
// 	case 5:
// 		return "<script>" + payloads[utils.RandRange(0, 4)] + "</script>", nil
// 	case 6:
// 		return "<script>" + payloads[utils.RandRange(0, 4)] + "</script>", nil
// 	default:
// 		return "", nil
// 	}
// }

type Keys struct {
	KeyType string
	KeyName string
	Value   string
}

func PostMethod(childURL, payload string) ([]Keys, error) {
	resp, err := soup.Get(childURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to URL %s: %v", childURL, err)
	}

	doc := soup.HTMLParse(resp)
	forms := doc.FindAll("form")

	var allKeys []Keys
	xssDetected := false

	for _, form := range forms {
		action := form.Attrs()["action"]
		method := strings.ToLower(form.Attrs()["method"])

		if method == "post" {
			newChildURL, err := url.JoinPath(childURL, action)
			if err != nil {
				return nil, fmt.Errorf("failed to join path %s to %s: %v", childURL, action, err)
			}

			utils.Warning("Target has form with POST method: " + utils.C + newChildURL)
			utils.Info("Collecting form input keys.....")

			var formKeys []Keys

			inputAreas := form.FindAll("input")
			textAreas := form.FindAll("textarea")
			inputAreas = append(inputAreas, textAreas...)

			for _, inputArea := range inputAreas {
				keyType := inputArea.Attrs()["type"]
				keyName := inputArea.Attrs()["name"]

				var key Keys
				if keyType == "submit" {
					utils.Info("Form key name: " + utils.G + keyName + utils.N + " value: " + utils.G + "<Submit Confirm>")
					key = Keys{
						KeyType: keyType,
						KeyName: keyName,
					}
				} else {
					utils.Info("Form key name: " + utils.G + keyName + utils.N + " value: " + utils.G + payload)
					key = Keys{
						KeyType: keyType,
						KeyName: keyName,
						Value:   payload,
					}
				}
				formKeys = append(formKeys, key)
			}

			// Construct form data
			// TODO:change it from form data to all types of data
			formData := url.Values{}
			for _, key := range formKeys {
				formData.Set(key.KeyName, key.Value)
			}

			utils.Info("Sending payload (POST) method...")
			resp, err := http.PostForm(newChildURL, formData)
			if err != nil {
				return nil, fmt.Errorf("failed to send POST request: %v", err)
			}
			defer resp.Body.Close()

			// Check for XSS
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read response body: %v", err)
			}
			if strings.Contains(string(body), payload) {
				utils.High("Detected XSS (POST) at " + childURL)
				utils.High("Post data: " + fmt.Sprintf("%+v", formKeys))
				xssDetected = true

				urlInfo := UrlInfo{
					BadURL:     childURL,
					BadPayload: payload,
				}
				w := ScanPayload{
					VulnType:  "Detected XSS(POST)",
					TimeFound: time.Now(),
					UrlInfo:   urlInfo,
				}
				var buffer bytes.Buffer
				encoder := json.NewEncoder(&buffer)
				encoder.SetEscapeHTML(false)
				err := encoder.Encode(w)

				cobra.CheckErr(err)

				err = utils.WriteToFile(buffer.Bytes())
				cobra.CheckErr(err)

			}

			allKeys = append(allKeys, formKeys...)
		}
	}

	if len(allKeys) == 0 {
		return nil, fmt.Errorf("no POST forms found on the page")
	}

	if !xssDetected {
		utils.Info("No XSS vulnerabilities detected in POST forms, but further testing is recommended")
	}

	return allKeys, nil
}

func GetMethod(childURL string, payload string) ([]Keys, error) {
	resp, err := soup.Get(childURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to URL %s: %v", childURL, err)
	}

	doc := soup.HTMLParse(resp)
	forms := doc.FindAll("form")

	var allKeys []Keys
	xssDetected := false

	for _, form := range forms {
		action := form.Attrs()["action"]
		method := strings.ToLower(form.Attrs()["method"])

		if method == "get" {
			newChildURL, err := url.JoinPath(childURL, action)
			if err != nil {
				return nil, fmt.Errorf("failed to join path %s to %s: %v", childURL, action, err)
			}

			utils.Warning("Target has form with GET method: " + utils.C + newChildURL)
			utils.Info("Collecting form input keys.....")
			var formKeys []Keys

			inputAreas := form.FindAll("input")
			textAreas := form.FindAll("textarea")
			inputAreas = append(inputAreas, textAreas...)

			for _, inputArea := range inputAreas {
				keyType := inputArea.Attrs()["type"]
				keyName := inputArea.Attrs()["name"]

				var key Keys
				if keyType == "submit" {
					utils.Info("Form key name: " + utils.G + keyName + utils.N + " value: " + utils.G + "<Submit Confirm>")
					key = Keys{
						KeyType: keyType,
						KeyName: keyName,
					}
				} else {
					utils.Info("Form key name: " + utils.G + keyName + utils.N + " value: " + utils.G + payload)
					key = Keys{
						KeyType: keyType,
						KeyName: keyName,
						Value:   payload,
					}
				}
				formKeys = append(formKeys, key)

			}

			utils.Info("Sending payload (GET) method...")
			resp, err := http.Get(newChildURL)
			if err != nil {
				return nil, fmt.Errorf("failed to send GET request: %v", err)
			}
			defer resp.Body.Close()
		}
	}
	return allKeys, nil
}

func ConnectAndRequest(childURL string) {
	payload, err := generatePayload(utils.RandRange(1, 6))
	if err != nil {
		log.Printf("failed to generate payload: %v", err)
		return
	}
	_, _ = PostMethod(childURL, payload)
}
