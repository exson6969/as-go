package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ResponseStructure struct {
	Event           string                 `json:"event"`
	EventType       string                 `json:"event_type"`
	AppID           string                 `json:"app_id"`
	UserID          string                 `json:"user_id"`
	MessageID       string                 `json:"message_id"`
	PageTitle       string                 `json:"page_title"`
	PageURL         string                 `json:"page_url"`
	BrowserLanguage string                 `json:"browser_language"`
	ScreenSize      string                 `json:"screen_size"`
	Attributes      map[string]interface{} `json:"attributes"`
	Traits          map[string]interface{} `json:"traits"`
}

func main() {
	http.HandleFunc("/", processRequest)
	http.ListenAndServe(":8080", nil)
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	transformedData := convertToResponseStructure(requestData)
	go sendDataToWebhook(transformedData)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request Processed")
}

func convertToResponseStructure(data map[string]interface{}) ResponseStructure {
	transformedData := ResponseStructure{
		Event:           data["ev"].(string),
		EventType:       data["et"].(string),
		AppID:           data["id"].(string),
		UserID:          data["uid"].(string),
		MessageID:       data["mid"].(string),
		PageTitle:       data["t"].(string),
		PageURL:         data["p"].(string),
		BrowserLanguage: data["l"].(string),
		ScreenSize:      data["sc"].(string),
		Attributes:      make(map[string]interface{}),
		Traits:          make(map[string]interface{}),
	}

	for key, value := range data {
		if strings.HasPrefix(key, "atrk") {
			attrNum := strings.TrimPrefix(key, "atrk")
			attrTypeKey := "atrt" + attrNum
			attrValueKey := "atrv" + attrNum

			attrType, typeExists := data[attrTypeKey].(string)
			attrValue, valueExists := data[attrValueKey].(string)

			if typeExists && valueExists {
				transformedData.Attributes[value.(string)] = map[string]interface{}{
					"value": attrValue,
					"type":  attrType,
				}
			}
		} else if strings.HasPrefix(key, "uatrk") {
			traitNum := strings.TrimPrefix(key, "uatrk")
			traitTypeKey := "uatrt" + traitNum
			traitValueKey := "uatrv" + traitNum

			traitType, typeExists := data[traitTypeKey].(string)
			traitValue, valueExists := data[traitValueKey].(string)

			if typeExists && valueExists {
				transformedData.Traits[value.(string)] = map[string]interface{}{
					"value": traitValue,
					"type":  traitType,
				}
			}
		}
	}

	return transformedData
}

func sendDataToWebhook(data ResponseStructure) {
	url := "https://webhook.site/18b6e968-b9a0-4f58-bb37-d386a3f99e29"

	jsonValue, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Success")
}
