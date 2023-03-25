package service

import (
	"chatgpt-go/global"
	"chatgpt-go/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func ChatConfig() (model.ChatConfig, error) {
	balance, err := fetchBalance()
	if err != nil {
		balance = "error"
	}

	reverseProxy := global.Config.System.ReverseProxy
	if reverseProxy == "" {
		reverseProxy = "-"
	}

	httpsProxy := global.Config.System.HttpsProxy

	if httpsProxy == "" {
		httpsProxy = "-"
	}

	socksProxy := "-"
	socksHost := global.Config.System.SocksHost
	socksPort := global.Config.System.SocksPort
	if socksHost != "" && socksPort != "" {
		socksProxy = fmt.Sprintf("%s:%s", socksHost, socksPort)
	}

	config := model.ChatConfig{
		Message: "",
		Data: model.ChatConfigData{
			APIModel:     "ChatGPTAPI",
			ReverseProxy: reverseProxy,
			TimeoutMs:    60000,
			SocksProxy:   socksProxy,
			HttpsProxy:   httpsProxy,
			Balance:      balance,
		},
		Status: "Success",
	}

	return config, nil
}

func fetchBalance() (string, error) {
	openAPIBaseURL := global.Config.System.OpenAPIBaseURL

	if global.Config.System.OpenAIKey == "" {
		return "-", nil
	}

	apiBaseURL := "https://api.openai.com"
	if openAPIBaseURL != "" {
		apiBaseURL = openAPIBaseURL
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dashboard/billing/credit_grants", apiBaseURL), nil)
	if err != nil {
		return "-", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+global.Config.System.OpenAIKey)

	resp, err := client.Do(req)
	if err != nil {
		return "-", err
	}
	defer resp.Body.Close()

	var data struct {
		TotalAvailable float64 `json:"total_available"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "-", err
	}

	return fmt.Sprintf("%.3f", data.TotalAvailable), nil
}
