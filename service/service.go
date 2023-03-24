package service

import (
	"chatgpt-go/global"
	"chatgpt-go/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ChatConfig() (model.ChatConfig, error) {
	balance, err := fetchBalance()
	if err != nil {
		balance = "error"
	}

	reverseProxy := os.Getenv("API_REVERSE_PROXY")
	if reverseProxy == "" {
		reverseProxy = "-"
	}

	httpsProxy := os.Getenv("HTTPS_PROXY")

	if httpsProxy == "" {
		httpsProxy = "-"
	}

	socksProxy := "-"
	socksHost := os.Getenv("SOCKS_PROXY_HOST")
	socksPort := os.Getenv("SOCKS_PROXY_PORT")
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
	openAPIBaseURL := os.Getenv("OPENAI_API_BASE_URL")

	if global.OpenAIKey == "" {
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
	req.Header.Set("Authorization", "Bearer "+global.OpenAIKey)

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

	fmt.Println(data)

	return fmt.Sprintf("%.3f", data.TotalAvailable), nil
}
